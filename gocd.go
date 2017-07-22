package gocd

import (
	"net/http"
	"net/url"
	"io"
	"bytes"
	"encoding/json"
)

const (
	libraryVersion = "1"
	userAgent = "go-gocd/" + libraryVersion
	mediaTypeV1 = "application/vnd.go.cd.v1+json"
)

type Client struct {
	client            *http.Client
	BaseURL           *url.URL
	UserAgent         string

	PipelineGroups    *PipelineGroupsService
	Stages            *StagesService
	Jobs              *JobsService
	PipelineTemplates *PipelineTemplatesService

	common            service
}

type service struct {
	client *Client
}

func NewClient(gocdBaseUrl string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(gocdBaseUrl)

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}
	c.common.client = c
	c.PipelineGroups = (*PipelineGroupsService)(&c.common)
	c.Stages = (*StagesService)(&c.common)
	c.Jobs = (*JobsService)(&c.common)
	c.PipelineTemplates = (*PipelineTemplatesService)(&c.common)
	return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)

	if err != nil {
		return nil, err
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
	req.Header.Set("Accept", mediaTypeV1)
	req.Header.Set("User-Agent", c.UserAgent)
	return req, nil
}

type Response struct {
	*http.Response
}

func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {

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

	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}

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

// CheckResponse checks the API response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range or equal to 202 Accepted.
// API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse. Any other
// response body will be silently ignored.
//
// The error type will be *RateLimitError for rate limit exceeded errors,
// *AcceptedError for 202 Accepted status codes,
// and *TwoFactorAuthError for two-factor authentication errors.
func CheckResponse(r *http.Response) error {
	//if r.StatusCode == http.StatusAccepted {
	//	return &AcceptedError{}
	//}
	//if c := r.StatusCode; 200 <= c && c <= 299 {
	//	return nil
	//}
	//errorResponse := &ErrorResponse{Response: r}
	//data, err := ioutil.ReadAll(r.Body)
	//if err == nil && data != nil {
	//	json.Unmarshal(data, errorResponse)
	//}
	//switch {
	//case r.StatusCode == http.StatusUnauthorized && strings.HasPrefix(r.Header.Get(headerOTP), "required"):
	//	return (*TwoFactorAuthError)(errorResponse)
	//case r.StatusCode == http.StatusForbidden && r.Header.Get(headerRateRemaining) == "0" && strings.HasPrefix(errorResponse.Message, "API rate limit exceeded for "):
	//	return &RateLimitError{
	//		Rate:     parseRate(r),
	//		Response: errorResponse.Response,
	//		Message:  errorResponse.Message,
	//	}
	//case r.StatusCode == http.StatusForbidden && errorResponse.DocumentationURL == "https://developer.github.com/v3#abuse-rate-limits":
	//	abuseRateLimitError := &AbuseRateLimitError{
	//		Response: errorResponse.Response,
	//		Message:  errorResponse.Message,
	//	}
	//	if v := r.Header["Retry-After"]; len(v) > 0 {
	//		// According to GitHub support, the "Retry-After" header value will be
	//		// an integer which represents the number of seconds that one should
	//		// wait before resuming making requests.
	//		retryAfterSeconds, _ := strconv.ParseInt(v[0], 10, 64) // Error handling is noop.
	//		retryAfter := time.Duration(retryAfterSeconds) * time.Second
	//		abuseRateLimitError.RetryAfter = &retryAfter
	//	}
	//	return abuseRateLimitError
	//default:
	//	return errorResponse
	//}
	return nil
}
