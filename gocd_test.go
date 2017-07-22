package gocd

import (
	"testing"
	"net/http"
	"io/ioutil"
	"strings"
)

func TestNewClient(t *testing.T) {

	mockClient := &http.Client{}
	mockClient.Transport = newMockTransport()

	client := NewClient("https://mock_endpoint/", mockClient)
	if client.BaseURL.Host != "mock_endpoint" {
		t.Error(
			"Expected: 'mock_endpoint'. ",
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

func newMockTransport() http.RoundTripper {
	return &transportMock{}
}

type transportMock struct{

}

func (t *transportMock) RoundTrip(req *http.Request) (*http.Response, error) {
	response := &http.Response{
		Header: make(http.Header),
		Request: req,
		StatusCode: http.StatusOK,
	}
	response.Header.Set("Content-Type", "application/json")
	response.Body = ioutil.NopCloser(strings.NewReader("response"))
	return response, nil
}
