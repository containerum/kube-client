package diff

import (
	"testing"

	"github.com/blang/semver"
	"github.com/containerum/kube-client/pkg/model"
)

func TestDiff(t *testing.T) {
	var oldDepl = model.Deployment{
		Version: semver.MustParse("1.0.0"),
		Containers: []model.Container{
			{
				Name:  "gateway",
				Image: "nginx",
			},
		},
	}
	var newDepl = model.Deployment{
		Containers: []model.Container{
			{
				Name:  "gateway",
				Image: "nginx:latest",
			},
		},
	}
	t.Log("\n", NewVersion(oldDepl, newDepl))
}
