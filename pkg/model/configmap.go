package model

type ConfigMapsList struct {
	ConfigMaps []ConfigMap `json:"configmaps"`
}

// ConfigMap --
type ConfigMap struct {
	Name      string            `json:"name"`
	CreatedAt *int64            `json:"created_at,omitempty"`
	Data      map[string]string `json:"data"`
}
