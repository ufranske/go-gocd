package gocd

import (
	"context"
	"net/url"
)

// EnvironmentsService exposes calls for interacting with Environment objects in the GoCD API.
type EnvironmentsService service

// EnvironmentsResponseLinks describes the HAL _link resource for the api response object for a collection of environment
// objects
//go:generate gocd-response-links-generator -type=EnvironmentsResponseLinks,EnvironmentLinks
type EnvironmentsResponseLinks struct {
	Self *url.URL `json:"self"`
	Doc  *url.URL `json:"doc"`
}

// EnvironmentLinks describes the HAL _link resource for the api response object for a collection of environment objects.
type EnvironmentLinks struct {
	Self *url.URL `json:"self"`
	Doc  *url.URL `json:"doc"`
	Find *url.URL `json:"find"`
}

// EnvironmentsResponse describes the response obejct for a plugin API call.
type EnvironmentsResponse struct {
	Links    *EnvironmentsResponseLinks `json:"_links"`
	Embedded struct {
		Environments []*Environment `json:"environments"`
	} `json:"_embedded"`
}

// Environment describes a group of pipelines and agents
type Environment struct {
	Links                *EnvironmentLinks       `json:"_links,omitempty"`
	Name                 string                 `json:"name"`
	Pipelines            []*Pipeline            `json:"pipelines,omitempty"`
	Agents               []*Agent               `json:"agents,omitempty"`
	EnvironmentVariables []*EnvironmentVariable `json:"environment_variables,omitempty"`
}

// List all environments
func (es *EnvironmentsService) List(ctx context.Context) (*EnvironmentsResponse, *APIResponse, error) {
	e := EnvironmentsResponse{}
	_, resp, err := es.client.getAction(ctx, &APIClientRequest{
		Path:         "admin/environments",
		ResponseBody: &e,
		APIVersion:   apiV2,
	})

	return &e, resp, err
}
