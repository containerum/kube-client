package requests_tests

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"math/rand"
	"testing"
	"time"

	"git.containerum.net/ch/kube-client/cmd"
	"git.containerum.net/ch/kube-client/pkg/model"
)

const (
	resourceAddr = "http://192.168.88.200:1213"
	cubeAPIaddr  = "http://192.168.88.200:1214"
)

func newResourceClient(test *testing.T) *cmd.Client {
	client, err := cmd.CreateCmdClient(
		cmd.ClientConfig{
			ResourceAddr: resourceAddr,
			APIurl:       cubeAPIaddr,
			User: cmd.User{
				Role: "admin",
			},
		})
	if err != nil {
		test.Fatalf("error while creating client: %v", err)
	}
	client.SetHeader("X-User-ID", "20b616d8-1ea7-4842-b8ec-c6e8226fda5b")
	return client
}

func newCubeAPIClient(test *testing.T) *cmd.Client {
	client, err := cmd.CreateCmdClient(
		cmd.ClientConfig{
			ResourceAddr: resourceAddr,
			APIurl:       cubeAPIaddr,
			User: cmd.User{
				Role: "admin",
			},
		})
	if err != nil {
		test.Fatalf("error while creating client: %v", err)
	}
	return client
}

func newFakeResourceNamespaces(test *testing.T) []model.Namespace {
	return []model.Namespace{
		{
			TariffID: "4563e8c1-fb41-416a-9798-e949a2616260",
		},
	}
}

func createResourceNamespace(test *testing.T, client *cmd.Client, namespace model.Namespace) {
	resp, err := client.Request.
		SetBody(namespace).
		Post(resourceAddr + "/namespace")
	if err = mapErr
}

func newFakeDeployment(test *testing.T, file string) model.Deployment {
	var deployment model.Deployment
	loadTestJSONdata(test, file, &deployment)
	return deployment
}

func newFakeResourceDeployment(test *testing.T) model.ResourceDeployment {
	deployment := model.ResourceDeployment{
		Name:     "gateway",
		Replicas: 4,
		Labels:   map[string]string{},
		Containers: []model.Container{
			{
				Name: "proxy", Image: "nginx",
				Limits: model.Limits{CPU: "1", Memory: "256"},
				Ports: &[]model.Port{
					{Name: "Gate", Port: 1080, Protocol: model.TCP},
				},
				Env: &[]model.Env{
					{Name: "TEAPOT", Value: "TRUE"},
				},
				Volume: &[]model.Volume{
					{Name: "Store", MountPath: "/mount/store"},
				},
			},
		},
	}
	return deployment
}

func newFakeKubeAPIdeployment(test *testing.T) model.Deployment {
	return newFakeDeployment(test, "test_data/kubeAPIdeployment.json")
}

func newFakeResourceUpdateImage(test *testing.T) model.UpdateImage {
	var updateImage model.UpdateImage
	loadTestJSONdata(test, "test_data/update_image.json", &updateImage)
	return updateImage
}

func newFakeKubeAPInamespace(test *testing.T) model.Namespace {
	var namespace model.Namespace
	loadTestJSONdata(test, "test_data/kube_api_namespace.json", &namespace)
	return namespace
}

func newFakeResourceVolume(test *testing.T) []model.ResourceVolume {
	var volume []model.ResourceVolume
	loadTestJSONdata(test, "test_data/fake_volumes.json", &volume)
	return volume
}

func loadTestJSONdata(test *testing.T, file string, data interface{}) {
	jsonData, err := ioutil.ReadFile(file)
	if err != nil {
		test.Fatalf("error wgile reading from %q: %v", file, err)
	}
	err = json.Unmarshal(jsonData, data)
	if err != nil {
		test.Fatalf("error while unmarshalling data: %v", err)
	}
}

func newRandomName(size int64) string {
	buf := &bytes.Buffer{}
	encoder := base64.NewEncoder(base64.RawURLEncoding, buf)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	io.CopyN(encoder, rnd, (3*size)/4)
	return buf.String()
}

func newResourceUpdateDeployment(test *testing.T) model.ResourceDeployment {
	var deployment model.ResourceDeployment
	loadTestJSONdata(test, "test_data/update_deployment.json", &deployment)
	return deployment
}
