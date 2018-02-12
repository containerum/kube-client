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
func (c *Client) GetNamespaceList(queries map[string]string) ([]model.Namespace, error) {
	resp, err := c.Request.
		SetQueryParams(queries).
		SetResult([]model.Namespace{}).
		Get(c.serverURL + getNamespaceList)
	if err != nil {
		return []model.Namespace{}, err
	}
	return *resp.Result().(*[]model.Namespace), nil
}

//GetNamespace return namespace by Name
func (c *Client) GetNamespace(ns string) (model.Namespace, error) {
	resp, err := c.Request.SetResult(model.Namespace{}).
		SetPathParams(map[string]string{
			"namespace": ns,
		}).
		Get(c.serverURL + getNamespace)
	if err != nil {
		return model.Namespace{}, err
	}
	return *resp.Result().(*model.Namespace), nil
}

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
