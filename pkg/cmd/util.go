package cmd

import (
	"git.containerum.net/ch/kube-client/pkg/model"
	"github.com/go-resty/resty"
)

func resourceError(resp *resty.Response) error {
	return resp.Error().(*model.ResourceError)
}

func firstNonNilErr(err error, errs ...error) error {
	if err != nil {
		return err
	}
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}
