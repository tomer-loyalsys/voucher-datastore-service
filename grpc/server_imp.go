// Copyright 2018 The Loyalsys Authors. All rights reserved.

package grpc

import (
	"github.com/loyalsys/couchbase"
	"github.com/loyalsys/error"
	"github.com/loyalsys/log"
	"github.com/loyalsys/voucher-datastore-service/datastore"
	"github.com/loyalsys/voucher-datastore-service/grpc/proto"
	"golang.org/x/net/context"
	"io"
)

type Service struct {
	Ds         *datastore.Datastore
	AppVersion string
}

func (s *Service) Ping(ctx context.Context, in *lsvoucherds.PingReq) (out *lsvoucherds.PingRes, err error) {
	defer RecoverAndResponse(ctx, &err)

	o := lsvoucherds.PingRes{Version: s.AppVersion}
	return &o, nil
}

func (s *Service) Test(ctx context.Context, in *lsvoucherds.TestReq) (out *lsvoucherds.TestRes, err error) {
	defer RecoverAndResponse(ctx, &err)

	o := lsvoucherds.TestRes{IsValid: true}
	return &o, nil
}

func (s *Service) UploadToPool(stream lsvoucherds.Grpc_UploadToPoolServer) (err error) {
	defer RecoverAndResponse(stream.Context(), &err)

	var region *string
	var customerId *string
	var poolId *string
	var uploadId *string
	var logTags *lslog.Tags
	isFirstStreamRecv := true

	for {
		in, err := stream.Recv()

		if err == io.EOF {
			uploadStat, err := s.addToPool(stream.Context(), *region, *customerId, *poolId, *uploadId)
			if err != nil {
				lslog.Errorf(stream.Context(), err, "failed to add pool upload to pool.").WithTags(logTags).Write()
				return lserr.WrapErrf(err, "failed to add pool upload to pool.")
			}

			return stream.SendAndClose(&lsvoucherds.UploadToPoolRes{TotalUpload: uploadStat.Total})
		}
		if err != nil {
			lslog.Errorf(stream.Context(), err, "error in stream receive.").WithTags(logTags).Write()
			return lserr.WrapErrf(err, "error in stream receive.")
		}

		// set input params to use when err = io.EOF
		region = &in.Region
		customerId = &in.CustomerId
		poolId = &in.PoolId
		uploadId = &in.UploadId
		logTags = &lslog.Tags{lslog.Tag_CustomerId: in.CustomerId}

		r := lscb.Region(in.Region)

		if isFirstStreamRecv {
			// check that pool upload not already exists from other request
			isExists, err := s.Ds.Voucher.Region(r).IsPoolUploadExists(in.CustomerId, in.PoolId, in.UploadId)
			if err != nil {
				lslog.Errorf(stream.Context(), err, "failed to check if pool upload exists.").WithTags(logTags).Write()
				return lserr.WrapErrf(err, "failed to check if pool upload exists.")
			}

			if *isExists {
				lslog.Errorf(stream.Context(), err, "pool upload with same upload id already exists.").WithTags(logTags).Write()
				return lserr.NewErrf("pool upload with same upload id already exists.")
			}

			isFirstStreamRecv = false
		}

		// upload vouchers
		for _, voucher := range in.Vouchers {
			err = s.Ds.Voucher.Region(r).AppendToPoolUpload(in.CustomerId, in.PoolId, in.UploadId, voucher)
			if err != nil {
				lslog.Errorf(stream.Context(), err, "failed to insert voucher to pool upload doc.").WithTags(logTags).Write()
				return lserr.WrapErrf(err, "failed to insert voucher to pool upload doc.")
			}
		}
	}
}

func (s *Service) GetPoolAvailability(ctx context.Context, in *lsvoucherds.GetPoolAvailabilityReq) (out *lsvoucherds.GetPoolAvailabilityRes, err error) {
	defer RecoverAndResponse(ctx, &err)

	region := lscb.Region(in.Region)
	logTags := &lslog.Tags{lslog.Tag_CustomerId: in.CustomerId}

	res := make(map[string]*lsvoucherds.GetPoolAvailabilityRes_GetPoolAvailability, len(in.PoolIds))
	for _, poolId := range in.PoolIds {
		total, available, err := s.Ds.Voucher.Region(region).GetPoolAvailability(in.CustomerId, poolId)
		if err != nil {
			lslog.Errorf(ctx, err, "failed to get pool availability.").WithTags(logTags).Write()
			return nil, lserr.WrapErrf(err, "failed to get pool availability.")
		}

		res[poolId] = &lsvoucherds.GetPoolAvailabilityRes_GetPoolAvailability{Total: *total, Available: *available}
	}

	o := lsvoucherds.GetPoolAvailabilityRes{PoolAvailability: res}
	return &o, nil
}

func (s *Service) DeletePool(ctx context.Context, in *lsvoucherds.DeletePoolReq) (out *lsvoucherds.DeletePoolRes, err error) {
	defer RecoverAndResponse(ctx, &err)

	region := lscb.Region(in.Region)
	logTags := &lslog.Tags{lslog.Tag_CustomerId: in.CustomerId}

	poolStatus, err := s.Ds.Voucher.Region(region).GetPoolStatus(in.CustomerId, in.PoolId)
	if err != nil {
		lslog.Errorf(ctx, err, "failed to get pool status.").WithTags(logTags).Write()
		return nil, lserr.WrapErrf(err, "failed to get pool status.")
	}

	for uploadId := range *poolStatus {
		err := s.Ds.Voucher.Region(region).DeletePoolUpload(in.CustomerId, in.PoolId, uploadId)
		if err != nil {
			lslog.Errorf(ctx, err, "failed to delete pool upload.").WithTags(logTags).Write()
			return nil, lserr.WrapErrf(err, "failed to delete pool upload.")
		}

		err = s.Ds.Voucher.Region(region).DeleteUsedPoolUpload(in.CustomerId, in.PoolId, uploadId)
		if err != nil {
			lslog.Errorf(ctx, err, "failed to delete used pool upload.").WithTags(logTags).Write()
			return nil, lserr.WrapErrf(err, "failed to used delete pool upload.")
		}
	}

	err = s.Ds.Voucher.Region(region).DeletePool(in.CustomerId, in.PoolId)
	if err != nil {
		lslog.Errorf(ctx, err, "failed to delete pool.").WithTags(logTags).Write()
		return nil, lserr.WrapErrf(err, "failed to delete pool.")
	}

	o := lsvoucherds.DeletePoolRes{}
	return &o, nil
}

func (s *Service) PopFromPool(ctx context.Context, in *lsvoucherds.PopFromPoolReq) (out *lsvoucherds.PopFromPoolRes, err error) {
	defer RecoverAndResponse(ctx, &err)

	region := lscb.Region(in.Region)
	logTags := &lslog.Tags{lslog.Tag_CustomerId: in.CustomerId}

	pool, err := s.Ds.Voucher.Region(region).GetPool(in.CustomerId, in.PoolId)
	if err != nil {
		lslog.Errorf(ctx, err, "failed to get pool.").WithTags(logTags).Write()
		return nil, lserr.WrapErrf(err, "failed to get pool.")
	}

	// check that pool has vouchers available
	if pool.Available <= 0 {
		lslog.Errorf(ctx, err, "pool is empty.").WithTags(logTags).Write()
		return nil, lserr.NewErrf("pool is empty.")
	}

	// get pool with upload id with available vouchers
	var poolUploadId *string
	for uploadId, status := range pool.Status {
		if status.Available > 0 {
			poolUploadId = &uploadId
			break
		}
	}

	if poolUploadId == nil {
		lslog.Errorf(ctx, err, "failed to get pool upload id with available vouchers.").WithTags(logTags).Write()
		return nil, lserr.NewErrf("failed to get pool upload id with available vouchers.")
	}

	// pop voucher from list
	voucher, err := s.Ds.Voucher.Region(region).PopFromPoolUpload(in.CustomerId, in.PoolId, *poolUploadId)
	if err != nil {
		lslog.Errorf(ctx, err, "failed to pop voucher from  pool upload.").WithTags(logTags).Write()
		return nil, lserr.WrapErrf(err, "failed to pop voucher from  pool upload.")
	}

	// push to used pool upload list
	err = s.Ds.Voucher.Region(region).PushToUsedPoolUpload(in.CustomerId, in.PoolId, *poolUploadId, *voucher)
	if err != nil {
		lslog.Errorf(ctx, err, "failed to add voucher to used list.").WithTags(logTags).Write()
		return nil, lserr.WrapErrf(err, "failed to add voucher to used list.")
	}

	o := lsvoucherds.PopFromPoolRes{Voucher: *voucher}
	return &o, nil
}
