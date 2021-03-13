package response_factory

import "net/http"

type successResponse struct {
	defaultResponse
}

func (r successResponse) HttpStatus() int {
	return http.StatusOK
}
