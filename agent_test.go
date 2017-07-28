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

	if agent.BuildDetails != nil {
		t.Error("Expected 'build_agents'. Got 'nil'.")
	}

	if agent.BuildDetails.Links.Job.String() != "https://ci.example.com/go/api/agents/adb9540a-b954-4571-9d9b-2f330739d4da" {
		t.Errorf(
			"Expected '%s'. Got '%s'",
			"https://ci.example.com/go/api/agents/adb9540a-b954-4571-9d9b-2f330739d4da",
			agent.Links.Self.String(),
		)
	}

	if agent.BuildDetails.Links.Doc.String() != "https://api.gocd.org/#agents" {
		t.Errorf(
			"Expected '%s'. Got '%s'",
			"https://api.gocd.org/#agents",
			agent.Links.Self.String(),
		)
	}

	if agent.BuildDetails.Links.Find.String() != "https://ci.example.com/go/api/agents/:uuid" {
		t.Errorf(
			"Expected '%s'. Got '%s'",
			"https://ci.example.com/go/api/agents/:uuid",
			agent.Links.Self.String(),
		)
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
		t.Errorf("Expected '1' agents. Got '%s'", len(agents))
	}

	testAgent(t, agents[0])

}

func testAgent(t *testing.T, agent *Agent) {
	if agent.Links.Self.String() != "https://ci.example.com/go/api/agents/adb9540a-b954-4571-9d9b-2f330739d4da" {
		t.Errorf(
			"Expected '%s'. Got '%s'",
			"https://ci.example.com/go/api/agents/adb9540a-b954-4571-9d9b-2f330739d4da",
			agent.Links.Self.String(),
		)
	}

	if agent.Links.Doc.String() != "https://api.gocd.org/#agents" {
		t.Errorf(
			"Expected '%s'. Got '%s'",
			"https://api.gocd.org/#agents",
			agent.Links.Self.String(),
		)
	}

	if agent.Links.Find.String() != "https://ci.example.com/go/api/agents/:uuid" {
		t.Errorf(
			"Expected '%s'. Got '%s'",
			"https://ci.example.com/go/api/agents/:uuid",
			agent.Links.Self.String(),
		)
	}

	if agent.Uuid != "adb9540a-b954-4571-9d9b-2f330739d4da" {
		t.Error(
			"Expected '%s'. Got '%s'.",
			"adb9540a-b954-4571-9d9b-2f330739d4da",
			agent.Uuid,
		)
	}

	if agent.Hostname != "agent01.example.com" {
		t.Errorf(
			"Expected 'agent01.example.com'. Got '%s'.",
			agent.Hostname,
		)
	}

	if agent.IpAddress != "10.12.20.47" {
		t.Errorf(
			"Expected '10.12.20.47'. Got '%s'.",
			agent.IpAddress,
		)
	}

	if agent.Sandbox != "/Users/ketanpadegaonkar/projects/gocd/gocd/agent" {
		t.Errorf(
			"Expected '%s'. Got '%s'.",
			"/Users/ketanpadegaonkar/projects/gocd/gocd/agent",
			agent.Sandbox,
		)
	}

	if agent.OperatingSystem != "Mac OS X" {
		t.Errorf(
			"Expected 'Mac OS X'. Got '%s'.",
			agent.OperatingSystem,
		)
	}

	if agent.FreeSpace != 84983328768 {
		t.Errorf(
			"Expected '%d'. Got '%d'.",
			84983328768,
			agent.FreeSpace,
		)
	}

	if agent.AgentConfigState != "Enabled" {
		t.Errorf(
			"Expected 'Enabled'. Got '%s'.",
			agent.AgentConfigState,
		)
	}

	if agent.AgentState != "Idle" {
		t.Errorf(
			"Expected 'Idle'. Got '%s'.",
			agent.AgentState,
		)
	}

	if agent.Resources[0] != "java" {
		t.Errorf(
			"Expected 'java'. Got '%s'.",
			agent.Resources[0],
		)
	}

	if agent.Resources[1] != "linux" {
		t.Errorf(
			"Expected 'linux'. Got '%s'.",
			agent.Resources[1],
		)
	}

	if agent.Resources[2] != "firefox" {
		t.Errorf(
			"Expected 'firefox'. Got '%s'.",
			agent.Resources[2],
		)
	}

	if agent.Environments[0] != "perf" {
		t.Errorf(
			"Expected 'perf'. Got '%s'.",
			agent.Environments[0],
		)
	}

	if agent.Environments[1] != "UAT" {
		t.Errorf(
			"Expected 'UAT'. Got '%s'.",
			agent.Environments[1],
		)
	}

	if agent.BuildState != "Idle" {
		t.Errorf(
			"Expected 'Idle'. Got '%s'.",
			agent.BuildState,
		)
	}
}
