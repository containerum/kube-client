package cmd

import (
	"fmt"
	"net/http"

	"git.containerum.net/ch/kube-client/pkg/model"
)

const (
	resourceIngressRootPath = "/namespace/{namespace}/ingress"
)

// AddIngress -- adds ingress to provided namespace
func (client *Client) AddIngress(namespace string, ingress model.ResourceIngress) error {
	resp, err := client.Request.
		SetPathParams(map[string]string{
			"namespace": namespace,
		}).SetBody(ingress).
		Post(client.resourceServiceAddr + resourceIngressRootPath)
	if err != nil {
		return err
	}
	switch resp.StatusCode() {
	case http.StatusOK, http.StatusAccepted:
		return nil
	default:
		if resp.Error() != nil {
			return fmt.Errorf("%v", resp.Error())
		}
		return fmt.Errorf("%s", resp.Status())
	}
}
