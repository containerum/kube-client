package client

import (
	"github.com/containerum/kube-client/pkg/rest"

	"github.com/containerum/kube-client/pkg/model"
)

const (
	deploymentPodsPath = "/projects/{project}/namespaces/{namespace}/deployments/{deployment}/pods"
	podsPath           = "/projects/{project}/namespaces/{namespace}/pods"
	podPath            = "/projects/{project}/namespaces/{namespace}/pods/{pod}"
)

// DeletePod -- deletes pod in provided namespace
func (client *Client) DeletePod(project, namespace, pod string) error {
	return client.RestAPI.Delete(rest.Rq{
		URL: rest.URL{
			Path: podPath,
			Params: rest.P{
				"pod":       pod,
				"namespace": namespace,
				"project":   project,
			},
		},
	})
}

// GetPod -- gets a particular pod by name.
func (client *Client) GetPod(project, namespace, pod string) (model.Pod, error) {
	var gainedPod model.Pod
	err := client.RestAPI.Get(rest.Rq{
		Result: &gainedPod,
		URL: rest.URL{
			Path: podPath,
			Params: rest.P{
				"namespace": namespace,
				"pod":       pod,
				"project":   project,
			},
		},
	})
	return gainedPod, err
}

// GetPodList -- returns list of pods in provided namespace
func (client *Client) GetPodList(project, namespace string) (model.PodsList, error) {
	var podList model.PodsList
	err := client.RestAPI.Get(rest.Rq{
		Result: &podList,
		URL: rest.URL{
			Path: podsPath,
			Params: rest.P{
				"namespace": namespace,
				"project":   project,
			},
		},
	})
	return podList, err
}

// GetDeploymentPodList -- returns list of pods in provided namespace and deployment
func (client *Client) GetDeploymentPodList(namespace, project, deployment string) (model.PodsList, error) {
	var podList model.PodsList
	err := client.RestAPI.Get(rest.Rq{
		Result: &podList,
		URL: rest.URL{
			Path: deploymentPodsPath,
			Params: rest.P{
				"namespace":  namespace,
				"project":   project,
				"deployment": deployment,
			},
		},
	})
	return podList, err
}
