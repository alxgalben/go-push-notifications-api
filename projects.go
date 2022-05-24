package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/emicklei/go-restful/v3"
)

type AddProject struct {
	Response []string `json:"response"`
}

/*
######################################################################################################
######################################  ADD TO PROJECT  ############################################
######################################################################################################
*/

func readAddProjectRequest(request *restful.Request) (requestQuery GetProjectRequest, err error) {

func (r *Service) AddProjectElement(request *restful.Request, response *restful.Response) {

	requestQuery, err := readAddProjectRequest(request)
	if err != nil {
		buildEndPointErrorResponse(response, http.StatusBadRequest, fmt.Sprintf("error:%s", err))
		return
	}

	err = r.validateGetProjectRequest(requestQuery)
	if err != nil {
		buildEndPointErrorResponse(response, http.StatusBadRequest, fmt.Sprintf("error:%s", err))
		return
	}

	r.Projects = append(r.Projects, *requestQuery.Val)
	return

}
