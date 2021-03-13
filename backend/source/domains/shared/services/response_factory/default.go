package response_factory

import "net/http"

type defaultResponse struct {
	data interface{}
}

func (r defaultResponse) HttpStatus() int {
	return http.StatusNoContent
}

func (r defaultResponse) ApplicationStatus() string {
	return statusOk
}

func (r defaultResponse) HasData() bool {
	return r.data != nil
}

func (r defaultResponse) GetData() interface{} {
	return r.data
}
