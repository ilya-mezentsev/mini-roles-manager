package mock

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func MustCreateRequestWithBody(u interface{}) *http.Request {
	us, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"localhost:8080",
		bytes.NewBuffer(us),
	)

	return req
}

func CreateRequestWithCookie(key, value string) *http.Request {
	req := httptest.NewRequest(
		http.MethodGet,
		"localhost:8080",
		nil,
	)

	req.AddCookie(&http.Cookie{
		Name:  key,
		Value: value,
	})

	return req
}

func CreateRequestWithHeader(key, value string) *http.Request {
	req := httptest.NewRequest(
		http.MethodGet,
		"localhost:8080",
		nil,
	)

	req.Header.Set(key, value)

	return req
}

func CreateSimpleRequest() *http.Request {
	return httptest.NewRequest(
		http.MethodGet,
		"localhost:8080",
		nil,
	)
}
