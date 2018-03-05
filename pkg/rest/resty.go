package rest

import (
	"net/http"
	"strings"

	"git.containerum.net/ch/kube-client/pkg/cherry"
	resty "github.com/go-resty/resty"
)

var (
	_ REST = &Resty{}
)

// Resty -- resty client,
// implements REST interface
type Resty struct {
	request *resty.Request
}

func NewResty() *Resty {
	return &Resty{
		request: resty.R(),
	}
}

// Get -- http get method
func (re *Resty) Get(body interface{}, params P, path ...string) error {
	resp, err := re.request.
		SetBody(body).
		SetError(cherry.Err{}).
		SetPathParams(params).
		Get(strings.Join(path, ""))
	if err = MapErrors(resp, err, http.StatusOK); err != nil {
		return err
	}
	copyInterface(body, resp.Result())
	return nil
}

// Put -- http put method
func (re *Resty) Put(body interface{}, params P, path ...string) error {
	resp, err := re.request.
		SetBody(body).
		SetError(cherry.Err{}).
		SetPathParams(params).
		Put(strings.Join(path, ""))
	if err = MapErrors(resp, err,
		http.StatusOK,
		http.StatusAccepted); err != nil {
		return err
	}
	copyInterface(body, resp.Result())
	return nil
}

// Post -- http post method
func (re *Resty) Post(body interface{}, params P, path ...string) error {
	resp, err := re.request.
		SetBody(body).
		SetError(cherry.Err{}).
		SetPathParams(params).
		Post(strings.Join(path, ""))
	if err = MapErrors(resp, err,
		http.StatusOK,
		http.StatusAccepted); err != nil {
		return err
	}
	copyInterface(body, resp.Result())
	return nil
}

// Delete -- http delete method
func (re *Resty) Delete(params P, path ...string) error {
	resp, err := re.request.
		SetError(cherry.Err{}).
		SetPathParams(params).
		Post(strings.Join(path, ""))
	return MapErrors(resp, err,
		http.StatusOK,
		http.StatusAccepted)
}
