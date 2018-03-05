package rest

type P map[string]string

type REST interface {
	Get(interface{}, P, ...string) error
	Put(interface{}, P, ...string) error
	Post(interface{}, P, ...string) error
	Delete(P, ...string) error
}
