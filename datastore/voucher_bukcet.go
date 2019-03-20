// Copyright 2018 The Loyalsys Authors. All rights reserved.

package datastore

import (
	"fmt"
	"github.com/loyalsys/couchbase"
	"github.com/loyalsys/error"
	"github.com/loyalsys/voucher-datastore-service/grpc/entity"
	"gopkg.in/couchbase/gocb.v1"
)

const (
	poolKeyFormat   = "cus::%v::pool::%v"
	uploadKeyFormat = "cus::%v::pool::%v::upload::%v"
)

type voucher struct{ bucket *gocb.Bucket }

func (v *voucher) GetBucketName() lscb.BucketType {
	return lscb.BucketTypeVoucher
}

func (v *voucher) AssignBucket(bucket *gocb.Bucket) {
	v.bucket = bucket
}

func (v *voucher) EnsureIndexes(manager *gocb.BucketManager) error {
	return nil
}

func (v *voucher) Ping() error {
	_, err := v.bucket.Ping([]gocb.ServiceType{gocb.MemdService, gocb.CapiService, gocb.N1qlService, gocb.FtsService})
	if err != nil {
		return lserr.WrapErrf(err, "failed to ping.")
	}

	return nil
}

func (v *voucher) IsPoolExists(customerId string, poolId string) (*bool, error) {
	key := fmt.Sprintf(poolKeyFormat, customerId, poolId)

	_, err := v.bucket.LookupIn(key).Exists("type").Execute()
	if err != nil {
		if gocb.IsKeyNotFoundError(err) {
			isExists := false
			return &isExists, nil
		} else {
			return nil, lserr.WrapErrf(err, "failed to lookup for `type` path in points tracking.")
		}
	}

	isExists := true

	return &isExists, nil
}

func (v *voucher) GetPool(customerId string, poolId string) (*lsvoucherds_entity.Pool, error) {
	key := fmt.Sprintf(poolKeyFormat, customerId, poolId)

	var pool *lsvoucherds_entity.Pool
	_, err := v.bucket.Get(key, &pool)
	if err != nil {
		return nil, lserr.WrapErrf(err, "failed to get pool.")
	}

	return pool, nil
}

func (v *voucher) CreatePool(customerId string, poolId string, pool lsvoucherds_entity.Pool) error {
	key := fmt.Sprintf(poolKeyFormat, customerId, poolId)

	_, err := v.bucket.Insert(key, pool, 0)
	if err != nil {
		return lserr.WrapErrf(err, "failed to create pool.")
	}

	return nil
}

func (v *voucher) GetPoolStatus(customerId string, poolId string) (total *int64, available *int64, err error) {
	key := fmt.Sprintf(poolKeyFormat, customerId, poolId)

	var totalVal int64
	_, err = v.bucket.MapGet(key, "total", &totalVal)
	if err != nil {
		return nil, nil, lserr.WrapErrf(err, "failed to get pool total status.")
	}

	var availableVal int64
	_, err = v.bucket.MapGet(key, "available", &availableVal)
	if err != nil {
		return nil, nil, lserr.WrapErrf(err, "failed to get pool available status.")
	}

	return &totalVal, &availableVal, nil
}

func (v *voucher) UpsertPool(customerId string, poolId string, uploadId string, uploadStatus lsvoucherds_entity.UploadStatus) error {
	key := fmt.Sprintf(poolKeyFormat, customerId, poolId)

	path := fmt.Sprintf("status.%v.", uploadId)
	_, err := v.bucket.MutateIn(key, 0, 0).Upsert(path, uploadStatus, false).
		Counter("total", uploadStatus.Total, false).
		Counter("available", uploadStatus.Total, false).
		Execute()

	if err != nil {
		return lserr.WrapErrf(err, "failed to upsert points tracking.")
	}

	return nil
}

func (v *voucher) DeletePool(customerId string, poolId string) error {
	key := fmt.Sprintf(poolKeyFormat, customerId, poolId)

	_, err := v.bucket.Remove(key, 0)
	if err != nil {
		return lserr.WrapErrf(err, "failed to delete pool.")
	}

	return nil
}

func (v *voucher) IsPoolUploadExists(customerId string, poolId string, uploadId string) (*bool, error) {
	key := fmt.Sprintf(uploadKeyFormat, customerId, poolId, uploadId)

	_, _, err := v.bucket.ListSize(key)
	if err != nil {
		if gocb.IsKeyNotFoundError(err) {
			isExists := false
			return &isExists, nil
		} else {
			return nil, lserr.WrapErrf(err, "failed to check if pool upload exists.")
		}
	}

	isExists := true

	return &isExists, nil
}

func (v *voucher) CreatePoolUpload(customerId string, poolId string, uploadId string) error {
	key := fmt.Sprintf(uploadKeyFormat, customerId, poolId, uploadId)

	_, err := v.bucket.Insert(key, []string{}, 0)
	if err != nil {
		return lserr.WrapErrf(err, "failed to add new voucher to set.")
	}

	return nil
}

func (v *voucher) AppendToPoolUpload(customerId string, poolId string, uploadId string, voucher string) error {
	key := fmt.Sprintf(uploadKeyFormat, customerId, poolId, uploadId)

	_, err := v.bucket.QueuePush(key, voucher, true)
	if err != nil {
		return lserr.WrapErrf(err, "failed to add new voucher to set.")
	}

	return nil
}

func (v *voucher) GetPoolUploadSize(customerId string, poolId string, uploadId string) (uint, error) {
	key := fmt.Sprintf(uploadKeyFormat, customerId, poolId, uploadId)

	size, _, err := v.bucket.QueueSize(key)
	if err != nil {
		return 0, lserr.WrapErrf(err, "failed to get upload vouchers size.")
	}

	return size, nil
}

func (v *voucher) DeletePoolUpload(customerId string, poolId string, uploadId string) error {
	key := fmt.Sprintf(uploadKeyFormat, customerId, poolId, uploadId)

	_, err := v.bucket.Remove(key, 0)
	if err != nil {
		return lserr.WrapErrf(err, "failed to delete pool upload.")
	}

	return nil
}
