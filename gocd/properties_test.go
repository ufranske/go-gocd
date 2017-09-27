package gocd

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"fmt"
	"net/http"
	"context"
)

func TestProperties(t *testing.T) {
	setup()
	defer teardown()

	t.Run("List", testPropertiesList)
	t.Run("Get", testPropertiesGet)
	t.Run("ListHistorical", testPropertiesListHistorical)
	//t.Run("RemoveLinks", tesPipelineTemplateRemoveLinks)
	//t.Run("Pipelines", testPipelineTemplatePipelines)
	//t.Run("Update", testPipelineTemplateUpdate)
	//t.Run("StageContainerI", testPipelineTemplateStageContainer)
}

func testPropertiesList(t *testing.T) {
	mux.HandleFunc("/go/properties/test-pipeline/5/test-stage/3/test-job", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Unexpected HTTP method")

		j, _ := ioutil.ReadFile("test/resources/properties.0.csv")
		fmt.Fprint(w, string(j))
	})

	p, _, err := client.Properties.List(context.Background(), PropertyRequest{
		Pipeline:        "test-pipeline",
		PipelineCounter: 5,
		Stage:           "test-stage",
		StageCounter:    3,
		Job:             "test-job",
	})

	assert.Nil(t, err)
	assert.Equal(t, []string{"cruise_agent", "cruise_timestamp_01_scheduled", "cruise_timestamp_02_assigned"}, p.Header)
	assert.Equal(t, []string{"myLocalAgent", "2015-07-09T11:59:08+05:30", "2015-07-09T11:59:16+05:30"}, p.DataFrame[0])
}

func testPropertiesGet(t *testing.T) {
	mux.HandleFunc("/go/properties/test-pipeline/5/test-stage/3/test-job/cruise_agent", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Unexpected HTTP method")
		fmt.Fprint(w, `cruise_agent
myLocalAgent`)
	})

	p, _, err := client.Properties.Get(context.Background(), "cruise_agent", PropertyRequest{
		Pipeline:        "test-pipeline",
		PipelineCounter: 5,
		Stage:           "test-stage",
		StageCounter:    3,
		Job:             "test-job",
	})

	assert.Nil(t, err)
	assert.Equal(t, []string{"cruise_agent"}, p.Header)
	assert.Equal(t, []string{"myLocalAgent"}, p.DataFrame[0])
}

func testPropertiesListHistorical(t *testing.T) {
	mux.HandleFunc("/go/properties/search", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Unexpected HTTP method")
		j, _ := ioutil.ReadFile("test/resources/properties.1.csv")
		fmt.Fprint(w, string(j))
	})
	p, _, err := client.Properties.ListHistorical(context.Background(), PropertyRequest{
		Pipeline:      "PipelineName",
		Stage:         "StageName",
		Job:           "JobName",
		LimitPipeline: "latest",
		Limit:         2,
	})

}
