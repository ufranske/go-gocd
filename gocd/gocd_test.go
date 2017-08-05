package gocd

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
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

func testMethod(t *testing.T, r *http.Request, want string) {
	assert.Equal(t, r.Method, want, "Unexpected HTTP method")
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
		testMethod(t, r, "GET")
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
