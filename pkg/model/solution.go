package model

type Solution struct {
	Env map[string]string `json:"env"`
	Run []ConfigFile      `json:"run,omitempty"`
}

type ConfigFile struct {
	Name string `json:"config_file"`
	Type string `json:"type"`
}

func (solution Solution) Copy() Solution {
	env := make(map[string]string, len(solution.Env))
	for k, v := range solution.Env {
		env[k] = v
	}
	return Solution{
		Env: env,
		Run: append(make([]ConfigFile, 0, len(solution.Run)), solution.Run...),
	}
}

func (solution Solution) GetFileByName(name string) (ConfigFile, bool) {
	for _, conFile := range solution.Run {
		if conFile.Name == name {
			return conFile, true
		}
	}
	return ConfigFile{}, false
}

func (solution Solution) GetFilesByType(tt string) []ConfigFile {
	confFiles := make([]ConfigFile, 0, len(solution.Run))
	for _, conFile := range solution.Run {
		if conFile.Type == tt {
			confFiles = append(confFiles, conFile)
		}
	}
	return confFiles
}

