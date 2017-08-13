package gocd

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPipelineService(t *testing.T) {
	setup()
	defer teardown()

	t.Run("Get", testPipelineServiceGet)
	t.Run("GetHistory", testPipelineServiceGetHistory)
	t.Run("GetStatus", testPipelineServiceGetStatus)
}

func testPipelineServiceGetStatus(t *testing.T) {
	mux.HandleFunc("/api/pipelines/test-pipeline/status", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Unexpected HTTP method")
		j, _ := ioutil.ReadFile("test/resources/pipeline.2.json")
		fmt.Fprint(w, string(j))
	})

	ps, _, err := client.Pipelines.GetStatus(context.Background(), "test-pipeline", 0)
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, ps)
	assert.False(t, ps.Locked)
	assert.True(t, ps.Paused)
	assert.False(t, ps.Schedulable)
}

func testPipelineServiceGet(t *testing.T) {
	mux.HandleFunc("/api/pipelines/test-pipeline/instance", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Unexpected HTTP method")
		j, _ := ioutil.ReadFile("test/resources/pipeline.0.json")
		fmt.Fprint(w, string(j))
	})

	p, _, err := client.Pipelines.Get(context.Background(), "test-pipeline", 0)
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, p)
	assert.Equal(t, p.Name, "test-pipeline")

	assert.Len(t, p.Stages, 1)

	s := p.Stages[0]
	assert.Equal(t, "stage1", s.Name)
	assert.Equal(t, false, s.FetchMaterials)
	assert.Equal(t, false, s.CleanWorkingDirectory)
	assert.Equal(t, false, s.NeverCleanupArtifacts)

	assert.Len(t, s.EnvironmentVariables, 0)
	assert.Nil(t, s.Approval)
}

func testPipelineServiceGetHistory(t *testing.T) {
	mux.HandleFunc("/api/pipelines/test-pipeline/history", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Unexpected HTTP method")
		j, _ := ioutil.ReadFile("test/resources/pipeline.1.json")
		fmt.Fprint(w, string(j))
	})
	ph, _, err := client.Pipelines.GetHistory(context.Background(), "test-pipeline", 0)
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, ph)
	assert.Len(t, ph.Pipelines, 2)

	h1 := ph.Pipelines[0]
	assert.True(t, h1.CanRun)
	assert.Equal(t, h1.Name, "pipeline1")
	assert.Equal(t, h1.NaturalOrder, 11)
	assert.Equal(t, h1.Comment, "")
	assert.Len(t, h1.Stages, 1)

	h1s := h1.Stages[0]
	assert.Equal(t, h1s.Name, "stage1")

	h2 := ph.Pipelines[1]
	assert.True(t, h2.CanRun)
	assert.Equal(t, h2.Name, "pipeline1")
	assert.Equal(t, h2.NaturalOrder, 10)
	assert.Equal(t, h2.Comment, "")
	assert.Len(t, h2.Stages, 1)

	h2s := h2.Stages[0]
	assert.Equal(t, h2s.Name, "stage1")

}
