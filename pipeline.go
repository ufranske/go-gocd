package gocd

import (
	"context"
	"fmt"
)

type PipelinesService service

type Pipeline struct {
	Name string `json:"name"`
}

type PipelineHistory struct {
	Pipelines []*PipelineInstance `json:"pipelines"`
}
type PipelineInstance struct {
}

func (pgs *PipelinesService) GetHistory(ctx context.Context, name string, offset int) (*PipelineHistory, *APIResponse, error) {
	stub := fmt.Sprintf("pipelines/%s/history", name)
	if offset > 0 {
		stub = fmt.Sprintf("%s/%d", stub, offset)
	}
	u, err := addOptions(stub)

	if err != nil {
		return nil, nil, err
	}

	req, err := pgs.client.NewRequest("GET", u, nil, apiV3)
	if err != nil {
		return nil, nil, err
	}

	pt := PipelineHistory{}
	resp, err := pgs.client.Do(ctx, req, &pt)
	if err != nil {
		return nil, resp, err
	}

	return &pt, resp, nil
}
