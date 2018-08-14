package reqtests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNamespace(test *testing.T) {
	client := newMockClient(test)
	fakeNamespaces := newFakeNamespaces(test)
	Convey("Test KubeAPI methods", test, func() {
		Convey("get namespace", func() {
			_, err := client.GetNamespace(kubeAPItestNamespace)
			So(err, ShouldBeNil)
		})
		Convey("get namespace list", func() {
			_, err := client.GetNamespaceList()
			So(err, ShouldBeNil)
		})
	})
	Convey("Test resource service methods", test, func() {
		Convey("get namespace", func() {
			gainedNamespace, err := client.ResourceGetNamespace(fakeNamespaces[0].ID)
			So(err, ShouldBeNil)
			So(gainedNamespace, ShouldResemble, fakeNamespaces[0])
		})
		Convey("get namespace list", func() {
			_, err := client.ResourceGetNamespaceList(0, 16)
			So(err, ShouldBeNil)
		})
		Convey("rename namespace", func() {
			err := client.RenameNamespace(fakeNamespaces[0].ID, "electron")
			So(err, ShouldBeNil)
			err = client.RenameNamespace("electron", fakeNamespaces[0].ID)
			So(err, ShouldBeNil)
		})
		Convey("set access", func() {
			err := client.SetNamespaceAccess(fakeNamespaces[0].ID, "fermi@da.com", "read")
			So(err, ShouldBeNil)
		})
		Convey("delete access", func() {
			err := client.DeleteNamespaceAccess(fakeNamespaces[0].ID, "fermi@da.com")
			So(err, ShouldBeNil)
		})
	})

}
