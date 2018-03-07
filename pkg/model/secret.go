package model

type SecretsList struct {
	Secrets []Secret `json:"secrets"`
}

// Secret --
type Secret struct {
	Name      string            `json:"name"`
	CreatedAt *int64            `json:"created_at,omitempty"`
	Data      map[string]string `json:"data"`
}
