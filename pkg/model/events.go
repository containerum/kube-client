package model

type EventKind string

const (
	EventError   EventKind = "error"
	EventWarning EventKind = "warning"
	EventInfo    EventKind = "info"
)

type ResourceType string

const (
	TypeNamespace  ResourceType = "namespace"
	TypeDeployment ResourceType = "deployment"
	TypePod        ResourceType = "pod"
	TypeService    ResourceType = "service"
	TypeIngress    ResourceType = "ingress"
	TypeVolume     ResourceType = "volume"
	TypeStorage    ResourceType = "storage"
	TypeConfigMap  ResourceType = "configmap"
	TypeSecret     ResourceType = "secret"
	TypeNode       ResourceType = "node"
	TypeUser       ResourceType = "user"
	TypeSystem     ResourceType = "system"
)

const (
	ResourceCreated  string = "ResourceCreated"
	ResourceModified string = "ResourceModified"
	ResourceDeleted  string = "ResourceDeleted"
)

type EventsList struct {
	Events []Event `json:"events" yaml:"events" bson:"events"`
}

// Event -- Containerum event
//
// swagger:model
type Event struct {
	Kind              EventKind    `json:"event_kind" yaml:"event_kind" bson:"eventkind"`
	Time              string       `json:"event_time" yaml:"event_time" bson:"eventtime"`
	Name              string       `json:"event_name" yaml:"event_name" bson:"eventname"`
	ResourceType      ResourceType `json:"resource_type" yaml:"resource_type" bson:"resourcetype"`
	ResourceName      string       `json:"resource_name" yaml:"resource_name" bson:"resourcename"`
	ResourceNamespace string       `json:"resource_namespace,omitempty" yaml:"resource_namespace,omitempty" bson:"resourcenamespace,omitempty"`
	Message           string       `json:"message,omitempty" yaml:"message,omitempty" bson:"message,omitempty"`
}
