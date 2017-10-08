package gocd

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	log "github.com/Sirupsen/logrus"
	"encoding/json"
)

// APIClientRequest helper struct to reduce amount of code.
type APIClientRequest struct {
	Method       string
	Path         string
	APIVersion   string
	RequestBody  interface{}
	ResponseType string
	ResponseBody interface{}
	Headers      map[string]string
}

// Handles any call to HEAD by returning whether or not we got a 2xx code.
func (c *Client) genericHeadAction(ctx context.Context, path string, apiversion string) (bool, *APIResponse, error) {
	_, resp, err := c.httpAction(ctx, &APIClientRequest{
		Method:       "HEAD",
		Path:         path,
		APIVersion:   apiversion,
		ResponseType: responseTypeJSON,
	})

	exists := resp.HTTP.StatusCode >= 300 || resp.HTTP.StatusCode < 200

	return exists, resp, err

}

func (c *Client) patchAction(ctx context.Context, r *APIClientRequest) (interface{}, *APIResponse, error) {
	r.Method = "PATCH"
	return c.httpAction(ctx, r)
}

func (c *Client) getAction(ctx context.Context, r *APIClientRequest) (interface{}, *APIResponse, error) {
	r.Method = "GET"
	return c.httpAction(ctx, r)
}

func (c *Client) postAction(ctx context.Context, r *APIClientRequest) (interface{}, *APIResponse, error) {
	r.Method = "POST"
	return c.httpAction(ctx, r)
}

func (c *Client) putAction(ctx context.Context, r *APIClientRequest) (interface{}, *APIResponse, error) {
	r.Method = "PUT"
	return c.httpAction(ctx, r)
}

// Returns a message from the DELETE action on the provided HTTP resource.
func (c *Client) deleteAction(ctx context.Context, path string, apiversion string) (string, *APIResponse, error) {
	a := StringResponse{}
	_, resp, err := c.httpAction(ctx, &APIClientRequest{
		Method:       "DELETE",
		Path:         path,
		APIVersion:   apiversion,
		ResponseType: responseTypeJSON,
		ResponseBody: &a,
	})

	return a.Message, resp, err
}

func (c *Client) httpAction(ctx context.Context, r *APIClientRequest) (interface{}, *APIResponse, error) {

	log.Debugf("HTTP Request")
	log.Debugf("%s %s", r.Method, r.Path)

	if r.ResponseType == "" {
		r.ResponseType = responseTypeJSON
	}

	var isVersioned bool
	var ver Versioned
	if ver, isVersioned = (r.RequestBody).(Versioned); isVersioned {
		if r.Headers == nil {
			r.Headers = map[string]string{}
		}
		r.Headers["If-Match"] = fmt.Sprintf("\"%s\"", ver.GetVersion())
	}

	// Build the request
	var reqBody interface{}
	if r.RequestBody != nil {
		reqBody = r.RequestBody
	} else {
		reqBody = nil
	}

	req, err := c.NewRequest(r.Method, r.Path, reqBody, r.APIVersion)
	if err != nil {
		return false, nil, err
	}

	if len(r.Headers) > 0 {
		for key, value := range r.Headers {
			req.HTTP.Header.Set(key, value)
		}
	}

	for header, values := range req.HTTP.Header {
		for _, value := range values {
			log.Debugf("%s: %s", header, value)
		}
	}
	if r.RequestBody != nil {
		log.Debug()
		log.Debug(r.RequestBody)
		log.Debug()
		log.Debug()
	}

	resp, err := c.Do(ctx, req, r.ResponseBody, r.ResponseType)
	if err != nil {
		log.Error(err)
		return r.ResponseBody, resp, err
	}

	if ver, isVersioned = (r.ResponseBody).(Versioned); isVersioned {
		parseVersions(resp.HTTP, ver)
	}

	if r.ResponseType == responseTypeJSON {
		log.Debug()
		log.Debug("Response")
		log.Debugf("%s %s", resp.HTTP.Proto, resp.HTTP.Status)
		for header, values := range resp.HTTP.Header {
			for _, value := range values {
				log.Debugf("%s: %s", header, value)
			}
		}
		log.Debug()
		b, _ := json.Marshal(r.ResponseBody)
		log.Debugf("%s", b)
	}

	return r.ResponseBody, resp, err
}

func parseVersions(response *http.Response, versioned Versioned) {
	etag := response.Header.Get("Etag")
	versioned.SetVersion(
		strings.Replace(etag, "\"", "", -1),
	)
}
