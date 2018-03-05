package rest

import "strings"

// P -- URL path params
type P map[string]string

// Q -- URL query params
type Q map[string]string

type URL struct {
	Templ  string
	Params P
}

func (u *URL) compile() string {
	addr := u.Templ
	for k, v := range u.Params {
		addr = strings.Replace(addr,
			"{"+k+"}", v, -1)
	}
	return addr
}

// Rq -- request params
type Rq struct {
	Result interface{}
	Body   interface{}
	Path   URL
	Query  Q
}

// REST -- rest client interface
type REST interface {
	Get(Rq) error
	Put(Rq) error
	Post(Rq) error
	Delete(Rq) error
}
