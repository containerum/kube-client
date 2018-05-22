package client

import (
	"github.com/containerum/kube-client/pkg/model"
	"github.com/containerum/kube-client/pkg/rest"
)

const (
	permissionPath = "/namespaces/{namespace}/access"
)

func (client *Client) GetPermission(namespace string) (model.Permission, error) {
	var perm model.Permission
	err := client.RestAPI.Get(rest.Rq{
		Result: &perm,
		URL: rest.URL{
			Path: kubeAPIIngressRootPath,
			Params: rest.P{
				"namespace": namespace,
			},
		},
	})
	return perm, err
}
