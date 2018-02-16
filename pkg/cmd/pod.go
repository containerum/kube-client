package cmd

import (
	"fmt"
	"net/http"
)

const (
	kubeAPIpodPath = "/namespaces/{namespace}/pods/{pod}"
)

// DeletePod -- deletes pod in provided namespace
func (client *Client) DeletePod(namespace, pod string) error {
	resp, err := client.Request.
		SetPathParams(map[string]string{
			"pod": pod,
		}).
		Delete(client.APIurl + kubeAPIpodPath)
	if err != nil {
		return err
	}
	switch resp.StatusCode() {
	case http.StatusOK, http.StatusAccepted:
		return nil
	default:
		return fmt.Errorf("%s", string(resp.Body()))
	}
}
