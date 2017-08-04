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

	templates, _, err := client.PipelineTemplates.List(context.Background())

	if err != nil {
		t.Error(err)
	}

	if len(templates) != 1 {
		t.Errorf("Expected '1' template, got '%d'", len(templates))
	}

	testGotStringSlice(t, []TestStringSlice{
		{templates[0].Name, "template0"},
		{templates[0].Embedded.Pipelines[0].Name, "up42"},
	})
}

func TestPipelineTemplateService_GetPipelineTemplate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/admin/templates/template1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testAuth(t, r, mockAuthorization)
		j, _ := ioutil.ReadFile("test/resources/pipelinetemplate.0.json")
		fmt.Fprint(w, string(j))
	})

	template, _, err := client.PipelineTemplates.Get(
		context.Background(),
		"template1",
	)

	if err != nil {
		t.Error(err)
	}

	if len(template.Stages) != 1 {
		t.Errorf("Expected '1' template, got '%d'", len(template.Stages))
	}

	testGotStringSlice(t, []TestStringSlice{
		{template.Name, "template1"},
		{template.Stages[0].Name, "up42_stage"},
		{template.Stages[0].Approval.Type, "success"},
	})
}
