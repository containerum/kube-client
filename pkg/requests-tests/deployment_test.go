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
			deployment := newFakeResourceDeployment(test)
			namespace := "pion"
			deployment.Name = "fermi"
			updateImage := model.UpdateImage{
				Container: deployment.Containers[0].Name,
				Image:     "mongo",
			}
			Convey("create deployment", func() {
				err := client.CreateDeployment(namespace, deployment)
				So(err, ShouldBeNil)
			})
			Convey("set container image", func() {
				err := client.SetContainerImage(namespace,
					deployment.Name, updateImage)
				So(err, ShouldBeNil)
			})
			Convey("replace deployment", func() {
				deployment = newResourceUpdateDeployment(test)
				err := client.ReplaceDeployment(namespace, deployment)
				So(err, ShouldBeNil)
			})
			Convey("set replicas", func() {
				err := client.SetReplicas(namespace, deployment.Name, 6)
				So(err, ShouldBeNil)
			})
			Convey("delete deployment", func() {
				err := client.DeleteDeployment(namespace, deployment.Name)
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
