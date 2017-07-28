package gocd

import (
	"net/http"
)

type APIResponse struct {
	*http.Response
}

func newResponse(r *http.Response) *APIResponse {
	response := &APIResponse{Response: r}
	return response
}