package rest

// P -- URL path params
type P map[string]string

// REST -- rest client interface
type REST interface {
	Get(interface{}, P, ...string) error
	Put(interface{}, P, ...string) error
	Post(interface{}, P, ...string) error
	Delete(P, ...string) error
}
