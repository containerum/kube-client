package model

type Secret struct {
	Created int64             `json:"created_at,omitempty"`
	Data    map[string]string `json:"data" binding:"required"`
	Name    string            `json:"name" binding:"required"`
}
