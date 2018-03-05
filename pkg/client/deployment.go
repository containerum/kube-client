package client

import (
	"git.containerum.net/ch/kube-client/pkg/model"
	. "git.containerum.net/ch/kube-client/pkg/rest"
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
	err := client.re.Get(Rq{
		Result: &depl,
		Path: URL{
			Templ: client.APIurl + kubeAPIdeploymentPath,
			Params: P{
				"namespace":  namespace,
				"deployment": deployment,
			},
		},
	})
	return depl, err
}

// GetDeploymentList -- consumes a namespace and a deployment names,
// returns a list of Deployments OR nil slice AND an error
func (client *Client) GetDeploymentList(namespace string) ([]model.Deployment, error) {
	var depls []model.Deployment
	err := client.re.Get(Rq{
		Result: &depls,
		Path: URL{
			Templ: client.APIurl + kubeAPIdeploymentsPath,
			Params: P{
				"namespace": namespace,
			},
		},
	})
	return depls, err
}

// DeleteDeployment -- consumes a namespace, a deployment,
// an user role and an ID
func (client *Client) DeleteDeployment(namespace, deployment string) error {
	return client.re.Delete(Rq{
		Path: URL{
			Templ: client.APIurl + resourceDeploymentPath,
			Params: P{
				"namespace":  namespace,
				"deployment": deployment,
			},
		},
	})
}

// CreateDeployment -- consumes a namespace, an user ID and a Role,
// returns nil if OK
func (client *Client) CreateDeployment(namespace string, deployment model.Deployment) error {
	return client.re.Post(Rq{
		Body: deployment,
		Path: URL{
			Templ: client.APIurl + resourceDeploymentRootPath,
			Params: P{
				"namespace": namespace,
			},
		},
	})
}

// SetContainerImage -- set or changes deployment container image
// Consumes namespace, deployment and container data
func (client *Client) SetContainerImage(namespace, deployment string, updateImage model.UpdateImage) error {
	return client.re.Put(Rq{
		Body: updateImage,
		Path: URL{
			Templ: client.APIurl + resourceImagePath,
			Params: P{
				"namespace":  namespace,
				"deployment": deployment,
			},
		},
	})
}

// ReplaceDeployment -- replaces deployment in provided namespace with new one
func (client *Client) ReplaceDeployment(namespace string, deployment model.Deployment) error {
	return client.re.Put(Rq{
		Body: deployment,
		Path: URL{
			Templ: client.APIurl + resourceDeploymentPath,
			Params: P{
				"namespace":  namespace,
				"deployment": deployment.Name,
			},
		},
	})
}

// SetReplicas -- sets or changes deployment replicas
func (client *Client) SetReplicas(namespace, deployment string, replicas int) error {
	return client.re.Put(Rq{
		Body: model.UpdateReplicas{
			Replicas: replicas,
		},
		Path: URL{
			Templ: client.APIurl + resourceReplicasPath,
			Params: P{
				"namespace":  namespace,
				"deployment": deployment,
			},
		},
	})
}
