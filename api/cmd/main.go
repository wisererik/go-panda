package main

import (
	"github.com/emicklei/go-restful"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-web"
	service "github.com/opensds/go-panda/api/pkg"

	"github.com/micro/go-log"
	//	_ "github.com/micro/go-plugins/client/grpc"
)

const (
	serviceName = "go.panda"
)

func main() {
	webService := web.NewService(
		web.Name(serviceName),
		web.Version("v0.1.0"),
	)

	webService.Init()
	handler := service.NewAPIService(client.DefaultClient)

	wc := restful.NewContainer()
	ws := new(restful.WebService)

	ws.
		Path("/v1").
		Doc("OpenSDS Multi-Cloud API").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/backends/{id}").To(handler.GetBackend)).
		Doc("Get backend details")

	ws.Route(ws.GET("/objects/{id}").To(handler.GetObject)).
		Doc("Get object details")

	ws.Route(ws.GET("/policies/{name}").To(handler.GetPolicy)).
		Doc("Get policy details")
	ws.Route(ws.POST("/policies/{name}").To(handler.CreatePolicy)).
		Doc("Create policy")
	ws.Route(ws.PUT("/policies/{id}").To(handler.UpdatePolicy)).
		Doc("Update policy")
	ws.Route(ws.DELETE("/policies/{id}").To(handler.DeletePolicy)).
		Doc("Delete policy")

	ws.Route(ws.GET("/connector/{name}").To(handler.GetConnector)).
		Doc("Get connector details")
	ws.Route(ws.POST("/connector/{name}").To(handler.CreateConnector)).
		Doc("Create connector")
	ws.Route(ws.PUT("/connector/{id}").To(handler.UpdateConnector)).
		Doc("Update connector")
	ws.Route(ws.DELETE("/connector/{id}").To(handler.DeleteConnector)).
		Doc("Delete connector")

	ws.Route(ws.GET("/plan/{name}").To(handler.GetPlan)).
		Doc("Get plan details")
	ws.Route(ws.POST("/plan/{name}").To(handler.CreatePlan)).
		Doc("Create plan")
	ws.Route(ws.PUT("/plan/{id}").To(handler.UpdatePlan)).
		Doc("Update plan")
	ws.Route(ws.DELETE("/plan/{id}").To(handler.DeletePlan)).
		Doc("Delete plan")

	wc.Add(ws)
	webService.Handle("/", wc)
	if err := webService.Run(); err != nil {
		log.Fatal(err)
	}

}
