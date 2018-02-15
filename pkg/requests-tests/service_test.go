package requests_tests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	kubeAPItestService = "ch-glusterfs"
)

func TestService(test *testing.T) {
	client := newCubeAPIClient(test)
	Convey("Test Kube API methods", test, func() {
		Convey("get service", func() {
			_, err := client.GetService(kubeAPItestNamespace, kubeAPItestService)
			So(err, ShouldBeNil)
		})
		Convey("get service list", func() {
			_, err := client.GetServiceList(kubeAPItestNamespace)
			So(err, ShouldBeNil)
		})
	})
}
