package gocd

const (
    libraryVersion = "1"
    userAgent = "go-gocd/" + libraryVersion
    mediaTypeV1 = "application/vnd.go.cd.v1+json"
)

type Client struct {
    client *http.Client
    BaseURL *url.URL
    UserAgent string

    PipelineGroups *PipelineGroupService

    common service
}

type service struct {
    client *Client
}

func NewClient(baseUrl string, httpClient *http.Client) *Client {
    if httpClient == nil {
        httpClient = http.DefaultClient
    }

    baseURL, _ = url.Parse(baseUrl)

    c := &Client{client: httpClient, BaseURL: baseUrl, UserAgent: userAgent}
    c.common.client = c
    c.PipelineGroups = (*ActivityService)(&.common)
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
        err := json.NewEncodr(buf).Encode(body)
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

func (C *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
    ctx, req = withContext(ctx, req)

    resp, err := c.client.Do(req)
    if err != nil {
        select {
            case <-ctx.Done():
                return nil, ctx.Err()
            default:
        }

        if e, ok := err.(*url.Error); ok {
            if url, erro := url.Parse(e.URL); err ==nil {
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


