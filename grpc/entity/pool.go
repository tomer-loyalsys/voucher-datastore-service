package lsvoucherds_entity

import "github.com/loyalsys/time"

const (
	TypePool = "pool"
)

type PoolStatus map[string]UploadStatus
type Pool struct {
	DocType    string     `json:"type"`
	CustomerId string     `json:"customerId"`
	Total      int64      `json:"total"`
	Available  int64      `json:"available"`
	Status     PoolStatus `json:"status"`
	CreatedAt  int64      `json:"createdAt"`
}

type UploadStatus struct {
	Total     int64 `json:"total"`
	Available int64 `json:"available"`
}

func NewPool(customerId string, total, available int64, poolStatus PoolStatus) Pool {
	return Pool{
		DocType:    TypePool,
		CustomerId: customerId,
		Total:      total,
		Available:  available,
		Status:     poolStatus,
		CreatedAt:  lstime.Millis.NowUnix(),
	}
}
