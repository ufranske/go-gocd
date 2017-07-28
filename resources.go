package gocd

import (
	"net/http"
	"net/url"
)

type APIResponse struct {
	*http.Response
}

func newResponse(r *http.Response) *APIResponse {
	response := &APIResponse{Response: r}
	return response
}

//go:generate gocd-response-links -type=ResponseLinks -output=responselinks_responselinks.go
type ResponseLinks struct {
	Doc    *url.URL
	Find   *url.URL
	Job    *url.URL
	Latest *url.URL
	Next   *url.URL
	Oldest *url.URL
	Self   *url.URL
}