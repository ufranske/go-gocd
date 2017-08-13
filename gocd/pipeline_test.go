package gocd

import (
	"testing"
	"context"
	"io/ioutil"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
)

func TestPipelineService(t *testing.T) {
	setup()
	defer teardown()

	t.Run("Get", testPipelineServiceGet)
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
