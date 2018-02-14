package requests_tests

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"git.containerum.net/ch/kube-client/pkg/cmd"

	"git.containerum.net/ch/kube-client/pkg/model"
)

func newResourceClient(test *testing.T) *cmd.Client {
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
	client.SetHeader("X-User-ID", "20b616d8-1ea7-4842-b8ec-c6e8226fda5b")
	return client
}

func newFakeDeployment(test *testing.T, file string) model.Deployment {
	var deployment model.Deployment
	loadTestJSONdata(test, file, &deployment)
	return deployment
}

func newFakeResourceDeployment(test *testing.T) model.Deployment {
	return newFakeDeployment(test, "test_data/deployment.json")
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

func newFakeResourceVolume(test *testing.T) model.ResourceVolume {
	var volume model.ResourceVolume
	loadTestJSONdata(test, "test_data/get_volume.json", &volume)
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
