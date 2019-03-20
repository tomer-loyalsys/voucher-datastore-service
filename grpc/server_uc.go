// Copyright 2018 The Loyalsys Authors. All rights reserved.

package grpc

import (
	"context"
	"github.com/loyalsys/couchbase"
	"github.com/loyalsys/error"
	"github.com/loyalsys/log"
	"github.com/loyalsys/voucher-datastore-service/grpc/entity"
)

func (s *Service) addToPool(ctx context.Context, region, customerId, poolId, uploadId string) (*lsvoucherds_entity.UploadStatus, error) {
	logTags := &lslog.Tags{lslog.Tag_CustomerId: customerId}
	r := lscb.Region(region)

	// get pool upload total size
	uploadSize, err := s.Ds.Voucher.Region(r).GetPoolUploadSize(customerId, poolId, uploadId)
	if err != nil {
		lslog.Errorf(ctx, err, "failed to get pool upload size.").WithTags(logTags).Write()
		return nil, lserr.WrapErrf(err, "failed to get pool upload size.")
	}

	// create upload statistics
	uploadStatus := lsvoucherds_entity.UploadStatus{Total: int64(uploadSize), Available: int64(uploadSize)}

	// create pool if not exist
	isPoolExists, err := s.Ds.Voucher.Region(r).IsPoolExists(customerId, poolId)
	if err != nil {
		lslog.Errorf(ctx, err, "failed to check if pool exists.").WithTags(logTags).Write()
		return nil, lserr.WrapErrf(err, "failed to check if pool exists.")
	}

	if *isPoolExists {
		err = s.Ds.Voucher.Region(r).UpsertPool(customerId, poolId, uploadId, uploadStatus)
		if err != nil {
			lslog.Errorf(ctx, err, "failed to upsert pool.").WithTags(logTags).Write()
			return nil, lserr.WrapErrf(err, "failed to upsert pool.")
		}
	} else {
		// create pool
		uploadStats := lsvoucherds_entity.PoolStatus{uploadId: uploadStatus}
		pool := lsvoucherds_entity.NewPool(customerId, uploadStatus.Total, uploadStatus.Total, uploadStats)

		err = s.Ds.Voucher.Region(r).CreatePool(customerId, poolId, pool)
		if err != nil {
			lslog.Errorf(ctx, err, "failed to create pool.").WithTags(logTags).Write()
			return nil, lserr.WrapErrf(err, "failed to create pool.")
		}
	}

	return &uploadStatus, nil
}
