package gocd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPipelineTemplate(t *testing.T) {
	setup()
	defer teardown()

	t.Run("List", testListPipelineTemplates)
	t.Run("Get", testGetPipelineTemplate)
	t.Run("Delete", testDeletePipelineTemplate)
	t.Run("RemoveLinks", tesPipelineTemplateRemoveLinks)
	t.Run("Pipelines", testPipelineTemplatePipelines)
}

func testPipelineTemplatePipelines(t *testing.T) {
	p := []*Pipeline{}
	pt := PipelineTemplate{Embedded: &embeddedPipelineTemplate{Pipelines: p}}

	assert.Exactly(t, p, pt.Pipelines())
}

func tesPipelineTemplateRemoveLinks(t *testing.T) {
	pt := PipelineTemplate{Links: &PipelineTemplateLinks{}}
	assert.NotNil(t, pt.Links)
	pt.RemoveLinks()
	assert.Nil(t, pt.Links)
}

func testDeletePipelineTemplate(t *testing.T) {

	b, err := json.Marshal(map[string]string{
		"message": "The template 'template2' was deleted successfully.",
	})
	if err != nil {
		t.Error(err)
	}

	mux.HandleFunc("/api/admin/templates/template2", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "DELETE", "Unexpected HTTP method")
		fmt.Fprint(w, string(b))
	})

	message, _, err := client.PipelineTemplates.Delete(context.Background(), "template2")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "The template 'template2' was deleted successfully.", message)

}

func testListPipelineTemplates(t *testing.T) {

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

func testGetPipelineTemplate(t *testing.T) {

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
