package client

import (
	"github.com/containerum/kube-client/pkg/model"
	"github.com/containerum/kube-client/pkg/rest"
)

const (
	accessesPath = "/projects/{project}/namespaces/{namespace}/access"
)

func (client *Client) GetNamespaceAccesses(project, namespace string) ([]model.UserAccess, error) {
	var access model.Namespace
	err := client.RestAPI.Get(rest.Rq{
		Result: &access,
		URL: rest.URL{
			Path: accessesPath,
			Params: rest.P{
				"namespace": namespace,
				"project":   project,
			},
		},
	})
	return access.Users, err
}

// SetNamespaceAccess -- sets/changes access to namespace for provided user
func (client *Client) SetNamespaceAccess(project, namespace, username string, access model.UserGroupAccess) error {
	return client.RestAPI.Put(rest.Rq{
		Body: model.ResourceUpdateUserAccess{
			Username: username,
			Access:   access,
		},
		URL: rest.URL{
			Path: namespaceAccessPath,
			Params: rest.P{
				"namespace": namespace,
				"project":   project,
			},
		},
	})
}

// DeleteNamespaceAccess -- deletes user access to namespace
func (client *Client) DeleteNamespaceAccess(project, namespace, username string) error {
	return client.RestAPI.Delete(rest.Rq{
		Body: model.ResourceUpdateUserAccess{
			Username: username,
		},
		URL: rest.URL{
			Path: namespaceAccessPath,
			Params: rest.P{
				"namespace": namespace,
				"project":   project,
			},
		},
	})
}
