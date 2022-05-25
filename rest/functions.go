package rest

import (
	"fmt"

	"net/http"

	"github.com/emicklei/go-restful/v3"
)

type DeleteNumberRequest struct {
	Number *int
}

func (r *Service) DeleteNumberRequest(request *restful.Request, response *restful.Response) {

	fmt.Printf("number: %v", requestQuery.Number)

	err = validateDeleteNumberRequest(requestQuery)

	if err != nil {

		buildEndPointErrorResponse(response, http.StatusBadRequest, fmt.Sprintf("error:%s", err))

		return

	}

	a := len(r.tokens)

	r.tokens = removeFromArray(*requestQuery.Number, r.tokens)

	b := len(r.tokens)

	err = numberDosentExistDeleteNumber(a, b)

	if err != nil {

		buildEndPointErrorResponse(response, http.StatusBadRequest, fmt.Sprintf("error:%s", err))

	}

}

func validateDeleteNumberRequest(requestQuery DeleteNumberRequest) error {

	if requestQuery.Number == nil {

		return fmt.Errorf("invalid number: %v", *requestQuery.Number)

	}

	if *requestQuery.Number < 0 {

		return fmt.Errorf("invalid number: %v", *requestQuery.Number)

	}

	return nil

}
func removeFromArray(val int, array []string) []int {

	for i, w := range array {

		if val == w {

			array[i] = array[len(array)-1]

			return array[:len(array)-1] // asa se sterge o valoare din array

		}

	}

	return array

}
func numberDosentExistDeleteNumber(a int, b int) error {

	if a == b {

		return fmt.Errorf("number doesn't exist")

	}

	return nil

}
func (r *Service) GetToken(request *restful.Request, response *restful.Response) {

}
