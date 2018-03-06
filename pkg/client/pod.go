package client

import (
	"git.containerum.net/ch/kube-client/pkg/rest"

	"git.containerum.net/ch/kube-client/pkg/model"
)

const (
	kubeAPIpodRootPath = "/namespaces/{namespace}/pods"
	kubeAPIpodPath     = "/namespaces/{namespace}/pods/{pod}"
)

// DeletePod -- deletes pod in provided namespace
func (client *Client) DeletePod(namespace, pod string) error {
	return client.re.Delete(rest.Rq{
		Path: rest.URL{
			Templ: client.APIurl + kubeAPIpodPath,
			Params: rest.P{
				"pod": pod,
			},
		},
	})
}

// GetPod -- gets a particular pod by name.
func (client *Client) GetPod(namespace, pod string) (model.Pod, error) {
	var gainedPod model.Pod
	err := client.re.Get(rest.Rq{
		Result: &gainedPod,
		Path: rest.URL{
			Templ: client.APIurl + kubeAPIpodPath,
			Params: rest.P{
				"namespace": namespace,
				"pod":       pod,
			},
		},
	})
	return gainedPod, err
}

// GetPodList -- returns list of pods in provided namespace
func (client *Client) GetPodList(namespace string) ([]model.Pod, error) {
	var podList []model.Pod
	err := client.re.Get(rest.Rq{
		Result: &podList,
		Path: rest.URL{
			Templ: client.APIurl + kubeAPIpodRootPath,
			Params: rest.P{
				"namespace": namespace,
			},
		},
	})
	return podList, err
}
