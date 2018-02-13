package requests_tests

import (
	"math/rand"
	"testing"
	"time"

	"git.containerum.net/ch/kube-client/pkg/cmd"
	"git.containerum.net/ch/kube-client/pkg/model"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	resourceTestNamespace = "test-namespace"
	kubeAPItestNamespace  = "5020aa84-4827-47da-87ee-5fc2cf18c111"
	kubeAPItestDeployment = "roma"
)

var (
	randomGen = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func TestDeployment(test *testing.T) {
	client, err := cmd.CreateCmdClient(
		cmd.ClientConfig{
			ResourceAddr: "http://192.168.88.200:1213",
			APIurl:       "http://192.168.88.200:1214",
			User: cmd.User{
				Role: "admin",
			},
		})
	if err != nil {
		test.Fatalf("error while creating client: %v", err)
	}
	Convey("Test resource service methods", test, func() {
		fakeResourceDeployment := newFakeResourceDeployment(test)
		fakeUpdateImage := newFakeResourceUpdateImage(test)
		Convey("deployment creation test",
			deploymentCreationTest(client, resourceTestNamespace, fakeResourceDeployment))
		Convey("set container image test",
			setContainerImageTest(client, resourceTestNamespace, fakeResourceDeployment.Name, fakeUpdateImage))
	})
	Convey("Test KubeAPI methods", test, func() {
		fakeKubeAPIdeployment := newFakeKubeAPIdeployment(test)
		Convey("get deployment test",
			getDeploymentTest(client, kubeAPItestNamespace,
				kubeAPItestDeployment, fakeKubeAPIdeployment))
		Convey("get deployment list",
			getDeploymentListTest(client, kubeAPItestNamespace, []model.Deployment{fakeKubeAPIdeployment}))
	})
}

func deploymentCreationTest(client *cmd.Client, namespace string, deployment model.Deployment) func() {
	return func() {
		err := client.CreateDeployment(namespace, deployment)
		So(err, ShouldBeNil)
	}
}

func setContainerImageTest(client *cmd.Client, namespace, deployment string, updateImage model.UpdateImage) func() {
	return func() {
		err := client.SetContainerImage(namespace, deployment, updateImage)
		So(err, ShouldBeNil)
	}
}
func getDeploymentTest(client *cmd.Client, namespace, deployment string, referenceDeployment model.Deployment) func() {
	return func() {
		/*gainedDeployment*/ _, err := client.GetDeployment(namespace, deployment)
		So(err, ShouldBeNil)
		//So(gainedDeployment, ShouldEqual, referenceDeployment)
	}
}

func getDeploymentListTest(client *cmd.Client, namespace string, referenceList []model.Deployment) func() {
	return func() {
		/*gainedDeploymentList*/ _, err := client.GetDeploymentList(namespace)
		So(err, ShouldBeNil)
		//So(gainedDeploymentList, ShouldEqual, referenceList)
	}
}
