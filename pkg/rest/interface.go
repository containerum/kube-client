package rest

// P -- URL path params
type P map[string]string

// Q -- URL query params
type Q map[string]string

// Rq -- request params
type Rq struct {
	Result interface{}
	Body   interface{}
	Params P
	Query  Q
	Path   string
}

// REST -- rest client interface
type REST interface {
	Get(Rq) error
	Put(Rq) error
	Post(Rq) error
	Delete(Rq) error
}
