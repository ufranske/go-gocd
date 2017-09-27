package gocd

import (
	"context"
	"fmt"
)

// AgentsService describes Actions which can be performed on agents
type PropertiesService service

// PropertyRequest describes the parameters to be submitted when calling/creating properties.
type PropertyRequest struct {
	Pipeline        string
	PipelineCounter int
	Stage           string
	StageCounter    int
	Job             string
	LimitPipeline   string
	Limit           int
}

func (ps *PropertiesService) List(ctx context.Context, pr *PropertyRequest) (*Properties, *APIResponse, error) {
	path := fmt.Sprintf("/properties/%s/%d/%s/%d/%s",
		pr.Pipeline, pr.PipelineCounter,
		pr.Stage, pr.StageCounter,
		pr.Job,
	)
	return ps.commonPropertiesAction(ctx, path)
}

func (ps *PropertiesService) Get(ctx context.Context, propertyName string, pr *PropertyRequest) (*Properties, *APIResponse, error) {
	path := fmt.Sprintf("/properties/%s/%d/%s/%d/%s/%s",
		pr.Pipeline, pr.PipelineCounter,
		pr.Stage, pr.StageCounter,
		pr.Job, propertyName,
	)
	return ps.commonPropertiesAction(ctx, path)
}

func (ps *PropertiesService) ListHistorical(ctx context.Context, pr *PropertyRequest) (*Properties, *APIResponse, error) {
	return nil, nil, nil
}

func (ps *PropertiesService) commonPropertiesAction(ctx context.Context, path string) (*Properties, *APIResponse, error) {
	p := Properties{
		UnmarshallWithHeader: true,
	}
	_, resp, err := ps.client.getAction(ctx, &APIClientRequest{
		Path:         path,
		ResponseBody: &p,
	})

	return &p, resp, err
}
