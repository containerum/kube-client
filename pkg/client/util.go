package client

import (
	"fmt"

	"git.containerum.net/ch/kube-client/pkg/cherry"
	"github.com/go-resty/resty"
)

// UnexpectedHTTPstatusError -- contains HTTP status code and message
type UnexpectedHTTPstatusError struct {
	Code int
	Msg  string
}

func (err *UnexpectedHTTPstatusError) Error() string {
	return "unexpected status: " + err.Msg
}

func MapErrors(resp *resty.Response, err error, okCodes ...int) error {
	if err != nil {
		return err
	}
	for _, code := range okCodes {
		if resp.StatusCode() == code && resp.Error() != nil {
			return nil
		}
	}
	if resp.Error() != nil {
		if err, ok := resp.Error().(*cherry.Err); ok {
			return err
		}
		return fmt.Errorf("%q", string(resp.Body()))
	}
	return &UnexpectedHTTPstatusError{
		Code: resp.StatusCode(),
		Msg:  resp.Status(),
	}
}
