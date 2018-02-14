package requests_tests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	resourceTestNamespace = "test-namespace"
	kubeAPItestNamespace  = "5020aa84-4827-47da-87ee-5fc2cf18c111"
	kubeAPItestDeployment = "roma"
)

func TestDeployment(test *testing.T) {
	Convey("Test deployment methods", test, func() {
		Convey("resource service methods", func() {
			client := newResourceClient(test)
			fakeResourceDeployment := newFakeResourceDeployment(test)
			fakeUpdateImage := newFakeResourceUpdateImage(test)
			Convey("create deployment", func() {
				err := client.CreateDeployment(resourceTestNamespace, fakeResourceDeployment)
				So(err, ShouldBeNil)
			})
			Convey("set container image", func() {
				err := client.SetContainerImage(resourceTestNamespace,
					fakeResourceDeployment.Name, fakeUpdateImage)
				So(err, ShouldBeNil)
			})
			Convey("replace deployment", func() {
				err := client.ReplaceDeployment(resourceTestNamespace, fakeResourceDeployment)
				So(err, ShouldBeNil)
			})
			Convey("set replicas", func() {
				err := client.SetReplicas(resourceTestNamespace, fakeResourceDeployment.Name, 4)
				So(err, ShouldBeNil)
			})
			Convey("delete deployment", func() {
				err := client.DeleteDeployment(resourceTestNamespace, fakeResourceDeployment.Name)
				So(err, ShouldBeNil)
			})
		})
		Convey("KubeAPI methods", func() {
			client := newCubeAPIClient(test)
			Convey("get deployment test", func() {
				_, err := client.GetDeployment(kubeAPItestNamespace, kubeAPItestDeployment)
				So(err, ShouldBeNil)
			})
			Convey("get deployment list", func() {
				_, err := client.GetDeploymentList(kubeAPItestNamespace)
				So(err, ShouldBeNil)
			})
		})
	})
}
