package model

type DeploymentStatus struct {
	CreatedAt           int64 `json:"created_at"`
	UpdatedAt           int64 `json:"updated_at"`
	Replicas            int   `json:"replicas"`
	ReadyReplicas       int   `json:"ready_replicas"`
	AvailableReplicas   int   `json:"available_replicas"`
	UnavailableReplicas int   `json:"unavailable_replicas"`
	UpdatedReplicas     int   `json:"updated_replicas"`
}

type UpdateReplicas struct {
	Replicas int `json:"replicas" binding:"required"`
}

type Deployment struct {
	Status     *DeploymentStatus `json:"status,omitempty"` //
	Containers []Container       `json:"containers"`       //
	Labels     map[string]string `json:"labels,omitempty"` //
	Name       string            `json:"name"`             // not UUID!
	Replicas   int               `json:"replicas"`         //
}

type Container struct {
	Image     string `json:"image"` //
	Name      string `json:"name"`  // not UUID!
	Resources struct {
		Requests Resource `json:"requests"` //
	} `json:"resources"` //
	Env          []Env    `json:"env, omitempty"`       //
	Commands     []string `json:"commands"`             //
	Ports        []Port   `json:"ports"`                //
	VolumeMounts []Volume `json:"volume_mounts"`        //
	ConfigMap    []Volume `json:"config_map,omitempty"` //
}

type Env struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

type Port struct {
	ContainerPort int `json:"container_port"`
}
