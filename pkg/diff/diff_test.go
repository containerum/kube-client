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
				Image: "redis",
			},
			{
				Name:  "test",
				Image: "redis",
			},
		},
	}
	var newDepl = model.Deployment{
		Containers: []model.Container{
			{
				Name:  "test",
				Image: "redis",
			},
			{
				Name:  "gateway",
				Image: "nginx",
			},
		},
	}
	fmt.Println(FromContainer(oldDepl.Containers[0]))
	t.Log("\n", NewVersion(oldDepl, newDepl))
}
