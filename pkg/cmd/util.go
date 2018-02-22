package cmd

import (
	"fmt"

	"git.containerum.net/ch/kube-client/pkg/cherry"
	"github.com/go-resty/resty"
)

func catchErr(err error, resp *resty.Response, okCodes ...int) error {
	if err != nil {
		return err
	}
	for _, code := range okCodes {
		if resp.StatusCode() == code {
			return nil
		}
	}
	if resp.Error() != nil {
		err, ok := resp.Error().(*cherry.Err)
		if !ok {
			return fmt.Errorf("%v", resp.Error())
		}
		err.StatusHTTP = resp.StatusCode()
		return err
	}
	return fmt.Errorf("%s", resp.Status())
}
