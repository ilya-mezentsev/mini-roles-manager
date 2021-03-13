package response_factory

import "net/http"

type forbiddenErrorResponse struct {
	defaultResponse
}

func (r forbiddenErrorResponse) ApplicationStatus() string {
	return statusError
}

func (r forbiddenErrorResponse) HttpStatus() int {
	return http.StatusForbidden
}
