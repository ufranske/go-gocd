package gocd

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPipelineTemplateService_ListPipelineTemplates(t *testing.T) {

	setup()
	defer teardown()

	mux.HandleFunc("/api/admin/templates", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Unexpected HTTP method")
		testAuth(t, r, mockAuthorization)
		j, _ := ioutil.ReadFile("test/resources/pipelinetemplates.0.json")
		fmt.Fprint(w, string(j))
	})

	templates, _, err := client.PipelineTemplates.List(context.Background())

	assert.Nil(t, err)
	assert.Len(t, templates, 1)

	for _, attribute := range []EqualityTest{
		{templates[0].Name, "template0"},
		{templates[0].Embedded.Pipelines[0].Name, "up42"},
	} {
		assert.Equal(t, attribute.got, attribute.wanted)
	}
}

func TestPipelineTemplateService_GetPipelineTemplate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/admin/templates/template1", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Unexpected HTTP method")
		testAuth(t, r, mockAuthorization)
		j, _ := ioutil.ReadFile("test/resources/pipelinetemplate.0.json")
		fmt.Fprint(w, string(j))
	})

	template, _, err := client.PipelineTemplates.Get(
		context.Background(),
		"template1",
	)

	assert.Nil(t, err)
	assert.Len(t, template.Stages, 1)

	for _, attribute := range []EqualityTest{
		{template.Name, "template1"},
		{template.Stages[0].Name, "up42_stage"},
		{template.Stages[0].Approval.Type, "success"},
	} {
		assert.Equal(t, attribute.got, attribute.wanted)
	}
}
