package requests_tests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	resourceTestVolume = "n1-volume"
)

func TestVolume(test *testing.T) {
	client := newResourceClient(test)
	Convey("Test volume methods", test, func() {
		Convey("resource api", func() {
			referenceVolumes := newFakeResourceVolume(test)
			Convey("get volume", func() {
				gainedVolume, err := client.GetVolume(referenceVolumes[0].Label, nil)
				So(err, ShouldBeNil)
				So(gainedVolume, ShouldResemble, referenceVolumes[0])
			})
			Convey("get volume list", func() {
				_, err := client.GetVolumeList(nil, nil)
				So(err, ShouldBeNil)
			})
		})
	})
}
