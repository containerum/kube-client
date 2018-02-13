package requests_test

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"git.containerum.net/ch/kube-client/pkg/cmd"
	"git.containerum.net/ch/kube-client/pkg/model"
)

const (
	testNamespace         = "test-namespace"
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
	client.SetHeaders(map[string]string{
		"X-User-Role": "admin",
	})
	{
		fakeResourceDeployment := newFakeResourceDeployment(test)
		fakeUpdateImage := newFakeResourceUpdateImage(test)
		test.Run("deployment creaton test",
			deploymentCreationTest(client, testNamespace, fakeResourceDeployment))
		test.Run("set container image test",
			setContainerImageTest(client, testNamespace, fakeResourceDeployment.Name, fakeUpdateImage))
	}
	{
		fakeKubeAPIdeployment := newFakeKubeAPIdeployment(test)
		test.Run("get deployment test",
			getDeploymentTest(client, kubeAPItestNamespace,
				kubeAPItestDeployment, fakeKubeAPIdeployment))
		test.Run("get deployment list",
			getDeploymentListTest(client, kubeAPItestNamespace, []model.Deployment{fakeKubeAPIdeployment}))
	}
}

func deploymentCreationTest(client *cmd.Client, namespace string, deployment model.Deployment) func(*testing.T) {
	return func(test *testing.T) {
		err := client.CreateDeployment(namespace, deployment)
		if err != nil {
			test.Fatalf("error while deployment creation: %v", err)
		}
	}
}

func setContainerImageTest(client *cmd.Client, namespace, deployment string, updateImage model.UpdateImage) func(*testing.T) {
	return func(test *testing.T) {
		client.SetHeaders(map[string]string{
			"X-User-ID":   "20b616d8-1ea7-4842-b8ec-c6e8226fda5b",
			"X-User-Role": "user",
		})
		err := client.SetContainerImage(namespace, deployment, updateImage)
		if err != nil {
			test.Fatalf("error while deployment creation: %v", err)
		}
	}
}
func getDeploymentTest(client *cmd.Client, namespace, deployment string, referenceDeployment model.Deployment) func(*testing.T) {
	return func(test *testing.T) {
		gainedDeployment, err := client.GetDeployment(namespace, deployment)
		if err != nil {
			test.Fatalf("error while getting deployment: %v", err)
		}
		if !reflect.DeepEqual(referenceDeployment.Containers, gainedDeployment.Containers) {
			test.Fatalf("gained deployment doesn't match reference deployment")
		}
	}
}

func getDeploymentListTest(client *cmd.Client, namespace string, referenceList []model.Deployment) func(*testing.T) {
	return func(test *testing.T) {
		gainedDeploymentList, err := client.GetDeploymentList(namespace)
		if err != nil {
			test.Fatalf("error while getting deployment: %v", err)
		}
		if !reflect.DeepEqual(referenceList, gainedDeploymentList) {
			test.Fatalf("gained deployment list doesn't match reference deployment list!")
		}
	}
}

func newFakeDeployment(test *testing.T, file string) model.Deployment {
	jsonData, err := ioutil.ReadFile(file)
	if err != nil {
		test.Fatalf("error while reading test data: %v", err)
	}
	var deployment model.Deployment
	if err := json.Unmarshal(jsonData, &deployment); err != nil {
		test.Fatalf("error while unmarshalling test response to deployment datastruct: %v", err)
	}
	return deployment
}

func newFakeResourceDeployment(test *testing.T) model.Deployment {
	return newFakeDeployment(test, "test_data/deployment.json")
}

func newFakeKubeAPIdeployment(test *testing.T) model.Deployment {
	return newFakeDeployment(test, "test_data/kubeAPIdeployment.json")
}

func newFakeResourceUpdateImage(test *testing.T) model.UpdateImage {
	jsonData, err := ioutil.ReadFile("test_data/update_image.json")
	if err != nil {
		test.Fatalf("error while reading test data: %v", err)
	}
	var updateImage model.UpdateImage
	if err := json.Unmarshal(jsonData, &updateImage); err != nil {
		test.Fatalf("error while unmarshalling test response to UpdateImage datastruct: %v", err)
	}
	return updateImage
}
