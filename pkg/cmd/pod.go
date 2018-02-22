package cmd

import (
	"net/http"

	"git.containerum.net/ch/kube-client/pkg/model"
)

const (
	kubeAPIpodRootPath = "/namespaces/{namespace}/pods"
	kubeAPIpodPath     = "/namespaces/{namespace}/pods/{pod}"
)

// DeletePod -- deletes pod in provided namespace
func (client *Client) DeletePod(namespace, pod string) error {
	resp, err := client.Request.
		SetPathParams(map[string]string{
			"pod": pod,
		}).
		Delete(client.APIurl + kubeAPIpodPath)
	return mapErrors(resp, err,
		http.StatusOK,
		http.StatusAccepted)
}

// GetPod -- gets a particular pod by name.
func (client *Client) GetPod(namespace, pod string) (model.Pod, error) {
	resp, err := client.Request.
		SetPathParams(map[string]string{
			"namespace": namespace,
			"pod":       pod,
		}).
		Get(client.APIurl + kubeAPIpodPath)
	if err = mapErrors(resp, err, http.StatusOK, http.StatusAccepted); err != nil {
		return model.Pod{}, err
	}
	return *resp.Result().(*model.Pod), nil
}

func (client *Client) GetPodList(namespace string) ([]model.Pod, error) {
	resp, err := client.Request.
		SetPathParams(map[string]string{
			"namespace": namespace,
		}).
		Get(client.APIurl + kubeAPIpodRootPath)
	if err = mapErrors(resp, err, http.StatusOK, http.StatusAccepted); err != nil {
		return nil, err
	}
	return *resp.Result().(*[]model.Pod), nil
}
