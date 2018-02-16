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
				_, err := client.GetVolume(referenceVolumes[1].Label, nil)
				So(err, ShouldBeNil)
			})
			Convey("get volume list", func() {
				_, err := client.GetVolumeList(nil, nil)
				So(err, ShouldBeNil)
			})
			Convey("set delete access", func() {
				So(client.SetVolumeAccess("foxtrot", "fermi@mail.com", "read"),
					ShouldBeNil)
				So(client.DeleteVolumeAccess("foxtrot", "fermi@mail.com"),
					ShouldBeNil)
			})
			Convey("rename volume", func() {
				So(client.RenameVolume("foxtrot", "polka"), ShouldBeNil)
				So(client.RenameVolume("polka", "foxtrot"), ShouldBeNil)
			})
		})
	})
}
