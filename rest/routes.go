package rest

import (
	"net/http"

	"github.com/emicklei/go-restful/v3"
)

const (
	token = "AAAA-fvtI1A:APA91bGQrxSmJA2YskZqQFkkG89UYUbUedAF9RaNaz2kYmZ1uAg-fQkyyQEFbxTbibDEcdCGM_1OE4NeNGY4l7aeShyJWxePGk8koo0rUfRbuVwMW_XBFaC083G0yHJ4FhX9PYyBhvvJ"
)

type Token struct {
	Token string `json:"token"`
}
type Subscriber struct {
	Username     string
	Token        string
	Applications []Application
}

type Application struct {
	AppId          string
	ExpirationDate date.Date
}

type Service struct {
	container *restful.Container
	tokens    []string
}

func NewService() (*Service, error) {
	r := &Service{
		container: restful.NewContainer(),
		tokens:    make([]string, 0),
	}

	r.container.Add(r.buildRoutes())

	if err := bootstrapSwagger(r.container); err != nil {
		return nil, err
	}

	return r, nil
}

func (r *Service) Container() *restful.Container {
	return r.container
}

func (r *Service) buildRoutes() *restful.WebService {
	ws := new(restful.WebService)

	ws.Path("/internship-api").
		Route(ws.GET("/health").
			Operation("health check").
			To(func(request *restful.Request, response *restful.Response) {
				_ = response.WriteErrorString(http.StatusOK, "internship-api is up and running!")
			}))

	ws.Route(ws.GET("/vapid").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), token).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), EndpointErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), EndpointErrorResponse{}).
		To(r.GetToken).
		Writes(map[string]string{}))

	ws.Route(ws.PUT("/blackList/{"+BlackListParameter+"}").
		Param(restful.QueryParameter("val", "numar").DataType("int").Required(true)).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), GetBlacklistResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), EndpointErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), EndpointErrorResponse{}).
		To(r.AddBlacklistElement).
		Writes(map[string]string{}))

	ws.Route(ws.DELETE("/subscriber/{"+Subscriber+"}").
		Param(restful.QueryParameter("subscriber").DataType("struct").Required(true)).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), token).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), token).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), token).
		To(r.DeleteNumberRequest).
		Writes(map[string]string{}))

	ws.Route(ws.GET("/blackList").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), GetFibonacciResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), EndpointErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), EndpointErrorResponse{}).
		To(r.printBlacklist).
		Writes(map[string]string{}))

	return ws
}
