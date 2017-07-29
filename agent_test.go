package gocd

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestAgnet_Get(t *testing.T) {

	setup()
	defer teardown()

	mux.HandleFunc("/api/agents/adb9540a-b954-4571-9d9b-2f330739d4da", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		j, _ := ioutil.ReadFile("test/resources/agent.0.json")
		fmt.Fprint(w, string(j))
	})

	agent, _, err := client.Agents.Get(context.Background(), "adb9540a-b954-4571-9d9b-2f330739d4da")
	if err != nil {
		t.Error(err)
	}

	testAgent(t, agent)

	testGotStringSlice(t, []TestStringSlice{
		{agent.BuildDetails.Links.Job.String(), "https://ci.example.com/go/tab/build/detail/up42/1/up42_stage/1/up42_job"},
		{agent.BuildDetails.Links.Stage.String(), "https://ci.example.com/go/pipelines/up42/1/up42_stage/1"},
		{agent.BuildDetails.Links.Pipeline.String(), "https://ci.example.com/go/tab/pipeline/history/up42"},
	})

	if agent.BuildDetails == nil {
		t.Error("Expected 'build_agents'. Got 'nil'.")
	}
}

func TestAgent_List(t *testing.T) {

	setup()
	defer teardown()

	mux.HandleFunc("/api/agents", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testAuth(t, r, mockAuthorization)
		j, _ := ioutil.ReadFile("test/resources/agents.1.json")
		fmt.Fprint(w, string(j))
	})

	agents, _, err := client.Agents.List(context.Background())
	if err != nil {
		t.Error(err)
	}

	if len(agents) != 1 {
		t.Errorf("Expected '1' agents. Got '%d'", len(agents))
	}

	testAgent(t, agents[0])
}

func testAgent(t *testing.T, agent *Agent) {
	testGotStringSlice(t, []TestStringSlice{
		{agent.Links.Self.String(), "https://ci.example.com/go/api/agents/adb9540a-b954-4571-9d9b-2f330739d4da"},
		{agent.Links.Doc.String(), "https://api.gocd.org/#agents"},
		{agent.Links.Find.String(), "https://ci.example.com/go/api/agents/:uuid"},
		{agent.Uuid, "adb9540a-b954-4571-9d9b-2f330739d4da"},
		{agent.Hostname, "agent01.example.com"},
		{agent.IpAddress, "10.12.20.47"},
		{agent.Sandbox, "/Users/ketanpadegaonkar/projects/gocd/gocd/agent"},
		{agent.OperatingSystem, "Mac OS X"},
		{agent.AgentConfigState, "Enabled"},
		{agent.AgentState, "Idle"},
		{agent.Resources[0], "java"},
		{agent.Resources[1], "linux"},
		{agent.Resources[2], "firefox"},
		{agent.Environments[0], "perf"},
		{agent.Environments[1], "UAT"},
		{agent.BuildState, "Idle"},
	})

	if agent.FreeSpace != 84983328768 {
		t.Errorf("Expected '84983328768'. Got '%d'.", agent.FreeSpace)
	}

}
