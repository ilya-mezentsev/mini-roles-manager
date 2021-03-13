package response_factory

import "net/http"

type serverErrorResponse struct {
	defaultResponse
}

func (r serverErrorResponse) ApplicationStatus() string {
	return statusError
}

func (r serverErrorResponse) HttpStatus() int {
	return http.StatusInternalServerError
}
