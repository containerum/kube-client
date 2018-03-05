package reqtests

import (
	"fmt"
	"testing"

	kubeClient "git.containerum.net/ch/kube-client/pkg/client"
	"git.containerum.net/ch/kube-client/pkg/model"
	"git.containerum.net/ch/kube-client/pkg/rest/remock"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	testAPIurl = "http://192.168.88.200"
)

const (
	kubeAPItestNamespace  = "5020aa84-4827-47da-87ee-5fc2cf18c111"
	kubeAPItestDeployment = "roma"
)

func TestDeployment(test *testing.T) {
	client, err := kubeClient.NewClient(kubeClient.Config{
		APIurl:  testAPIurl,
		RestAPI: remock.NewMock(),
		User: kubeClient.User{
			Role: "user",
		},
	})
	if err != nil {
		test.Fatalf("error while test cli init: %v", err)
	}
	Convey("Test deployment methods", test, func() {
		Convey("resource service methods", func() {
			deployment := newFakeDeployment(test)
			namespace := "pion"
			deployment.Name = newRandomName(6)
			updateImage := model.UpdateImage{
				Container: deployment.Containers[0].Name,
				Image:     "mongo",
			}
			er := client.CreateDeployment(namespace, deployment)
			fmt.Printf("%#v", er)
			So(er, ShouldBeNil)

			err = client.SetContainerImage(namespace,
				deployment.Name, updateImage)
			So(err, ShouldBeNil)
			deployment.Labels["color"] = "blue"
			err = client.ReplaceDeployment(namespace, deployment)
			So(err, ShouldBeNil)

			err = client.SetReplicas(namespace, deployment.Name, 6)
			So(err, ShouldBeNil)

			err = client.DeleteDeployment(namespace, deployment.Name)
			So(err, ShouldBeNil)

		})
		Convey("KubeAPI methods", func() {
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
