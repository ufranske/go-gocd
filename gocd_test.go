package gocd

import (
	"testing"
	"net/http"
	"io/ioutil"
	"strings"
	"net/http/httptest"
	"fmt"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the GitHub client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

// setup sets up a test HTTP server along with a gocd.Client that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup(s transportMock) {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	mockClient := &http.Client{}
	mockClient.Transport = newMockTransport(s)

	// github client configured to use test server
	client = NewClient(server.URL, mockClient)
}

// teardown closes the test HTTP server.
func teardown() {
	server.Close()
}

func newMockTransport(r transportMock) http.RoundTripper {
	responseData, _ := ioutil.ReadFile(r.Response)

	return &transportMock{
		Response: string(responseData),
	}
}

type transportMock struct {
	Response string
}

func testMethod(t *testing.T, r *http.Request, want string) {
	fmt.Print(r.Header)
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testAuth(t *testing.T, r *http.Request, want string) {
	fmt.Print(r.Header)
	if got := r.Header["WWW-Authenticate"][1]; got != want {
		t.Errorf("Auth expected: %v, want %v", got, want)
	}
}

func (t *transportMock) RoundTrip(req *http.Request) (*http.Response, error) {
	response := &http.Response{
		Header: make(http.Header),
		Request: req,
		StatusCode: http.StatusOK,
	}
	response.Header.Set("Content-Type", "application/json")
	response.Body = ioutil.NopCloser(strings.NewReader(t.Response))
	return response, nil
}

func TestNewClient(t *testing.T) {

	mockClient := &http.Client{}
	mockClient.Transport = newMockTransport(transportMock{
		"tests/pipelinetemplates/0.response.json",
	})

	client := NewClient("https://ci.example.com/go", mockClient)
	if client.BaseURL.Host != "ci.example.com" {
		t.Error(
			"Expected: 'ci.example.com'. ",
			"Got: '", client.BaseURL, "'",
		)
	}

	if client.PipelineGroups == nil {
		t.Error("`PipelineGroups` missing from `client`.")
	}

	if client.Stages == nil {
		t.Error("`Stages` missing from `client`.")
	}

	if client.Jobs == nil {
		t.Error("`Jobs` missing from `client`.")
	}

	if client.PipelineTemplates == nil {
		t.Error("`PipelineTemplates` missing from `client`.")
	}
}

