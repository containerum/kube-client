package rest

// P -- URL path params
type P map[string]string

// REST -- rest client interface
type REST interface {
	Get(result interface{}, params P, path ...string) error
	Put(result, body interface{}, params P, path ...string) error
	Post(result, body interface{}, params P, path ...string) error
	Delete(params P, path ...string) error
}
