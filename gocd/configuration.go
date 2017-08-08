package gocd

import (
	"context"
)

// ConfigurationService describes the HAL _link resource for the api response object for a pipelineconfig
type ConfigurationService service

type ConfigXML struct {

}

// Get will retrieve all agents, their status, and metadata from the GoCD Server.
// Get returns a list of pipeline instanves describing the pipeline history.
func (cd *ConfigurationService) Get(ctx context.Context) (*ConfigXML, *APIResponse, error) {
	req, err := cd.client.NewRequest("GET", "admin/config.xml", nil, "")
	if err != nil {
		return nil, nil, err
	}

	cx := ConfigXML{}
	resp, err := cd.client.Do(ctx, req, &cx, responseTypeXML)
	if err != nil {
		return nil, resp, err
	}

	return &cx, resp, nil
}
