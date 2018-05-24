package model

import "time"

// Volume -- volume representation
//
//swagger:model
type Volume struct {
	ID               string    `json:"id"`
	CreateTime       time.Time `json:"create_time"`
	Label            string    `json:"label"`
	Access           string    `json:"access"`
	AccessChangeTime time.Time `json:"access_change_time"`
	Capacity         int       `json:"capacity"`
	Replicas         int       `json:"replicas"`
}

// CreateVolume --
//swagger:ignore
type CreateVolume struct {
	TariffID string `json:"tariff-id"`
	Label    string `json:"label"`
}

// ResourceUpdateName -- contains new resource name
//swagger:ignore
type ResourceUpdateName struct {
	Label string `json:"label"`
}
