package gocd

import (
	"context"
	"fmt"
	"strings"
)

type PipelineConfigsService service

type PipelineConfig struct {
}

type PipelineConfigRequest struct {
	Group    string    `json:"group"`
	Pipeline *Pipeline `json:"pipeline"`
}

func (pcs *PipelineConfigsService) Update(ctx context.Context, group string, name string, version string, p *Pipeline) (*Pipeline, *APIResponse, error) {
	u, err := addOptions(fmt.Sprintf("admin/pipelines/%s", name))
	if err != nil {
		return nil, nil, err
	}

	pt := &PipelineConfigRequest{
		Group:    group,
		Pipeline: p,
	}

	req, err := pcs.client.NewRequest("PUT", u, pt, apiV4)
	if err != nil {
		return nil, nil, err
	}

	req.Http.Header.Set("If-Match", fmt.Sprintf("\"%s\"", version))

	pc := Pipeline{}
	resp, err := pcs.client.Do(ctx, req, &pc)
	if err != nil {
		return nil, resp, err
	}

	return &pc, resp, nil

}

func (pcs *PipelineConfigsService) Create(ctx context.Context, group string, p *Pipeline) (*Pipeline, *APIResponse, error) {
	u, err := addOptions("admin/pipelines")
	if err != nil {
		return nil, nil, err
	}

	pt := &PipelineConfigRequest{
		Group:    group,
		Pipeline: p,
	}

	req, err := pcs.client.NewRequest("POST", u, pt, apiV4)
	if err != nil {
		return nil, nil, err
	}

	pc := Pipeline{}
	resp, err := pcs.client.Do(ctx, req, &pc)
	if err != nil {
		return nil, resp, err
	}
	pc.Version = strings.Replace(resp.Http.Header.Get("Etag"), "\"", "", -1)

	return &pc, resp, nil
}

func (pcs *PipelineConfigsService) Delete(ctx context.Context, name string) (string, *APIResponse, error) {
	return pcs.client.genericDeleteAction(ctx, fmt.Sprintf("admin/pipelines/%s", name))
}
