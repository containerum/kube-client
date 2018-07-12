package diff

import (
	"testing"

	"github.com/blang/semver"
	"github.com/containerum/kube-client/pkg/model"
)

func TestNewVersion(t *testing.T) {
	var oldDepl = model.Deployment{
		Version: semver.MustParse("1.0.1"),
		Containers: []model.Container{
			{
				Name:  "gateway",
				Image: "nginx:sdfsdfs",
			},
		},
	}
	var newDepl = model.Deployment{
		Containers: []model.Container{
			{
				Name:  "gateway",
				Image: "nginx:sdfsfd",
			},
		},
	}
	t.Log("\nOld version:", oldDepl.Version, "\nNew version:", NewVersion(oldDepl, newDepl))
}
