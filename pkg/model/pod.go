package model

// Pod --
type Pod struct {
	Name            string             `json:"name"`
	Containers      []Container        `json:"containers"`
	ImagePullSecret *map[string]string `json:"image_pull_secret,omitempty"`
	Status          *PodStatus         `json:"status,omitempty"`
	Hostname        *string            `json:"hostname,omitempty"`
	Deploy          *string            `json:"deploy,omitempty"`
	//total CPU usage by all containers in this pod
	TotalCPU uint `json:"total_cpu,omitempty"`
	//total RAM usage by all containers in this pod
	TotalMemory uint `json:"total_memory,omitempty"`
	//creation date in RFC3339 format
	CreatedAt *string `json:"created_at,omitempty"`
}

// PodStatus --
type PodStatus struct {
	Phase        string `json:"phase"`
	RestartCount int    `json:"restart_count"`
	StartAt      string `json:"start_at"`
}

// UpdateImage --
type UpdateImage struct {
	Container string `json:"container_name"`
	Image     string `json:"image"`
}
