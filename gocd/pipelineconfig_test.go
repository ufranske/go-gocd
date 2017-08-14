package gocd

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPipelineConfig(t *testing.T) {
	setup()
	defer teardown()
	t.Run("Create", testPipelineConfigCreate)
	t.Run("Update", testPipelineConfigUpdate)
}

func testPipelineConfigCreate(t *testing.T) {
	mux.HandleFunc("/api/admin/pipelines", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST", "Unexpected HTTP method")
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(
			t,
			"{\n  \"group\": \"test-group\",\n  \"pipeline\": {\n    \"name\": \"\",\n    \"stages\": null,\n    \"Version\": \"\"\n  }\n}\n",
			string(b))
		j, _ := ioutil.ReadFile("test/resources/pipelineconfig.0.json")
		fmt.Fprint(w, string(j))
	})

	p := Pipeline{}
	pgs, _, err := client.PipelineConfigs.Create(context.Background(), "test-group", &p)
	if err != nil {
		t.Error(t, err)
	}

	assert.NotNil(t, pgs)
}

func testPipelineConfigUpdate(t *testing.T) {
	mux.HandleFunc("/api/admin/pipelines/test-name", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PUT", r.Method, "Unexpected HTTP method")
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(
			t,
			"{\n  \"group\": \"test-group\",\n  \"pipeline\": {\n    \"name\": \"\",\n    \"stages\": null,\n    \"Version\": \"\"\n  }\n}\n",
			string(b))
		j, _ := ioutil.ReadFile("test/resources/pipelineconfig.0.json")
		fmt.Fprint(w, string(j))
	})

	p := Pipeline{}
	pcs, _, err := client.PipelineConfigs.Update(context.Background(),
		"test-group", "test-name", "test-version", &p)
	if err != nil {
		t.Error(t, err)
	}

	assert.NotNil(t, pcs)
}
