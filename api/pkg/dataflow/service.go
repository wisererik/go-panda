package dataflow

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/micro/go-log"
	"github.com/micro/go-micro/client"
	//	"github.com/micro/go-micro/errors"
	"encoding/json"

	"github.com/opensds/go-panda/backend/proto"
	"github.com/opensds/go-panda/dataflow/proto"
	"github.com/opensds/go-panda/s3/proto"
	"golang.org/x/net/context"
)

const (
	backendService  = "backend"
	s3Service       = "s3"
	dataflowService = "dataflow"
)

type APIService struct {
	backendClient  backend.BackendService
	s3Client       s3.S3Service
	dataflowClient dataflow.DataFlowService
}

func NewAPIService(c client.Client) *APIService {
	return &APIService{
		backendClient:  backend.NewBackendService(backendService, c),
		s3Client:       s3.NewS3Service(s3Service, c),
		dataflowClient: dataflow.NewDataFlowService(dataflowService, c),
	}
}

func (s *APIService) GetBackend(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("id")
	log.Logf("Received request for backend details: %s", id)
	ctx := context.Background()
	res, err := s.backendClient.GetBackend(ctx, &backend.GetBackendRequest{Id: id})
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	log.Log("Get backend details successfully.")
	response.WriteEntity(res)
}

func (s *APIService) GetObject(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("id")
	log.Logf("Received request for object details: %s", id)
	ctx := context.Background()
	res, err := s.s3Client.GetObject(ctx, &s3.GetObjectRequest{Id: id})
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	log.Log("Get object details successfully.")
	response.WriteEntity(res)
}

func (s *APIService) GetPolicy(request *restful.Request, response *restful.Response) {
	name := request.PathParameter("name")
	log.Logf("Received request for policy[name=%s] details.", name)
	ctx := context.Background()
	res, err := s.dataflowClient.GetPolicy(ctx, &dataflow.GetPolicyRequest{Name: name})
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	//For debug -- begin
	log.Logf("Get policy reponse:%v",res)
	log.Logf("res.ErrCode:%d",res.ErrCode)
	jsons, errs := json.Marshal(res)
	if errs != nil {
		log.Logf(errs.Error())
	} else {
		log.Logf("res: %s.\n", jsons)
	}
	//For debug -- end

	log.Log("Get policy details successfully.")
	response.WriteEntity(res)
}

func (s *APIService) CreatePolicy(request *restful.Request, response *restful.Response) {
	//name := request.PathParameter("name")
	log.Logf("Received request for create policy.\n")
	ctx := context.Background()
	pol := dataflow.Policy{}
	err := request.ReadEntity(&pol)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		log.Logf("Read request body failed, err:%v.\n", err)
		return
	}

	//For debug --begin
	jsons, errs := json.Marshal(pol)
	if errs != nil {
		log.Logf(errs.Error())
	} else {
		log.Logf("Req body: %s.\n", jsons)
	}
	//For debug --end

	//name,_ := request.BodyParameter("name")
	//tenant,_ := request.BodyParameter("tenant")
	//desc,_ := request.BodyParameter("description")

	res, err := s.dataflowClient.CreatePolicy(ctx, &dataflow.CreatePolicyRequest{Pol: &pol})
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	log.Log("Get policy details successfully.")
	response.WriteEntity(res)
}

func (s *APIService) UpdatePolicy(request *restful.Request, response *restful.Response) {
	log.Logf("Received request for update policy.\n")
	ctx := context.Background()
	pol := dataflow.Policy{}
	err := request.ReadEntity(&pol)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		log.Logf("Read request body failed, err:%v.\n", err)
		return
	}

	//For debug --begin
	jsons, errs := json.Marshal(pol)
	if errs != nil {
		log.Logf(errs.Error())
	}else {
		log.Logf("Req body: %s.\n", jsons)
	}
	//For debug --end

	res, err := s.dataflowClient.UpdatePolicy(ctx, &dataflow.UpdatePolicyRequest{Pol:&pol})
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	log.Log("Update policy successfully.")
	response.WriteEntity(res)

	//For debug --begin
	jsons1, errs1 := json.Marshal(res)
	if errs1 != nil {
		log.Logf(errs1.Error())
	}else {
		log.Logf("Rsp body: %s.\n", jsons1)
	}
	//For debug --end
}

func (s *APIService) DeletePolicy(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("id")
	log.Logf("Received request for delete policy[id=%s] details.", id)
	ctx := context.Background()
	res, err := s.dataflowClient.DeletePolicy(ctx, &dataflow.DeletePolicyRequest{Id: id})
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	log.Log("Delete policy end, err = %d.", err)
	response.WriteEntity(res)
}

func (s *APIService) GetConnector(request *restful.Request, response *restful.Response) {
	name := request.PathParameter("name")
	log.Logf("Received request for connector[name=%s] details.", name)
	ctx := context.Background()
	res, err := s.dataflowClient.GetConnector(ctx, &dataflow.GetConnectorRequest{Name: name})
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	//For debug -- begin
	log.Logf("Get policy reponse:%v",res)
	log.Logf("res.ErrCode:%d",res.ErrCode)
	jsons, errs := json.Marshal(res)
	if errs != nil {
		log.Logf(errs.Error())
	}else {
		log.Logf("res: %s.\n", jsons)
	}
	//For debug -- end

	log.Log("Get policy details successfully.")
	response.WriteEntity(res)
}

func (s *APIService) CreateConnector(request *restful.Request, response *restful.Response) {
	//name := request.PathParameter("name")
	log.Logf("Received request for create connector.\n")
	ctx := context.Background()
	conn := dataflow.Connector{}
	err := request.ReadEntity(&conn)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		log.Logf("Read request body failed, err:%v.\n", err)
		return
	}

	//For debug --begin
	jsons, errs := json.Marshal(conn)
	if errs != nil {
		log.Logf(errs.Error())
	}else {
		log.Logf("Req body: %s.\n", jsons)
	}
	//For debug --end

	res, err := s.dataflowClient.CreateConnector(ctx, &dataflow.CreateConnectorRequest{Conn:&conn})
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	log.Log("Create connector successfully.")
	response.WriteEntity(res)
}

func (s *APIService) UpdateConnector(request *restful.Request, response *restful.Response) {
	log.Logf("Received request for update connector.\n")
	ctx := context.Background()
	conn := dataflow.Connector{}
	err := request.ReadEntity(&conn)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		log.Logf("Read request body failed, err:%v.\n", err)
		return
	}

	//For debug --begin
	jsons, errs := json.Marshal(conn)
	if errs != nil {
		log.Logf(errs.Error())
	}else {
		log.Logf("Req body: %s.\n", jsons)
	}
	//For debug --end

	res, err := s.dataflowClient.UpdateConnector(ctx, &dataflow.UpdateConnectorRequest{Conn:&conn})
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	log.Log("Update connector successfully.")
	response.WriteEntity(res)

	//For debug --begin
	jsons1, errs1 := json.Marshal(res)
	if errs1 != nil {
		log.Logf(errs1.Error())
	}else {
		log.Logf("Rsp body: %s.\n", jsons1)
	}
	//For debug --end
}

func (s *APIService) DeleteConnector(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("id")
	log.Logf("Received request for delete connector[id=%s] details.", id)
	ctx := context.Background()
	res, err := s.dataflowClient.DeleteConnector(ctx, &dataflow.DeleteConnectorRequest{Id: id})
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	log.Log("Delete connector end, err = %d.", err)
	response.WriteEntity(res)
}

func (s *APIService) GetPlan(request *restful.Request, response *restful.Response) {
	name := request.PathParameter("name")
	log.Logf("Received request for plan[name=%s] details.", name)
	ctx := context.Background()
	res, err := s.dataflowClient.GetPlan(ctx, &dataflow.GetPlanRequest{Name: name})
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	//For debug -- begin
	log.Logf("Get plan reponse:%v",res)
	log.Logf("res.ErrCode:%d",res.ErrCode)
	jsons, errs := json.Marshal(res)
	if errs != nil {
		log.Logf(errs.Error())
	}else {
		log.Logf("res: %s.\n", jsons)
	}
	//For debug -- end

	log.Log("Get plan details successfully.")
	response.WriteEntity(res)
}

func (s *APIService) CreatePlan(request *restful.Request, response *restful.Response) {
	//name := request.PathParameter("name")
	log.Logf("Received request for create plan.\n")
	ctx := context.Background()
	plan := dataflow.Plan{}
	err := request.ReadEntity(&plan)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		log.Logf("Read request body failed, err:%v.\n", err)
		return
	}

	//For debug --begin
	jsons, errs := json.Marshal(plan)
	if errs != nil {
		log.Logf(errs.Error())
	}else {
		log.Logf("Req body: %s.\n", jsons)
	}
	//For debug --end

	res, err := s.dataflowClient.CreatePlan(ctx, &dataflow.CreatePlanRequest{Plan:&plan})
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	log.Log("Create plan successfully.")
	response.WriteEntity(res)
}

func (s *APIService) UpdatePlan(request *restful.Request, response *restful.Response) {
	log.Logf("Received request for update plan.\n")
	ctx := context.Background()
	plan := dataflow.Plan{}
	err := request.ReadEntity(&plan)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		log.Logf("Read request body failed, err:%v.\n", err)
		return
	}

	//For debug --begin
	jsons, errs := json.Marshal(plan)
	if errs != nil {
		log.Logf(errs.Error())
	}else {
		log.Logf("Req body: %s.\n", jsons)
	}
	//For debug --end

	res, err := s.dataflowClient.UpdatePlan(ctx, &dataflow.UpdatePlanRequest{Plan:&plan})
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	log.Log("Update plan successfully.")
	response.WriteEntity(res)

	//For debug --begin
	jsons1, errs1 := json.Marshal(res)
	if errs1 != nil {
		log.Logf(errs1.Error())
	}else {
		log.Logf("Rsp body: %s.\n", jsons1)
	}
	//For debug --end
}

func (s *APIService) DeletePlan(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("id")
	log.Logf("Received request for delete plan[id=%s] details.", id)
	ctx := context.Background()
	res, err := s.dataflowClient.DeletePlan(ctx, &dataflow.DeletePlanRequest{Id: id})
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	response.WriteEntity(res)
}
