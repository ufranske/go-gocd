package gocd

import (
	"context"
	"fmt"
	"net/url"
)

type AgentsService service

//go:generate gocd-response-links -type=AgentsLinks,AgentLinks
type AgentsLinks struct {
	Self *url.URL `json:"self"`
	Doc  *url.URL `json:"doc"`
}

type AgentLinks struct {
	Self *url.URL `json:"self"`
	Doc  *url.URL `json:"doc"`
	Find *url.URL `json:"find"`
}

type AgentsResponse struct {
	Links    *AgentsLinks `json:"_links,omitempty"`
	Embedded *struct {
		Agents []*Agent `json:"agents"`
	} `json:"_embedded,omitempty"`
}

type Agent struct {
	Uuid             string        `json:"uuid"`
	Hostname         string        `json:"hostname"`
	ElasticAgentId   string        `json:"elastic_agent_id"`
	ElasticPluginId  string        `json:"elastic_plugin_id"`
	IpAddress        string        `json:"ip_address"`
	Sandbox          string        `json:"sandbox"`
	OperatingSystem  string        `json:"operating_system"`
	FreeSpace        int64         `json:"free_space"`
	AgentConfigState string        `json:"agent_config_state"`
	AgentState       string        `json:"agent_state"`
	Resources        []string      `json:"resources"`
	Environments     []string      `json:"environments"`
	BuildState       string        `json:"build_state"`
	BuildDetails     *BuildDetails `json:"build_details"`
	Links            *AgentLinks   `json:"_links,omitempty"`
	client           *Client
}

func (a *Agent) JobRunHistory(ctx context.Context) ([]*Job, error) {
	jobs, _, err := a.client.Agents.JobRunHistory(ctx, a.Uuid)
	if err != nil {
		return nil, err
	}

	return jobs, nil
}

type AgentUpdate struct {
	Hostname         string   `json:"hostname,omitempty"`
	Resources        []string `json:"resources,omitempty"`
	Environments     []string `json:"environments,omitempty"`
	AgentConfigState string   `json:"agent_config_state,omitempty"`
}

type AgentBulkUpdate struct {
	Uuids            []string                   `json:"uuids"`
	Operations       *AgentBulkOperationsUpdate `json:"operations,omitempty"`
	AgentConfigState string                     `json:"agent_config_state,omitempty"`
}

type AgentBulkOperationsUpdate struct {
	Environments *AgentBulkOperationUpdate `json:"environments,omitempty"`
	Resources    *AgentBulkOperationUpdate `json:"resources,omitempty"`
}

type AgentBulkOperationUpdate struct {
	Add    []string `json:"add,omitempty"`
	Remove []string `json:"remove,omitempty"`
}

type BuildDetails struct {
	Links    *BuildDetailsLinks `json:"_links"`
	Pipeline string             `json:"pipeline"`
	Stage    string             `json:"stage"`
	Job      string             `json:"job"`
}

//go:generate gocd-response-links -type=BuildDetailsLinks
type BuildDetailsLinks struct {
	Job      *url.URL `json:"job"`
	Stage    *url.URL `json:"stage"`
	Pipeline *url.URL `json:"pipeline"`
}

func (a *Agent) RemoveLinks() {
	a.Links = nil
}

func (s *AgentsService) List(ctx context.Context) ([]*Agent, *APIResponse, error) {
	u, err := addOptions("agents")

	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil, apiV4)
	if err != nil {
		return nil, nil, err
	}

	r := AgentsResponse{}
	resp, err := s.client.Do(ctx, req, &r)
	if err != nil {
		return nil, resp, err
	}

	for _, agent := range r.Embedded.Agents {
		agent.client = s.client
	}

	return r.Embedded.Agents, resp, nil
}

func (s *AgentsService) Get(ctx context.Context, uuid string) (*Agent, *APIResponse, error) {
	return s.handleAgentRequest(ctx, "GET", uuid, nil)
}

func (s *AgentsService) Update(ctx context.Context, uuid string, agent AgentUpdate) (*Agent, *APIResponse, error) {
	return s.handleAgentRequest(ctx, "PATCH", uuid, &agent)
}

func (s *AgentsService) Delete(ctx context.Context, uuid string) (string, *APIResponse, error) {
	u, err := addOptions(fmt.Sprintf("agents/%s", uuid))
	if err != nil {
		return "", nil, err
	}

	req, err := s.client.NewRequest("DELETE", u, nil, apiV4)
	if err != nil {
		return "", nil, err
	}

	a := StringResponse{}
	resp, err := s.client.Do(ctx, req, &a)
	if err != nil {
		return "", resp, err
	}

	return a.Message, resp, nil
}

func (s *AgentsService) BulkUpdate(ctx context.Context, agents AgentBulkUpdate) (string, *APIResponse, error) {
	u, err := addOptions("agents")
	if err != nil {
		return "", nil, err
	}

	req, err := s.client.NewRequest("PATCH", u, agents, apiV4)
	if err != nil {
		return "", nil, err
	}
	a := StringResponse{}
	resp, err := s.client.Do(ctx, req, &a)
	if err != nil {
		return "", resp, err
	}

	return a.Message, resp, nil
}

func (s *AgentsService) JobRunHistory(ctx context.Context, uuid string) ([]*Job, *APIResponse, error) {
	u, err := addOptions(fmt.Sprintf("agents/%s/job_run_history", uuid))
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil, apiV4)
	if err != nil {
		return nil, nil, err
	}
	a := JobRunHistoryResponse{}
	resp, err := s.client.Do(ctx, req, &a)
	if err != nil {
		return nil, resp, err
	}

	return a.Jobs, resp, nil

}

func (s *AgentsService) handleAgentRequest(ctx context.Context, action string, uuid string, body *AgentUpdate) (*Agent, *APIResponse, error) {
	u, err := addOptions(fmt.Sprintf("agents/%s", uuid))
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(action, u, body, apiV4)
	if err != nil {
		return nil, nil, err
	}

	a := Agent{client: s.client}
	resp, err := s.client.Do(ctx, req, &a)
	if err != nil {
		return nil, resp, err
	}

	return &a, resp, nil
}
