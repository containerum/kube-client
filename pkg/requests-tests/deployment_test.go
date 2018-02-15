package requests_tests

import (
	"testing"

	"git.containerum.net/ch/kube-client/pkg/model"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	kubeAPItestNamespace  = "5020aa84-4827-47da-87ee-5fc2cf18c111"
	kubeAPItestDeployment = "roma"
)

func TestDeployment(test *testing.T) {
	Convey("Test deployment methods", test, func() {
		Convey("resource service methods", func() {
			client := newResourceClient(test)
			//fakeNamespaces := newFakeResourceNamespaces(test)
			fakeResourceDeployment := newFakeResourceDeployment(test)

			namespace := "pion"
			fakeResourceDeployment.Name = "fermi"
			updateImage := model.UpdateImage{
				Container: fakeResourceDeployment.Containers[0].Name,
				Image:     "mongo",
			}
			Convey("create deployment", func() {
				err := client.CreateDeployment(namespace, fakeResourceDeployment)
				So(err, ShouldBeNil)
			})
			Convey("set container image", func() {
				err := client.SetContainerImage(namespace,
					fakeResourceDeployment.Name, updateImage)
				So(err, ShouldBeNil)
			})
			Convey("replace deployment", func() {
				err := client.ReplaceDeployment(namespace, fakeResourceDeployment)
				So(err, ShouldBeNil)
			})
			Convey("set replicas", func() {
				err := client.SetReplicas(namespace, fakeResourceDeployment.Name, 6)
				So(err, ShouldBeNil)
			})
			Convey("delete deployment", func() {
				err := client.DeleteDeployment(namespace, fakeResourceDeployment.Name)
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
