package requests_test

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
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
	{
		fakeResourceDeployment := newFakeResourceDeployment(test)
		test.Run("deployment creaton test",
			deploymentCreationTest(client, testNamespace, fakeResourceDeployment))
	}
	{
		test.Run("get deployment test",
			getDeploymentTest(client, kubeAPItestNamespace, kubeAPItestDeployment))
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

func getDeploymentTest(client *cmd.Client, namespace, deployment string) func(*testing.T) {
	return func(test *testing.T) {
		_, err := client.GetDeployment(namespace, deployment)
		if err != nil {
			test.Fatalf("error while getting deployment: %v", err)
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
