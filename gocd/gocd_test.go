package gocd

import (
	"bytes"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

const (
	mockAuthorization = "Basic bW9ja1VzZXJuYW1lOm1vY2tQYXNzd29yZA=="
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the GitHub client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

type EqualityTest struct {
	got    string
	wanted string
}

// setup sets up a test HTTP server along with a gocd.Client that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// gocd client configured to use test server
	client = NewClient(&Configuration{
		Server:   server.URL,
		Username: "mockUsername",
		Password: "mockPassword",
	}, nil)
}

// teardown closes the test HTTP server.
func teardown() {
	server.Close()
}

func TestCheckResponse(t *testing.T) {
	t.Run("ValidHTTP", testCheckResponseValid)
	t.Run("FailHTTP", testCheckResponseInvalid)
}

type closingbuffer struct {
	*bytes.Buffer
}

func (cb *closingbuffer) Close() error {
	return nil
}

func testCheckResponseInvalid(t *testing.T) {
	var rc1, rc2 io.ReadCloser

	cb1 := &closingbuffer{bytes.NewBufferString("Hi!")}
	cb2 := &closingbuffer{bytes.NewBufferString("Hi!")}
	rc1 = cb1
	rc2 = cb2

	err := CheckResponse(&http.Response{
		StatusCode: 199,
		Status:     "Failed",
		Body:       rc1,
	})
	assert.NotNil(t, err)

	err = CheckResponse(&http.Response{
		StatusCode: 400,
		Status:     "Failed",
		Body:       rc2,
	})
	assert.NotNil(t, err)

}

func testCheckResponseValid(t *testing.T) {
	err := CheckResponse(&http.Response{
		StatusCode: 200,
	})
	assert.Nil(t, err)
}

func testAuth(t *testing.T, r *http.Request, want string) {
	assert.Contains(t, r.Header, "Authorization")
	assert.Contains(t, r.Header["Authorization"], want)
}

func TestNewClient(t *testing.T) {

	c := NewClient(&Configuration{
		Server:   server.URL,
		Username: "mockUsername",
		Password: "mockPassword",
	}, nil)

	// Make sure expected values are present.
	for _, attribute := range []EqualityTest{
		{c.BaseURL.String(), server.URL},
		{c.UserAgent, userAgent},
	} {
		assert.Equal(t, attribute.got, attribute.wanted)
	}

	// Make sure values expected to have nil, have nil.
	for _, attribute := range []interface{}{
		c.PipelineGroups,
		c.Stages,
		c.Jobs,
		c.PipelineTemplates,
	} {
		assert.NotNil(t, attribute)
	}
}

func TestDo(t *testing.T) {
	setup()
	defer teardown()

	type foo struct {
		A string
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Unexpected HTTP method")
		fmt.Fprint(w, `{"A":"a"}`)
	})

	req, _ := client.NewRequest("GET", "/", nil, "api-version")
	body := new(foo)
	client.Do(context.Background(), req, body)

	want := &foo{"a"}
	if !reflect.DeepEqual(body, want) {
		t.Errorf("Response body = %v, want %v", body, want)
	}
}
