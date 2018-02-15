package cmd

import (
	"net/http"

	"git.containerum.net/ch/kube-client/pkg/model"
)

const (
	kubeAPIdeploymentPath  = "/namespaces/{namespace}/deployments/{deployment}"
	kubeAPIdeploymentsPath = "/namespaces/{namespace}/deployments"

	resourceDeploymentRootPath = "/namespace/{namespace}/deployment"
	resourceDeploymentPath     = resourceDeploymentRootPath + "/{deployment}"
	resourceImagePath          = resourceDeploymentPath + "/image"
	resourceReplicasPath       = "/namespace/{namespace}/deployment/{deployment}/replicas"
)

// GetDeployment -- consumes a namespace and a deployment names,
// returns a Deployment data OR uninitialized struct AND an error
func (client *Client) GetDeployment(namespace, deployment string) (model.Deployment, error) {
	resp, err := client.Request.
		SetPathParams(map[string]string{
			"namespace":  namespace,
			"deployment": deployment,
		}).SetResult(model.Deployment{}).
		SetError(model.ResourceError{}).
		Get(client.serverURL + kubeAPIdeploymentPath)
	if err := catchErr(err, resp, http.StatusOK); err != nil {
		return model.Deployment{}, err
	}
	return *resp.Result().(*model.Deployment), nil
}

// GetDeploymentList -- consumes a namespace and a deployment names,
// returns a list of Deployments OR nil slice AND an error
func (client *Client) GetDeploymentList(namespace string) ([]model.Deployment, error) {
	resp, err := client.Request.
		SetPathParams(map[string]string{
			"namespace": namespace,
		}).SetResult([]model.Deployment{}).
		SetError(model.ResourceError{}).
		Get(client.serverURL + kubeAPIdeploymentsPath)
	if err := catchErr(err, resp, http.StatusOK); err != nil {
		return nil, err
	}
	return *resp.Result().(*[]model.Deployment), nil
}

// DeleteDeployment -- consumes a namespace, a deployment,
// an user role and an ID
func (client *Client) DeleteDeployment(namespace, deployment string) error {
	resp, err := client.Request.
		SetPathParams(map[string]string{
			"namespace":  namespace,
			"deployment": deployment,
		}).SetError(model.ResourceError{}).
		Delete(client.resourceServiceAddr + resourceDeploymentPath)
	return catchErr(err, resp, http.StatusOK)
}

// CreateDeployment -- consumes a namespace, an user ID and a Role,
// returns nil if OK
func (client *Client) CreateDeployment(namespace string, deployment model.Deployment) error {
	resp, err := client.Request.
		SetPathParams(map[string]string{
			"namespace": namespace,
		}).SetBody(deployment).
		SetError(model.ResourceError{}).
		Post(client.resourceServiceAddr + resourceDeploymentRootPath)
	return catchErr(err, resp,
		http.StatusOK,
		http.StatusCreated,
		http.StatusAccepted)
}

func (client *Client) SetContainerImage(namespace, deployment string, updateImage model.UpdateImage) error {
	resp, err := client.Request.
		SetPathParams(map[string]string{
			"namespace":  namespace,
			"deployment": deployment,
		}).SetBody(updateImage).
		Put(client.resourceServiceAddr + resourceImagePath)
	return catchErr(err, resp,
		http.StatusAccepted,
		http.StatusOK,
		http.StatusNoContent)
}

func (client *Client) ReplaceDeployment(namespace string, deployment model.Deployment) error {
	resp, err := client.Request.
		SetPathParams(map[string]string{
			"namespace":  namespace,
			"deployment": deployment.Name,
		}).SetBody(deployment).
		Put(client.resourceServiceAddr + resourceDeploymentPath)
	return catchErr(err, resp,
		http.StatusAccepted,
		http.StatusOK,
		http.StatusNoContent)
}

func (client *Client) SetReplicas(namespace, deployment string, replicas int) error {
	resp, err := client.Request.SetPathParams(map[string]string{
		"namespace":  namespace,
		"deployment": deployment,
	}).SetBody(model.UpdateReplicas{replicas}).
		Put(client.resourceServiceAddr + resourceReplicasPath)
	return catchErr(err, resp,
		http.StatusAccepted,
		http.StatusOK)
}
