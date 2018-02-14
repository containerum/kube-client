package requests_tests

import (
	"testing"

	"git.containerum.net/ch/kube-client/pkg/cmd"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	resourceTestVolume = "n1-volume"
)

func TestVolume(test *testing.T) {
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
	Convey("Test resource API methods", test, func() {
		referenceVolume := newFakeResourceVolume(test)
		Convey("get volume", func() {
			gainedVolume, err := client.GetVolume(resourceTestVolume, nil)
			So(err, ShouldBeNil)
			So(gainedVolume, ShouldResemble, referenceVolume)
		})
	})
}
