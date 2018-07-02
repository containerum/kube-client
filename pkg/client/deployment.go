package client

import (
	"github.com/blang/semver"
	"github.com/containerum/kube-client/pkg/model"
	"github.com/containerum/kube-client/pkg/rest"
)

const (
	deploymentsPath                   = "/projects/{project}/namespaces/{namespace}/deployments"
	deploymentPath                    = "/projects/{project}/namespaces/{namespace}/deployments/{deployment}"
	deploymentVersionsPath            = "/projects/{project}/namespaces/{namespace}/deployments/{deployment}/versions"
	deploymentVersionPath             = "/projects/{project}/namespaces/{namespace}/deployments/{deployment}/versions/{version}"
	imagePath                         = "/projects/{project}/namespaces/{namespace}/deployments/{deployment}/image"
	replicasPath                      = "/projects/{project}/namespaces/{namespace}/deployments/{deployment}/replicas"
	deploymentDiffWithPreviousVersion = "/projects/{project}/namespaces/{namespace}/deployments/{deployment}/versions/{version}/diff"
	deploymentDiffBetweenVersions     = "/projects/{project}/namespaces/{namespace}/deployments/{deployment}/versions/{left-version}/diff/{right-version}"
)

// GetDeployment -- consumes a namespace and a deployment names,
// returns a Deployment data OR uninitialized struct AND an error
func (client *Client) GetDeployment(project, namespace, deployment string) (model.Deployment, error) {
	var depl model.Deployment
	err := client.RestAPI.Get(rest.Rq{
		Result: &depl,
		URL: rest.URL{
			Path: deploymentPath,
			Params: rest.P{
				"namespace":  namespace,
				"deployment": deployment,
				"project":    project,
			},
		},
	})
	return depl, err
}

// GetDeploymentList -- consumes a namespace and a deployment names,
// returns a list of Deployments OR nil slice AND an error
func (client *Client) GetDeploymentList(project, namespace string) (model.DeploymentsList, error) {
	var depls model.DeploymentsList
	err := client.RestAPI.Get(rest.Rq{
		Result: &depls,
		URL: rest.URL{
			Path: deploymentsPath,
			Params: rest.P{
				"namespace": namespace,
				"project":   project,
			},
		},
	})
	return depls, err
}

// DeleteDeployment -- consumes a namespace, a deployment,
// an user role and an ID
func (client *Client) DeleteDeployment(project, namespace, deployment string) error {
	return client.RestAPI.Delete(rest.Rq{
		URL: rest.URL{
			Path: deploymentPath,
			Params: rest.P{
				"namespace":  namespace,
				"deployment": deployment,
				"project":    project,
			},
		},
	})
}

// CreateDeployment -- consumes a namespace, an user ID and a Role,
// returns nil if OK
func (client *Client) CreateDeployment(project, namespace string, deployment model.Deployment) error {
	return client.RestAPI.Post(rest.Rq{
		Body: deployment,
		URL: rest.URL{
			Path: deploymentsPath,
			Params: rest.P{
				"namespace": namespace,
				"project":   project,
			},
		},
	})
}

// SetContainerImage -- set or changes deployment container image
// Consumes namespace, deployment and container data
func (client *Client) SetContainerImage(project, namespace, deployment string, updateImage model.UpdateImage) error {
	return client.RestAPI.Put(rest.Rq{
		Body: updateImage,
		URL: rest.URL{
			Path: imagePath,
			Params: rest.P{
				"namespace":  namespace,
				"deployment": deployment,
				"project":    project,
			},
		},
	})
}

// ReplaceDeployment -- replaces deployment in provided namespace with new one
func (client *Client) ReplaceDeployment(project, namespace string, deployment model.Deployment) error {
	return client.RestAPI.Put(rest.Rq{
		Body: deployment,
		URL: rest.URL{
			Path: deploymentPath,
			Params: rest.P{
				"namespace":  namespace,
				"deployment": deployment.Name,
				"project":    project,
			},
		},
	})
}

// SetReplicas -- sets or changes deployment replicas
func (client *Client) SetReplicas(project, namespace, deployment string, replicas int) error {
	return client.RestAPI.Put(rest.Rq{
		Body: model.UpdateReplicas{
			Replicas: replicas,
		},
		URL: rest.URL{
			Path: replicasPath,
			Params: rest.P{
				"namespace":  namespace,
				"deployment": deployment,
				"project":    project,
			},
		},
	})
}

// Returns list of defferent deployment versions
func (client *Client) GetDeploymentVersions(project, namespace, deplName string) (model.DeploymentsList, error) {
	var list model.DeploymentsList
	return list, client.RestAPI.Get(rest.Rq{
		Result: &list,
		URL: rest.URL{
			Path: deploymentVersionsPath,
			Params: rest.P{
				"namespace":  namespace,
				"deployment": deplName,
				"project":    project,
			},
		},
	})
}

// Create pods from deployment with specific version
func (client *Client) RunDeploymentVersion(project, namespace, deplName string, version semver.Version) error {
	return client.RestAPI.Post(rest.Rq{
		URL: rest.URL{
			Path: deploymentVersionPath,
			Params: rest.P{
				"namespace":  namespace,
				"deployment": deplName,
				"version":    version.String(),
				"project":    project,
			},
		},
	})
}

func (client *Client) GetDeploymentDiffWithPreviousVersion(namespace, deployment string, version semver.Version) (string, error) {
	var diff model.DeploymentDiff
	return diff.Diff, client.RestAPI.Get(rest.Rq{
		Result: &diff,
		URL: rest.URL{
			Path: deploymentDiffWithPreviousVersion,
			Params: rest.P{
				"namespace":  namespace,
				"deployment": deployment,
				"version":    version.String(),
			},
		},
	})
}

func (client *Client) GetDeloymentVersionBetweenVersions(namespace, deployment string, leftVersion, rightVersion semver.Version) (string, error) {
	var diff model.DeploymentDiff
	return diff.Diff, client.RestAPI.Get(rest.Rq{
		Result: &diff,
		URL: rest.URL{
			Path: deploymentDiffBetweenVersions,
			Params: rest.P{
				"namespace":     namespace,
				"deployment":    deployment,
				"left-version":  leftVersion.String(),
				"right-version": rightVersion.String(),
			},
		},
	})
}

func (client *Client) DeleteDeploymentVersion(namespace, deployment string, version semver.Version) error {
	return client.RestAPI.Delete(rest.Rq{
		URL: rest.URL{
			Path: deploymentVersionPath,
			Params: rest.P{
				"namespace":  namespace,
				"deployment": deployment,
				"version":    version.String(),
			},
		},
	})
}
