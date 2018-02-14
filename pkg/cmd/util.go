package cmd

import (
	"git.containerum.net/ch/kube-client/pkg/model"
	"github.com/go-resty/resty"
)

func resourceError(resp *resty.Response) error {
	return resp.Error().(*model.ResourceError)
}
