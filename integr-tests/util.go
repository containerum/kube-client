package test

import (
	"testing"

	kubeClient "git.containerum.net/ch/kube-client/pkg/client"
	"git.containerum.net/ch/kube-client/pkg/rest/re"
	"git.containerum.net/ch/kube-client/pkg/rest/remock"
)

func newMockClient(test *testing.T) *kubeClient.Client {
	client, err := kubeClient.NewClient(kubeClient.Config{
		APIurl:  "http://192.168.88.200",
		RestAPI: remock.NewMock(),
		User: kubeClient.User{
			Role: "user",
		},
	})
	if err != nil {
		test.Fatalf("error while client initialisation: %v", err)
	}
	return client
}
func newClient(test *testing.T) *kubeClient.Client {
	client, err := kubeClient.NewClient(
		kubeClient.Config{
			RestAPI: re.NewResty(),
			APIurl:  "http://192.168.88.200",
			User: kubeClient.User{
				Role: "user",
			},
		})
	if err != nil {
		test.Fatalf("error while creating client: %v", err)
	}
	return client
}
