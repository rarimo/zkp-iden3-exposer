package helpers

import (
	"net/http"
)

type FetcherError struct {
	name     string
	response *http.Response
}

func (e *FetcherError) Error() string {
	return "FetcherError"
}

func NewFetcherError(resp *http.Response) *FetcherError {
	return &FetcherError{
		name:     "FetcherError",
		response: resp,
	}
}
