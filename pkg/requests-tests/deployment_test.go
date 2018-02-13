package requests_test

import (
	"encoding/json"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"git.containerum.net/ch/kube-client/pkg/cmd"
	"git.containerum.net/ch/kube-client/pkg/model"
)

const (
	testNamespace = "test-namespace"
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
	fakeDeployment := newFakeDeployment(test)
	test.Run("deployment creaton test",
		deploymentCreationTest(client, testNamespace, fakeDeployment))
	test.Run("get deployment test",
		getDeploymentTest(client, testNamespace, fakeDeployment))
}

func deploymentCreationTest(client *cmd.Client, namespace string, deployment model.Deployment) func(*testing.T) {
	return func(test *testing.T) {
		err := client.CreateDeployment(namespace, deployment)
		if err != nil {
			test.Fatalf("error while deployment creation: %v", err)
		}
	}
}

func getDeploymentTest(client *cmd.Client, namespace string, deployment model.Deployment) func(*testing.T) {
	return func(test *testing.T) {
		createdDeployment, err := client.GetDeployment(namespace, deployment.Name)
		if err != nil {
			test.Fatalf("error while getting deployment: %v", err)
		}
		if !reflect.DeepEqual(createdDeployment, deployment) {
			test.Fatalf("created deployment doesn't match provided data!")
		}
	}
}

func newFakeDeployment(test *testing.T) model.Deployment {
	jsonStr := `{
		"containers": [
			{
				"image": "nginx", 
				"name": "first", 
				"resources": 
					{
						"requests": 
							{"cpu": "100m", "memory": "128Mi"}
					}, 
				"env": [
					{
						"value": "world", "name": "hello"
					}
				], 
				"commands": [],
				"ports": [{"containerPort": 10000}],
				"volumeMounts": [
					{
						"name": "default-volume", "mountPath": "blabla", "subPath": "home"
					}
				]
			}
		], 
		"labels": {"name": "value"}, 
		"name": "nginx", 
		"replicas": 1
	}`
	var deployment model.Deployment
	if err := json.Unmarshal([]byte(jsonStr), &deployment); err != nil {
		test.Fatalf("error while unmarshalling test response to deployment datastruct: %v", err)
	}
	return deployment
}
