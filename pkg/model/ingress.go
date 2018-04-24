package model

// Ingress
//
// swagger:model
type Ingress struct {
	// required: true
	Name string `json:"name"`
	//creation date in RFC3339 format
	CreatedAt *string `json:"created_at,omitempty"`
	// required: true
	Rules []Rule `json:"rules"`
}

// Rule --
//
// swagger:model
type Rule struct {
	// required: true
	Host      string  `json:"host"`
	TLSSecret *string `json:"tls_secret,omitempty"`
	// required: true
	Path []Path `json:"path"`
}

// Path --
//
// swagger:model
type Path struct {
	// required: true
	Path string `json:"path"`
	// required: true
	ServiceName string `json:"service_name"`
	// required: true
	ServicePort int `json:"service_port"`
}
