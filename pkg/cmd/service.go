package cmd

import (
	"git.containerum.net/ch/kube-client/pkg/model"
)

const (
	servicePath  = "/namespaces/{namespace}/services/{service}"
	servicesPath = "/namespaces/{namespace}/services"
)

// GetService -- consume namespace id and srvice name
// returns Service OR uninitialized Service struct AND error
func (client *Client) GetService(namespace, serviceName string) (model.Service, error) {
	resp, err := client.Request.
		SetResult(model.Service{}).
		SetPathParams(map[string]string{
			"namespace": namespace,
			"service":   serviceName,
		}).
		Get(client.serverURL + servicePath)
	if err != nil {
		return model.Service{}, err
	}

	return *resp.Result().(*model.Service), nil
}

// GetServiceList -- consume namespace name
// returns slice of Services OR nil slice AND error
func (client *Client) GetServiceList(namespace string) ([]model.Service, error) {
	resp, err := client.Request.
		SetResult([]model.Service{}).
		SetPathParams(map[string]string{
			"namespace": namespace,
		}).
		Get(client.serverURL + servicesPath)
	if err != nil {
		return nil, err
	}

	return *resp.Result().(*[]model.Service), nil
}

func (client *Client) CreateService(namespace string, service model.Service) (model.Service, error) {
	resp, err := client.Request.
		SetResult(model.Service{}).
		SetBody(service).
		SetPathParams(map[string]string{
			"namespace": namespace,
		}).Post(client.serverURL + servicesPath)
	if err != nil {
		return model.Service{}, err
	}
	return *resp.Result().(*model.Service), nil
}
