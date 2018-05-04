package model

// AvailableSolutionsList -- list of available solutions
//
// swagger:model
type AvailableSolutionsList struct {
	Solutions []AvailableSolution `json:"solutions"`
}

// AvailableSolution -- solution which user can run
//
// swagger:model
type AvailableSolution struct {
	Name   string          `json:"name"`
	Limits *SolutionLimits `json:"limits"`
	Images []string        `json:"images"`
	URL    string          `json:"url"`
	Active bool
}

// SolutionLimits -- solution resources limits
//
// swagger:model
type SolutionLimits struct {
	CPU string `json:"cpu"`
	RAM string `json:"ram"`
}

func (solution AvailableSolution) Copy() AvailableSolution {
	return AvailableSolution{
		Name: solution.Name,
		Limits: func() *SolutionLimits {
			if solution.Limits == nil {
				return nil
			}
			cp := *solution.Limits
			return &cp
		}(),
		Images: append([]string{}, solution.Images...),
		URL:    solution.URL,
		Active: solution.Active,
	}
}

// SolutionEnv -- solution environment variables
//
// swagger:model
type SolutionEnv struct {
	Env map[string]string `json:"env"`
}

// SolutionResources -- list of solution resources
//
// swagger:model
type SolutionResources struct {
	Resources map[string]int `json:"resources"`
}

type ConfigFile struct {
	Name string `json:"config_file"`
	Type string `json:"type"`
}

// UserSolutionsList -- list of running solution
//
// swagger:model
type UserSolutionsList struct {
	Solutions []UserSolution `json:"solutions"`
}

// UserSolution -- running solution
//
// swagger:model
type UserSolution struct {
	Branch string            `json:"branch"`
	Env    map[string]string `json:"env"`
	// required: true
	Template string `json:"template"`
	// required: true
	Name string `json:"name"`
	// required: true
	Namespace string `json:"namespace"`
}

// RunSolutionResponce -- responce to run solution request
//
// swagger:model
type RunSolutionResponce struct {
	Created    int      `json:"created"`
	NotCreated int      `json:"not_created"`
	Errors     []string `json:"errors,omitempty"`
}
