package gocd

import "context"

type PipelineGroupsService service

type PipelineGroup struct {
	Name string `json:"name"`
	Pipelines []struct {
		Name string   `json:"name"`
		Stages []struct {
			Name string `json:"name"`
		}`json:"stages"`
		Materials []struct {
			Description string `json:"description"`
			Fingerprint string `json:"fingerprint"`
			Type string `json:"type"`
		}`json:"materials"`
		Label string `json:"label"`
	} `json:"pipelines"`
}

func (pgs *PipelineGroupsService) List(ctx context.Context) ([]*PipelineGroup, *APIResponse, error) {
	u, err := addOptions("config/pipeline_groups")

	if err != nil {
		return nil, nil, err
	}

	req, err := pgs.client.NewRequest("GET", u, nil, "")
	if err != nil {
		return nil, nil, err
	}

	pg := []*PipelineGroup{}
	resp, err := pgs.client.Do(ctx, req, &pg)
	if err != nil {
		return nil, resp, err
	}

	return pg, resp, nil
}
