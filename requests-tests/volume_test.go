package reqtests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	resourceTestVolume = "n1-volume"
)

func TestVolume(test *testing.T) {
	client := newMockClient(test)
	Convey("Test volume methods", test, func() {
		Convey("resource api", func() {
			Convey("get volume and list", func() {
				volumes, err := client.GetVolumeList("")
				So(err, ShouldBeNil)
				So(len(volumes.Volumes), ShouldBeGreaterThan, 0)
				gainedVolume, err := client.GetVolume("", volumes.Volumes[0].Name)
				So(err, ShouldBeNil)
				So(gainedVolume, ShouldResemble, volumes.Volumes[0])
			})
		})
	})
}
