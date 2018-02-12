package requests_test

import (
	"encoding/json"
	"math/rand"
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
	client, err := cmd.CreateCmdClient(cmd.User{
		Role: "admin",
	})
	if err != nil {
		test.Fatalf("error while creating client: %v", err)
	}
	test.Run("deployment test", deploymentCreationTest(client))
}

func deploymentCreationTest(client *cmd.Client) func(*testing.T) {
	return func(test *testing.T) {
		err := client.CreateDeployment(testNamespace, newFakeDeployment(test))
		if err != nil {
			test.Fatalf("error while deployment creation: %v", err)
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
