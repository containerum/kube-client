package rest

import (
	"fmt"
	"net/http"
	"reflect"

	"git.containerum.net/ch/kube-client/pkg/cherry"
	"github.com/go-resty/resty"
)

// UnexpectedHTTPstatusError -- contains HTTP status code and message
type UnexpectedHTTPstatusError struct {
	Status  int
	Message string
}

func (err *UnexpectedHTTPstatusError) Error() string {
	return fmt.Sprintf("unexpected status [HTTP %d %s] %q",
		err.Status, http.StatusText(err.Status), err.Message)
}

// MapErrors -- trys to extract errors from resty response,
// check http statuses and pack resulting info to error
func MapErrors(resp *resty.Response, err error, okCodes ...int) error {
	if err != nil {
		return err
	}
	for _, code := range okCodes {
		if resp.StatusCode() == code {
			return nil
		}
	}
	if resp.Error() != nil {
		if err, ok := resp.Error().(*cherry.Err); ok &&
			err != nil &&
			err.ID != (cherry.ErrID{}) {
			return err
		}
	}
	return &UnexpectedHTTPstatusError{
		Status:  resp.StatusCode(),
		Message: string(resp.Body()),
	}
}

func copyInterface(dst, src interface{}) {
	if src == nil || dst == nil {
		return
	}
	value := reflect.ValueOf(dst).Elem()
	srcValue := reflect.ValueOf(src)
	if srcValue.Kind() == reflect.Ptr {
		srcValue = srcValue.Elem()
	}
	if value.CanSet() {
		value.Set(srcValue)
	} else {
		panic(fmt.Sprintf("[rest] can't set %v value", value.Type()))
	}
}
