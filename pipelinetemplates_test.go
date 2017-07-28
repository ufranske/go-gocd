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

	testGotStringSlice(t, []TestStringSlice{
		{(*templates)[0].Name, "template0"},
		{(*templates)[0].Pipelines[0], "up42"},
		{string(len((*templates))), "1"},
	})
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

	testGotStringSlice(t, []TestStringSlice{
		{template.Name, "template1"},
		{string(len(template.Stages)), "1"},
	})
}
