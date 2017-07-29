package gocd

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const (
	libraryVersion = "1"
	userAgent      = "go-gocd/" + libraryVersion
	apiV1          = "application/vnd.go.cd.v1+json"
	apiV2          = "application/vnd.go.cd.v2+json"
	apiV3          = "application/vnd.go.cd.v3+json"
	apiV4          = "application/vnd.go.cd.v4+json"
)

type APIResponse struct {
	*http.Response
}

func newResponse(r *http.Response) *APIResponse {
	response := &APIResponse{Response: r}
	return response
}

type DeleteResponse struct {
	Message string `json:"message"`
}

type ClientInterface interface{}

type Client struct {
	client    *http.Client
	BaseURL   *url.URL
	UserAgent string
	Auth      *Auth

	Agents            *AgentsService
	PipelineGroups    *PipelineGroupsService
	Stages            *StagesService
	Jobs              *JobsService
	PipelineTemplates *PipelineTemplatesService

	common service
	cookie string
}

type service struct {
	client *Client
}

type Auth struct {
	Username string
	Password string
}

func NewClient(gocdBaseUrl string, auth *Auth, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(gocdBaseUrl)

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent, Auth: auth}
	c.common.client = c
	c.Agents = (*AgentsService)(&c.common)
	c.PipelineGroups = (*PipelineGroupsService)(&c.common)
	c.Stages = (*StagesService)(&c.common)
	c.Jobs = (*JobsService)(&c.common)
	c.PipelineTemplates = (*PipelineTemplatesService)(&c.common)
	return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}, apiVersion string) (*http.Request, error) {
	rel, err := url.Parse("api/" + urlStr)

	if err != nil {
		return nil, err
	}

	if apiVersion == "" {
		apiVersion = apiV1
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", apiVersion)
	req.Header.Set("User-Agent", c.UserAgent)

	if c.cookie == "" {
		req.SetBasicAuth(c.Auth.Username, c.Auth.Password)
	} else {
		req.Header.Set("Cookie", c.cookie)
	}

	return req, nil
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*APIResponse, error) {

	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		if e, ok := err.(*url.Error); ok {
			if url, err := url.Parse(e.URL); err == nil {
				e.URL = sanitizeURL(url).String()
				return nil, e
			}
		}

		return nil, err
	}

	response := newResponse(resp)

	//err = CheckResponse(resp)
	//if err != nil {
	//	return response, err
	//}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
	}

	return response, err

}

// sanitizeURL redacts the client_secret parameter from the URL which may be
// exposed to the user.
func sanitizeURL(uri *url.URL) *url.URL {
	if uri == nil {
		return nil
	}
	params := uri.Query()
	if len(params.Get("client_secret")) > 0 {
		params.Set("client_secret", "REDACTED")
		uri.RawQuery = params.Encode()
	}
	return uri
}

// addOptions adds the parameters in opt as URL query parameters to s. opt
// must be a struct whose fields may contain "url" tags.
//func addOptions(s string, opt interface{}) (string, error) {
func addOptions(s string) (string, error) {
	//v := reflect.ValueOf(opt)
	//if v.Kind() == reflect.Ptr && v.IsNil() {
	//	return s, nil
	//}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	return u.String(), nil
}
