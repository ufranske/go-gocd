package gocd

import (
	"context"
)

// ConfigurationService describes the HAL _link resource for the api response object for a pipelineconfig
type ConfigurationService service

// Get will retrieve all agents, their status, and metadata from the GoCD Server.
// Get returns a list of pipeline instanves describing the pipeline history.
func (cd *ConfigurationService) Get(ctx context.Context) (*PipelineStatus, *APIResponse, error) {
	req, err := cd.client.NewRequest("GET", "admin/config.xml", nil, "")
	if err != nil {
		return nil, nil, err
	}

	ps := PipelineStatus{}
	resp, err := cd.client.Do(ctx, req, &ps, responseTypeXML)
	if err != nil {
		return nil, resp, err
	}

	return &ps, resp, nil
}
