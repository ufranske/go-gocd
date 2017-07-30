package gocd

import (
	"context"
)

type PipelineConfigsService service

type PipelineConfig struct {
}

type PipelineConfigRequest struct {
	Group    string `json:"group"`
	Pipeline *Pipeline `json:"pipeline"`
}

func (pcs *PipelineConfigsService) Create(ctx context.Context, group string, p *Pipeline) (*PipelineConfig, *APIResponse, error) {
	u, err := addOptions("admin/pipelines")
	if err != nil {
		return nil, nil, err
	}

	pt := PipelineConfigRequest{
		Group:    group,
		Pipeline: p,
	}

	req, err := pcs.client.NewRequest("POST", u, pt, apiV4)
	if err != nil {
		return nil, nil, err
	}

	pc := PipelineConfig{}
	resp, err := pcs.client.Do(ctx, req, &pc)
	if err != nil {
		return nil, resp, err
	}

	return &pc, resp, nil
}
