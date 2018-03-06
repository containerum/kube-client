package test

import (
	"testing"

	"git.containerum.net/ch/kube-client/pkg/model"
)

func TestDeployment(test *testing.T) {
	client := newClient(test)
	_, err := client.Login(model.Login{
		Username: "helpik94@yandex.ru",
		Password: "12345678",
	})
	if err != nil {
		test.Fatalf("error while login: %v", err)
	}
}
