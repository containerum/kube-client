package remock

import (
	kubeAPImodels "git.containerum.net/ch/kube-api/pkg/model"
	"git.containerum.net/ch/kube-client/pkg/model"
)

func ValidateDeployment(deployment model.Deployment) error {
	deploymentWithOwner := kubeAPImodels.DeploymentWithOwner{
		deployment,
		"owner",
	}
	errs := kubeAPImodels.ValidateDeployment(deploymentWithOwner)
	for _, err := range errs {
		return err
	}
	return nil
}
