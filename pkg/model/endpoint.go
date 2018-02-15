package model

type Endpoint struct {
	Name      string   `json:"name" binding:"required"`
	Owner     *string  `json:"owner,omitempty" binding:"required"`
	CreatedAt *int64   `json:"created_at,omitempty"`
	Addresses []string `json:"addresses" binding:"required"`
	Ports     []Port   `json:"ports" binding:"required"`
}
