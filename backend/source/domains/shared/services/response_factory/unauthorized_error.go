package response_factory

import "net/http"

type unauthorizedErrorResponse struct {
	defaultResponse
}

func (r unauthorizedErrorResponse) ApplicationStatus() string {
	return statusError
}

func (r unauthorizedErrorResponse) HttpStatus() int {
	return http.StatusUnauthorized
}
