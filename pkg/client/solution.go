package client

import (
	"github.com/containerum/kube-client/pkg/rest"
	"github.com/containerum/kube-client/pkg/model"
)

const (
	solutionListPath      = "/solutions"
	solutionEnvPath       = solutionListPath + "/{solution}/env"
	solutionResourcesPath = solutionListPath + "{solution}/resources"
)

// GetSolutionList -- returns list of public solutions
func (client *Client) GetSolutionList() (model.AvailableSolutionsList, error) {
	var solutionList model.AvailableSolutionsList
	err := client.RestAPI.Get(rest.Rq{
		Result: &solutionList,
		URL: rest.URL{
			Path: solutionListPath,
		},
	})
	return solutionList, err
}
