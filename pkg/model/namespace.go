package model

// represents namespace limits and user resources,
//
// swagger:model
type Resources struct {
	// Hard resource limits
	//
	// required: true
	Hard Resource  `json:"hard"`
	Used *Resource `json:"used,omitempty"`
}

// represents namespace CPU and RAM
//
// swagger:model
type Resource struct {
	// CPU in m
	//
	// required: true
	CPU uint `json:"cpu"`
	// RAM in Mi
	//
	// required: true
	Memory uint `json:"memory"`
}

// contains new namespace name
//
// swagger:model
type UpdateNamespaceName struct {
	// required: true
	Label string `json:"label"`
}

// namespace representation provided by resource-service
// https://ch.pages.containerum.net/api-docs/modules/resource-service/index.html#get-namespace
//
// swagger:model
type Namespace struct {
	CreatedAt     *string   `json:"created_at,omitempty"`
	Label         string    `json:"label,omitempty"`
	Access        string    `json:"access,omitempty"`
	MaxExtService *uint     `json:"max_ext_service,omitempty"`
	MaxIntService *uint     `json:"max_int_service,omitempty"`
	MaxTraffic    *uint     `json:"max_traffic,omitempty"`
	Volumes       []Volume  `json:"volumes,omitempty"`
	Resources     Resources `json:"resources"`
}
