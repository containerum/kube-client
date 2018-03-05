package client

import (
	"git.containerum.net/ch/kube-client/pkg/rest"

	"git.containerum.net/ch/kube-client/pkg/model"
)

const (
	kubeAPIdeploymentPath  = "/namespaces/{namespace}/deployments/{deployment}"
	kubeAPIdeploymentsPath = "/namespaces/{namespace}/deployments"

	resourceDeploymentRootPath = "/namespace/{namespace}/deployment"
	resourceDeploymentPath     = "/namespace/{namespace}/deployment/{deployment}"
	resourceImagePath          = "/namespace/{namespace}/deployment/{deployment}/image"
	resourceReplicasPath       = "/namespace/{namespace}/deployment/{deployment}/replicas"
)

// GetDeployment -- consumes a namespace and a deployment names,
// returns a Deployment data OR uninitialized struct AND an error
func (client *Client) GetDeployment(namespace, deployment string) (model.Deployment, error) {
	var depl model.Deployment
	err := client.re.Get(&depl, rest.P{
		"namespace":  namespace,
		"deployment": deployment,
	}, client.APIurl, kubeAPIdeploymentPath)
	return depl, err
}

// GetDeploymentList -- consumes a namespace and a deployment names,
// returns a list of Deployments OR nil slice AND an error
func (client *Client) GetDeploymentList(namespace string) ([]model.Deployment, error) {
	var depls []model.Deployment
	err := client.re.Get(&depls, rest.P{
		"namespace": namespace,
	}, client.APIurl, kubeAPIdeploymentsPath)
	return depls, err
}

// DeleteDeployment -- consumes a namespace, a deployment,
// an user role and an ID
func (client *Client) DeleteDeployment(namespace, deployment string) error {
	return client.re.Delete(rest.P{
		"namespace":  namespace,
		"deployment": deployment,
	}, client.APIurl, resourceDeploymentPath)
}

// CreateDeployment -- consumes a namespace, an user ID and a Role,
// returns nil if OK
func (client *Client) CreateDeployment(namespace string, deployment model.Deployment) error {
	return client.re.Post(nil, deployment, rest.P{
		"namespace": namespace,
	}, client.APIurl, resourceDeploymentRootPath)
}

// SetContainerImage -- set or changes deployment container image
// Consumes namespace, deployment and container data
func (client *Client) SetContainerImage(namespace, deployment string, updateImage model.UpdateImage) error {
	return client.re.Put(nil, updateImage, rest.P{
		"namespace":  namespace,
		"deployment": deployment,
	}, client.APIurl, resourceImagePath)
}

// ReplaceDeployment -- replaces deployment in provided namespace with new one
func (client *Client) ReplaceDeployment(namespace string, deployment model.Deployment) error {
	return client.re.Put(nil, deployment, rest.P{
		"namespace":  namespace,
		"deployment": deployment.Name,
	}, client.APIurl, resourceDeploymentPath)
}

// SetReplicas -- sets or changes deployment replicas
func (client *Client) SetReplicas(namespace, deployment string, replicas int) error {
	return client.re.Put(nil, model.UpdateReplicas{
		Replicas: replicas,
	}, rest.P{
		"namespace":  namespace,
		"deployment": deployment,
	}, client.APIurl, resourceReplicasPath)
}
