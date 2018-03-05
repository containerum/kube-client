package rest

// P -- URL path params
type P map[string]string
type Q map[string]string

// REST -- rest client interface

type Rq struct {
	Result interface{}
	Body   interface{}
	Params P
	Query  Q
	Path   string
}

type REST interface {
	Get(Rq) error
	Put(Rq) error
	Post(Rq) error
	Delete(Rq) error
}
