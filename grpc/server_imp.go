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

	for {
		in, err := stream.Recv()

		if err == io.EOF {
			uploadStat, err := s.addToPool(stream.Context(), *region, *customerId, *poolId, *uploadId)
			if err != nil {
				lslog.Errorf(stream.Context(), err, "failed to add upload to pool.").WithTags(logTags).Write()
				return lserr.WrapErrf(err, "failed to add upload to pool.")
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
		for _, voucher := range in.Vouchers {
			// add voucher batch to pool upload
			err = s.Ds.Voucher.Region(r).AppendToPoolUpload(in.CustomerId, in.PoolId, in.UploadId, voucher)
			if err != nil {
				lslog.Errorf(stream.Context(), err, "failed to insert voucher to upload doc.").WithTags(logTags).Write()
				return lserr.WrapErrf(err, "failed to insert voucher to upload doc.")
			}
		}
	}
}

func (s *Service) GetPoolStatus(ctx context.Context, in *lsvoucherds.GetPoolStatusReq) (out *lsvoucherds.GetPoolStatusRes, err error) {
	defer RecoverAndResponse(ctx, &err)

	region := lscb.Region(in.Region)
	logTags := &lslog.Tags{lslog.Tag_CustomerId: in.CustomerId}

	res := make(map[string]*lsvoucherds.GetPoolStatusRes_PoolStatus, len(in.PoolIds))
	for _, poolId := range in.PoolIds {
		total, available, err := s.Ds.Voucher.Region(region).GetPoolStatus(in.CustomerId, poolId)
		if err != nil {
			lslog.Errorf(ctx, err, "failed to get pool status.").WithTags(logTags).Write()
			return nil, lserr.WrapErrf(err, "failed to get pool status.")
		}

		res[poolId] = &lsvoucherds.GetPoolStatusRes_PoolStatus{Total: *total, Available: *available}
	}

	o := lsvoucherds.GetPoolStatusRes{PoolsStatus: res}
	return &o, nil
}
