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
				Image: "nginx:1.1",
			},
			{
				Name:  "feed",
				Image: "wordpress:2.9.2",
			},
		},
	}
	var newDepl = model.Deployment{
		Containers: []model.Container{
			{
				Name:  "gateway",
				Image: "caddy:1.5",
			},
			{
				Name:  "storage",
				Image: "mongo:45.0.9",
			},
			{
				Name:  "ai",
				Image: "pytnon:1.1.4",
			},
			{
				Name:  "blog",
				Image: "box",
			},
		},
	}
	t.Log("\n", Diff(newDepl, oldDepl))
}
