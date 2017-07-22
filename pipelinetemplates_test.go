package gocd

import (
	"testing"
	"net/http"
	"context"
)

func TestPipelineTemplateService_ListPipelineTemplates(t *testing.T) {

	setup(transportMock{
		"tests/pipelinetemplates/0.response.json",
	})
	defer teardown()

	mux.HandleFunc("api/admin/templates", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testAuth(t, r, "mockUsername:mockPassword")
	})

	templates, _, err := client.PipelineTemplates.ListPipelines(context.Background())

	if err != nil {
		t.Error(err)
	}

	if len((*templates)) != 1 {
		t.Error(
			"Expected template length: '1 ",
			"Got: '", (*templates)[0].Name, "'",
		)
	}

	if (*templates)[0].Name != "template1" {
		t.Error(
			"Expected: 'template1. ",
			"Got: '", (*templates)[0].Name, "'",
		)
	}

	if (*templates)[0].Pipelines[0] != "up42" {
		t.Error(
			"Expected: 'up42. ",
			"Got: '", (*templates)[0].Pipelines[0], "'",
		)
	}
}
