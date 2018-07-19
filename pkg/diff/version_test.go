package diff

import (
	"testing"

	"github.com/blang/semver"
	"github.com/containerum/kube-client/pkg/model"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNoChange(t *testing.T) {
	var oldDepl = model.Deployment{
		Version: semver.MustParse("1.0.0"),
		Containers: []model.Container{
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
		},
	}
	Convey("Version should be...", t, func() {
		So(NewVersion(oldDepl, newDepl).String(), ShouldEqual, oldDepl.Version.String())
	})
}

func TestChangeSemverMinorVersion(t *testing.T) {
	var oldDepl = model.Deployment{
		Version: semver.MustParse("1.0.0"),
		Containers: []model.Container{
			{
				Name:  "test",
				Image: "redis:1.0.0",
			},
		},
	}
	var newDepl = model.Deployment{
		Containers: []model.Container{
			{
				Name:  "test",
				Image: "redis:1.2.3",
			},
		},
	}
	Convey("Version should be...", t, func() {
		So(NewVersion(oldDepl, newDepl).String(), ShouldEqual, "1.1.0")
	})
}

func TestReplaceContainerName(t *testing.T) {
	var oldDepl = model.Deployment{
		Version: semver.MustParse("1.0.0"),
		Containers: []model.Container{
			{
				Name:  "test",
				Image: "redis",
			},
		},
	}
	var newDepl = model.Deployment{
		Containers: []model.Container{
			{
				Name:  "test-new",
				Image: "redis",
			},
		},
	}
	Convey("Version should be...", t, func() {
		So(NewVersion(oldDepl, newDepl).String(), ShouldEqual, "2.0.0")
	})
}

func TestNoVersionToSemver(t *testing.T) {
	var oldDepl = model.Deployment{
		Version: semver.MustParse("1.0.0"),
		Containers: []model.Container{
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
				Image: "redis:1.1",
			},
		},
	}
	Convey("Version should be...", t, func() {
		So(NewVersion(oldDepl, newDepl).String(), ShouldEqual, "2.0.0")
	})
}

func TestSemverToNoVersion(t *testing.T) {
	var oldDepl = model.Deployment{
		Version: semver.MustParse("1.0.0"),
		Containers: []model.Container{
			{
				Name:  "test",
				Image: "redis:1.0",
			},
		},
	}
	var newDepl = model.Deployment{
		Containers: []model.Container{
			{
				Name:  "test",
				Image: "redis",
			},
		},
	}
	Convey("Version should be...", t, func() {
		So(NewVersion(oldDepl, newDepl).String(), ShouldEqual, "2.0.0")
	})
}

func TestSemverToNotSemver(t *testing.T) {
	var oldDepl = model.Deployment{
		Version: semver.MustParse("1.0.0"),
		Containers: []model.Container{
			{
				Name:  "test",
				Image: "redis:1.0",
			},
		},
	}
	var newDepl = model.Deployment{
		Containers: []model.Container{
			{
				Name:  "test",
				Image: "redis:test",
			},
		},
	}
	Convey("Version should be...", t, func() {
		So(NewVersion(oldDepl, newDepl).String(), ShouldEqual, "2.0.0")
	})
}

func TestNotSemverToSemver(t *testing.T) {
	var oldDepl = model.Deployment{
		Version: semver.MustParse("1.0.0"),
		Containers: []model.Container{
			{
				Name:  "test",
				Image: "redis:test",
			},
		},
	}
	var newDepl = model.Deployment{
		Containers: []model.Container{
			{
				Name:  "test",
				Image: "redis:1.0",
			},
		},
	}
	Convey("Version should be...", t, func() {
		So(NewVersion(oldDepl, newDepl).String(), ShouldEqual, "2.0.0")
	})
}

func TestReplaceOneImage(t *testing.T) {
	var oldDepl = model.Deployment{
		Version: semver.MustParse("1.0.0"),
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
	var newDepl = model.Deployment{
		Containers: []model.Container{
			{
				Name:  "test",
				Image: "nginx",
			},
			{
				Name:  "gateway",
				Image: "nginx",
			},
		},
	}
	Convey("Version should be...", t, func() {
		So(NewVersion(oldDepl, newDepl).String(), ShouldEqual, "2.0.0")
	})
}

func TestAddContainer(t *testing.T) {
	var oldDepl = model.Deployment{
		Version: semver.MustParse("1.0.0"),
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
			{
				Name:  "test-1",
				Image: "redis",
			},
		},
	}
	Convey("Version should be...", t, func() {
		So(NewVersion(oldDepl, newDepl).String(), ShouldEqual, "1.1.0")
	})
}

func TestDeleteContainer(t *testing.T) {
	var oldDepl = model.Deployment{
		Version: semver.MustParse("1.0.0"),
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
	var newDepl = model.Deployment{
		Containers: []model.Container{
			{
				Name:  "test",
				Image: "redis",
			},
		},
	}
	Convey("Version should be...", t, func() {
		So(NewVersion(oldDepl, newDepl).String(), ShouldEqual, "2.0.0")
	})
}
