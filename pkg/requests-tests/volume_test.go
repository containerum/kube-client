package requests_tests

import (
	"testing"

	"git.containerum.net/ch/kube-client/pkg/model"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	resourceTestVolume = "n1-volume"
)

func TestVolume(test *testing.T) {
	client := newResourceClient(test)
	Convey("Test volume methods", test, func() {
		Convey("resource api", func() {
			var volumes []model.ResourceVolume
			Convey("get volume and list", func() {
				var err error
				volumes, err = client.GetVolumeList(nil, nil)
				So(err, ShouldBeNil)
				So(len(volumes), ShouldBeGreaterThan, 0)
				gainedVolume, err := client.GetVolume(volumes[0].Label, nil)
				So(err, ShouldBeNil)
				So(gainedVolume, ShouldResemble, volumes[0])
			})
		})
	})
}
