package response_factory

import "net/http"

type clientErrorResponse struct {
	defaultResponse
}

func (r clientErrorResponse) ApplicationStatus() string {
	return statusError
}

func (r clientErrorResponse) HttpStatus() int {
	return http.StatusBadRequest
}
