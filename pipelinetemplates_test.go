package gocd

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPipelineTemplateService_ListPipelineTemplates(t *testing.T) {

	setup()
	defer teardown()

	mux.HandleFunc("/api/admin/templates", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testAuth(t, r, mockAuthorization)
		j, _ := ioutil.ReadFile("test/resources/pipelinetemplates.0.json")
		fmt.Fprint(w, string(j))
	})

	templates, _, err := client.PipelineTemplates.ListPipelineTemplates(context.Background())

	if err != nil {
		t.Error(err)
	}

	if len((*templates)) != 1 {
		t.Error(
			"Expected template length: '1 ",
			"Got: '", (*templates)[0].Name, "'",
		)
	}

	if (*templates)[0].Name != "template0" {
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

func TestPipelineTemplateService_GetPipelineTemplates(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/admin/templates/template1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testAuth(t, r, mockAuthorization)
		j, _ := ioutil.ReadFile("test/resources/pipelinetemplates.1.json")
		fmt.Fprint(w, string(j))
	})
	template, _, err := client.PipelineTemplates.GetPipelineTemplate(
		context.Background(),
		"template1",
	)

	if err != nil {
		t.Error(err)
	}

	if template.Name != "template1" {
		t.Error(fmt.Printf("Expected template name: 'template1'. Got: '%s'.", template.Name))
	}

	if len(template.Stages) != 1 {
		t.Error(fmt.Printf("Expected 1 stage. Got '%d'.", len(template.Stages)))
	}
}
