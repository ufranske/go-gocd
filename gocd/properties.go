package gocd

import (
	"context"
	"fmt"
	"bytes"
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

// PropertyCreateResponse handles the parsing of the response when creating a property
type PropertyCreateResponse struct {
	Name  string
	Value string
}

func (ps *PropertiesService) List(ctx context.Context, pr *PropertyRequest) (*Properties, *APIResponse, error) {
	path := fmt.Sprintf("/properties/%s/%d/%s/%d/%s",
		pr.Pipeline, pr.PipelineCounter,
		pr.Stage, pr.StageCounter,
		pr.Job,
	)
	return ps.commonPropertiesAction(ctx, path)
}

func (ps *PropertiesService) Get(ctx context.Context, name string, pr *PropertyRequest) (*Properties, *APIResponse, error) {
	path := fmt.Sprintf("/properties/%s/%d/%s/%d/%s/%s",
		pr.Pipeline, pr.PipelineCounter,
		pr.Stage, pr.StageCounter,
		pr.Job, name,
	)
	return ps.commonPropertiesAction(ctx, path)
}

func (ps *PropertiesService) Create(ctx context.Context, name string, value string, pr *PropertyRequest) (bool, *APIResponse, error) {
	path := fmt.Sprintf("/properties/%s/%d/%s/%d/%s/%s",
		pr.Pipeline, pr.PipelineCounter,
		pr.Stage, pr.StageCounter,
		pr.Job, name,
	)


	responseBuffer := bytes.NewBuffer([]byte(""))
	_, resp, err := ps.client.postAction(ctx, &APIClientRequest{
		Path:         path,
		ResponseType: responseTypeText,
		ResponseBody: responseBuffer,
	})
	responseString := responseBuffer.String()

	r := fmt.Sprintf("Property '%s' created with value '%s'", name, value)

	return responseString == r, resp, err
}

func (ps *PropertiesService) ListHistorical(ctx context.Context, pr *PropertyRequest) (*Properties, *APIResponse, error) {
	u := ps.client.BaseURL
	q := u.Query()
	q.Set("pipelineName", pr.Pipeline)
	q.Set("stageName", pr.Stage)
	q.Set("jobName", pr.Job)
	if pr.Limit >= 0 && pr.LimitPipeline != "" {
		q.Set("limitCount", fmt.Sprintf("%d", pr.Limit))
		q.Set("limitPipeline", pr.LimitPipeline)
	}
	u.RawQuery = q.Encode()
	return ps.commonPropertiesAction(ctx, "/properties/search")
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
