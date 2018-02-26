package reqtests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNamespace(test *testing.T) {
	client := newClient(test)
	fakeNamespaces := newFakeNamespaces(test)
	Convey("Test KubeAPI methods", test, func() {
		Convey("get namespace", func() {
			_, err := client.GetNamespace(kubeAPItestNamespace)
			So(err, ShouldBeNil)
		})
		Convey("get namespace list", func() {
			_, err := client.GetNamespaceList(map[string]string{})
			So(err, ShouldBeNil)
		})
	})
	Convey("Test resource service methods", test, func() {
		Convey("get namespace", func() {
			gainedNamespace, err := client.ResourceGetNamespace(fakeNamespaces[0].Label, nil)
			So(err, ShouldBeNil)
			So(gainedNamespace, ShouldResemble, fakeNamespaces[0])
		})
		Convey("get namespace list", func() {
			_, err := client.ResourceGetNamespaceList(0, 16, "")
			So(err, ShouldBeNil)
		})
	})
}
