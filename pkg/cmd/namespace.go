package cmd

import (
	"git.containerum.net/ch/kube-client/pkg/model"
)

//ListOptions -
type ListOptions struct {
	Owner string
}

const (
	getNamespace         = "/namespaces/{namespace}"
	getNamespaceList     = "/namespaces"
	serviceNamespacePath = "/namespace/{namespace}"
)

//GetNamespaceList return namespace list. Can use query filters: owner
func (client *Client) GetNamespaceList(queries map[string]string) ([]model.Namespace, error) {
	resp, err := client.Request.
		SetQueryParams(queries).
		SetResult([]model.Namespace{}).
		Get(client.serverURL + getNamespaceList)
	if err != nil {
		return []model.Namespace{}, err
	}
	return *resp.Result().(*[]model.Namespace), nil
}

//GetNamespace return namespace by Name
func (client *Client) GetNamespace(ns string) (model.Namespace, error) {
	resp, err := client.Request.SetResult(model.Namespace{}).
		SetPathParams(map[string]string{
			"namespace": ns,
		}).
		Get(client.serverURL + getNamespace)
	if err != nil {
		return model.Namespace{}, err
	}
	return *resp.Result().(*model.Namespace), nil
}

// ResourceGetNamespace -- consumes a namespace and an optional user ID
// returns a namespace data OR an error
func (client *Client) ResourceGetNamespace(namespace, userID string) (model.ResourceNamespace, error) {
	req := client.Request.
		SetPathParams(map[string]string{
			"namespace": namespace,
		}).SetResult(model.ResourceNamespace{})
	if userID != "" {
		req.SetQueryParam("user-id", userID)
	}
	resp, err := req.Get(client.resourceServiceAddr + serviceNamespacePath)
	if err != nil {
		return model.ResourceNamespace{}, nil
	}
	return *resp.Result().(*model.ResourceNamespace), nil
}
