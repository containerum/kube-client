package cmd

import (
	"git.containerum.net/ch/kube-client/pkg/model"
)

const (
	getService = "/namespaces/{namespace}/services/{service}"
)

// GetService -- consume namespace id and srvice name
// returns Service or uninitialized Service struct AND error
func (client *Client) GetService(namespace, serviceName string) (model.Service, error) {
	resp, err := client.Request.
		SetResult(model.Service{}).
		SetPathParams(map[string]string{
			"namespace": namespace,
			"service":   serviceName,
		}).
		Get(client.serverURL + getService)
	if err != nil {
		return model.Service{}, err
	}
	return *resp.Result().(*model.Service), nil
}
