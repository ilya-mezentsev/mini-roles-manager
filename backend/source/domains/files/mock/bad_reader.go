package mock

import "errors"

type ErroredReader struct {
}

func (e ErroredReader) Read([]byte) (n int, err error) {
	return 0, errors.New("some-error")
}
