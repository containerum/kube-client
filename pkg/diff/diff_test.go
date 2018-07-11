package diff

import (
	"testing"

	"fmt"

	"github.com/blang/semver"
	"github.com/containerum/kube-client/pkg/model"
)

func TestDiff(t *testing.T) {
	var oldDepl = model.Deployment{
		Version: semver.MustParse("1.0.0"),
		Containers: []model.Container{
			{
				Name:  "gateway",
				Image: "nginx:1.0",
			},
		},
	}
	var newDepl = model.Deployment{
		Containers: []model.Container{
			{
				Name:  "gateway",
				Image: "nginx:2.0",
			},
		},
	}
	fmt.Println(FromContainer(oldDepl.Containers[0]))
	t.Log("\n", NewVersion(oldDepl, newDepl))
}
