package requests_tests

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"git.containerum.net/ch/kube-client/pkg/model"
)

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
