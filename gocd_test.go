package gocd

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
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

type TestStringSlice struct {
	got  string
	want string
}

// setup sets up a test HTTP server along with a gocd.Client that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// gocd client configured to use test server
	client = NewClient(
		server.URL,
		&Auth{
			Username: "mockUsername",
			Password: "mockPassword",
		},
		nil,
		false,
	)
}

// teardown closes the test HTTP server.
func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {

	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testAuth(t *testing.T, r *http.Request, want string) {

	if val, ok := r.Header["Authorization"]; ok {
		if got := val[0]; got != want {
			t.Errorf("Auth expected: %v, want %v", got, want)
		}
	} else {
		t.Error("'Authorization' header not found")
	}
}

func testStringInSlice(t *testing.T, s []string, e string) {
	for _, a := range s {
		if a == e {
			return
		}
	}
	t.Errorf("Expected '%s' in '%s'.", e, strings.Join(s, ","))
}

func testGotStringSlice(t *testing.T, got_want []TestStringSlice) {
	for index, test := range got_want {
		if test.got != test.want {
			t.Errorf("Expected '%s'. Got '%s' in '%d'", test.want, test.got, index)
		}
	}
}

func TestNewClient(t *testing.T) {

	c := NewClient(
		"http://ci.example.com/go",
		&Auth{
			Username: "mockUsername",
			Password: "mockPassword",
		},
		nil,
		false,
	)

	testGotStringSlice(t, []TestStringSlice{
		{c.BaseURL.String(), "http://ci.example.com/go"},
		{c.UserAgent, userAgent},
		{c.BaseURL.Host, "ci.example.com"},
	})

	if c.PipelineGroups == nil {
		t.Error("`PipelineGroups` missing from `client`.")
	}

	if c.Stages == nil {
		t.Error("`Stages` missing from `client`.")
	}

	if c.Jobs == nil {
		t.Error("`Jobs` missing from `client`.")
	}

	if c.PipelineTemplates == nil {
		t.Error("`PipelineTemplates` missing from `client`.")
	}
}

func TestDo(t *testing.T) {
	setup()
	defer teardown()

	type foo struct {
		A string
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, want %v", r.Method, m)
		}
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
