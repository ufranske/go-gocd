package gocd

import (
	"context"
	"fmt"
	"net/url"
)

type AgentsService service

type AgentResponse struct {
	Links ResponseLinks `json:"_links,omitempty"`
	Embedded struct {
		Agents []*Agent `json:"agents"`
	}`json:"_embedded"`
}

type Agent struct {
	Uuid             string `json:"uuid"`
	Hostname         string `json:"hostname"`
	ElasticAgentId   string `json:"elastic_agent_id"`
	ElasticPluginId  string `json:"elastic_plugin_id"`
	IpAddress        string `json:"ip_address"`
	Sandbox          string `json:"sandbox"`
	OperatingSystem  string `json:"operating_system"`
	FreeSpace        int64 `json:"free_space"`
	AgentConfigState string `json:"agent_config_state"`
	AgentState       string `json:"agent_state"`
	Resources        []string `json:"resources"`
	Environments     []string `json:"environments"`
	BuildState       string `json:"build_state"`
	BuildDetails     *BuildDetails `json:"build_details"`
	Links            *ResponseLinks `json:"_links"`
}

type BuildDetails struct {
	Links    *BuildDetails_ResponseLinks `json:"_links"`
	Pipeline string `json:"pipeline"`
	Stage    string `json:"stage"`
	Job      string `json:"job"`
}

//go:generate gocd-response-links -type=BuildDetails_ResponseLinks -output=responselinks_builddetails.go
type BuildDetails_ResponseLinks struct {
	Doc    *url.URL
	Find   *url.URL
	Job    *url.URL
	Latest *url.URL
	Next   *url.URL
	Oldest *url.URL
	Self   *url.URL
}


func (s *AgentsService) Get(ctx context.Context, uuid string) (*Agent, *APIResponse, error) {
	u, err := addOptions(fmt.Sprintf("agents/%s", uuid))
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil, apiV4)
	if err != nil {
		return nil, nil, err
	}

	a := Agent{}
	resp, err := s.client.Do(ctx, req, &a)
	if err != nil {
		return nil, resp, err
	}

	return &a, resp, nil
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

	r := AgentResponse{}
	resp, err := s.client.Do(ctx, req, &r)
	if err != nil {
		return nil, resp, err
	}

	//return &r.Embedded, resp, nil
	return r.Embedded.Agents, resp, nil
}
