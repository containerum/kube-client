package cmd

import (
	"git.containerum.net/ch/kube-client/pkg/model"
)

const (
	deploymentPath  = "/namespaces/{namespace}/deployments/{deployment}"
	deploymentsPath = "/namespaces/{namespace}/deployments"

	resourceDeploymentPath = "/namespace/{namespace}/deployment/{deployment}"
	resourceImagePath      = resourceDeploymentPath + "/image"
	resourceReplicasPath   = resourceDeploymentPath + "/replicas"
)

// GetDeployment -- consumes a namespace and a deployment names,
// returns a Deployment data OR uninitialized struct AND an error
func (client *Client) GetDeployment(namespace, deployment string) (model.Deployment, error) {
	resp, err := client.Request.
		SetPathParams(map[string]string{
			"namespace":  namespace,
			"deployment": deployment,
		}).SetResult(model.Deployment{}).
		Get(client.serverURL + deploymentPath)
	if err != nil {
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
		Get(client.serverURL + deploymentsPath)
	if err != nil {
		return nil, err
	}
	return *resp.Result().(*[]model.Deployment), nil
}

// CreateDeployment -- consumes a namespace, an user ID and a Role,
// returns nil if OK
func (client *Client) CreateDeployment(namespace string, deployment model.Deployment) error {
	_, err := client.Request.
		SetPathParams(map[string]string{
			"namespace": namespace,
		}).SetBody(deployment).
		Post(client.resourceServiceAddr + "/namespace/{namespace}/deployment")
	return err
}

func (client *Client) SetContainerImage(namespace, deployment string, containerImage model.ContainerImage) error {
	_, err := client.Request.
		SetPathParams(map[string]string{
			"namespace":  namespace,
			"deployment": deployment,
		}).SetBody(containerImage).
		Put(client.resourceServiceAddr + resourceImagePath)
	return err
}

func (client *Client) ReplaceDeployment(namespace string, deployment model.Deployment) error {
	_, err := client.Request.
		SetPathParams(map[string]string{
			"namespace":  namespace,
			"deployment": deployment.Name,
		}).SetBody(deployment).
		Put(client.resourceServiceAddr + resourceDeploymentPath)
	return err
}

func (client *Client) SetReplicas(namespace, deployment string, replicas int) error {
	_, err := client.Request.SetPathParams(map[string]string{
		"namespace":  namespace,
		"deployment": deployment,
	}).SetBody(model.UpdateReplicas{replicas}).
		Put(client.resourceServiceAddr + resourceReplicasPath)
	return err
}
