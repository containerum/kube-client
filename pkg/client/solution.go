package client

import (
	"github.com/containerum/kube-client/pkg/model"
	"github.com/containerum/kube-client/pkg/rest"
)

const (
	templatesPath   = "/templates"
	templateEnvPath = "/templates/{template}/env"
	templateResPath = "/templates/{template}/resources"
	solutionsPath   = "/solutions"
	solutionPath    = "/solutions/{solution}"
)

// GetSolutionsTemplatesList -- returns list of public solutions templates
func (client *Client) GetSolutionsTemplatesList() (model.AvailableSolutionsList, error) {
	var solutionList model.AvailableSolutionsList
	err := client.RestAPI.Get(rest.Rq{
		Result: &solutionList,
		URL: rest.URL{
			Path: templatesPath,
		},
	})
	return solutionList, err
}

// GetSolutionsTemplateEnv -- returns list of public solution template envs
func (client *Client) GetSolutionsTemplateEnv(templateName string) (model.SolutionEnv, error) {
	var solutionEnv model.SolutionEnv
	err := client.RestAPI.Get(rest.Rq{
		Result: &solutionEnv,
		URL: rest.URL{
			Path: templateEnvPath,
			Params: rest.P{
				"template": templateName,
			},
		},
	})
	return solutionEnv, err
}

// GetSolutionsTemplateResources -- returns count of public solution template resources
func (client *Client) GetSolutionsTemplateResources(templateName string) (model.SolutionResources, error) {
	var solutionRes model.SolutionResources
	err := client.RestAPI.Get(rest.Rq{
		Result: &solutionRes,
		URL: rest.URL{
			Path: templateResPath,
			Params: rest.P{
				"template": templateName,
			},
		},
	})
	return solutionRes, err
}

// RunSolution -- creates new solution
func (client *Client) RunSolution(solution model.UserSolution) (model.RunSolutionResponse, error) {
	var resp model.RunSolutionResponse
	err := client.RestAPI.Post(rest.Rq{
		Result: &resp,
		Body:   solution.Copy(),
		URL: rest.URL{
			Path: solutionsPath,
		},
	})
	return resp, err
}

// GetSolutionsList -- returns list of users running solutions
func (client *Client) GetSolutionsList() (model.UserSolutionsList, error) {
	var solutionList model.UserSolutionsList
	err := client.RestAPI.Get(rest.Rq{
		Result: &solutionList,
		URL: rest.URL{
			Path: solutionsPath,
		},
	})
	return solutionList, err
}

// GetSolution -- returns user running solutions
func (client *Client) GetSolution(solutionName string) (model.UserSolutionsList, error) {
	var solutionList model.UserSolutionsList
	err := client.RestAPI.Get(rest.Rq{
		Result: &solutionList,
		URL: rest.URL{
			Path: solutionPath,
			Params: rest.P{
				"solution": solutionName,
			},
		},
	})
	return solutionList, err
}
