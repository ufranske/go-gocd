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

type Response struct {
	Links    ResponseLinks `json:"_links,omitempty"`
	Embedded ResponseEmbedded `json:"_embedded"`
}
