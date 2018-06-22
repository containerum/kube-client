package client

import (
	"github.com/containerum/kube-client/pkg/model"
	"github.com/containerum/kube-client/pkg/rest"
)

const (
	ingressesPath = "/projects/{project}/namespaces/{namespace}/ingresses"
	ingressPath   = "/projects/{project}/namespaces/{namespace}/ingresses/{domain}"
)

// AddIngress -- adds ingress to provided namespace
func (client *Client) AddIngress(project, namespace string, ingress model.Ingress) error {
	return client.RestAPI.Post(rest.Rq{
		Body: ingress,
		URL: rest.URL{
			Path: ingressesPath,
			Params: rest.P{
				"namespace": namespace,
				"project":   project,
			},
		},
	})
}

// GetIngressList -- returns list of ingresses.
func (client *Client) GetIngressList(project, namespace string) (model.IngressesList, error) {
	var ingressList model.IngressesList
	err := client.RestAPI.Get(rest.Rq{
		Result: &ingressList,
		URL: rest.URL{
			Path: ingressesPath,
			Params: rest.P{
				"namespace": namespace,
				"project":   project,
			},
		},
	})
	return ingressList, err
}

// GetIngressList -- returns ingress with specified domain.
func (client *Client) GetIngress(project, namespace, domain string) (model.Ingress, error) {
	var ingress model.Ingress
	err := client.RestAPI.Get(rest.Rq{
		Result: &ingress,
		URL: rest.URL{
			Path: ingressPath,
			Params: rest.P{
				"namespace": namespace,
				"domain":    domain,
				"project":   project,
			},
		},
	})
	return ingress, err
}

// UpdateIngress -- updates ingress on provided domain with new one
func (client *Client) UpdateIngress(project, namespace, domain string, ingress model.Ingress) error {
	return client.RestAPI.Put(rest.Rq{
		Body: ingress,
		URL: rest.URL{
			Path: ingressPath,
			Params: rest.P{
				"namespace": namespace,
				"domain":    domain,
				"project":   project,
			},
		},
	})
}

// DeleteIngress -- deletes ingress on provided domain
func (client *Client) DeleteIngress(project, namespace, domain string) error {
	return client.RestAPI.Delete(rest.Rq{
		URL: rest.URL{
			Path: ingressPath,
			Params: rest.P{
				"namespace": namespace,
				"domain":    domain,
				"project":   project,
			},
		},
	})
}
