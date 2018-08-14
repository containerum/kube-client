package reqtests

import (
	"testing"

	"github.com/containerum/kube-client/pkg/model"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	kubeAPItestService = "ch-glusterfs"
)

func TestService(test *testing.T) {
	client := newMockClient(test)
	Convey("Test Kube API methods", test, func() {
		Convey("get service", func() {
			_, err := client.GetService(kubeAPItestNamespace, kubeAPItestService)
			So(err, ShouldBeNil)
		})
		Convey("get service list", func() {
			_, err := client.GetServiceList(kubeAPItestNamespace)
			So(err, ShouldBeNil)
		})
	})
	Convey("Test resource manager methods", test, func() {
		nsList, err := client.GetNamespaceList()
		So(err, ShouldBeNil)
		So(len(nsList.Namespaces), ShouldBeGreaterThan, 0)
		namespace := nsList.Namespaces[0]
		testService := model.Service{
			Name:   "sophocles",
			Deploy: "mappet",
			IPs:    nil,
			Domain: "",
			Ports: []model.ServicePort{
				{
					Name:       "wasteland_entry",
					TargetPort: 11666,
					Protocol:   model.UDP,
				},
			},
		}
		createdService, err := client.CreateService(namespace.ID, testService)
		So(err, ShouldBeNil)
		So(createdService, ShouldResemble, testService)

		serviceList, err := client.GetServiceList(namespace.ID)
		So(err, ShouldBeNil)
		So(len(serviceList.Services), ShouldBeGreaterThan, 0)

		gainedService, err := client.GetService(namespace.ID, createdService.Name)
		So(err, ShouldBeNil)
		So(gainedService, ShouldResemble, createdService)

		updatedService, err := client.UpdateService(namespace.ID, testService)
		So(err, ShouldBeNil)
		So(updatedService, ShouldResemble, testService)

		So(client.DeleteService(namespace.ID, testService.Name), ShouldBeNil)
	})
}
