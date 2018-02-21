package model

import "time"

type Resources struct {
	Hard Resource  `json:"hard"`           //
	Used *Resource `json:"used,omitempty"` //
}

type Resource struct {
	CPU    string `json:"cpu"`    //
	Memory string `json:"memory"` //
}

type UpdateNamespaceName struct {
	Label string `json:"label"` //
}

// ResourceNamespace -- namespace representation
// provided by resource-service
// https://ch.pages.containerum.net/api-docs/modules/resource-service/index.html#get-namespace
type Namespace struct {
	CreateTime    time.Time        `json:"create_time"`     //
	TariffID      string           `json:"tariff_id"`       //
	Label         string           `json:"label"`           //
	Access        string           `json:"access"`          //
	MaxExtService uint             `json:"max_ext_service"` //
	MaxIntService uint             `json:"max_int_service"` //
	MaxTraffic    uint             `json:"max_traffic"`     //
	Volumes       []ResourceVolume `json:"volumes"`         //
	Resources     Resources        `json:"resources"`       //
}
