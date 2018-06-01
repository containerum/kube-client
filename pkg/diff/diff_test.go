package diff

import (
	"testing"

	"github.com/containerum/kube-client/pkg/model"
)

func TestDiff(t *testing.T) {
	var oldDepl = model.Deployment{
		Containers: []model.Container{
			{
				Name:  "gateway",
				Image: "nginx",
			},
			{
				Name:  "feed",
				Image: "wordpress",
			},
		},
	}
	var newDepl = model.Deployment{
		Containers: []model.Container{
			{
				Name:  "gateway",
				Image: "caddy",
			},
			{
				Name:  "storage",
				Image: "mongo",
			},
			{
				Name:  "ai",
				Image: "pytnon",
			},
			{
				Name:  "blog",
				Image: "box",
			},
		},
	}
	t.Log("\n", Diff(newDepl, oldDepl))
}
