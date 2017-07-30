package gocd

import (
	"context"
	"fmt"
)

type PipelinesService service

type Pipeline struct {
	Name                  string     `json:"name"`
	LabelTemplate         string     `json:"label_template,omitempty"`
	EnablePipelineLocking bool       `json:"enable_pipeline_locking,omitempty"`
	Template              string     `json:"template,omitempty"`
	Materials             []Material `json:"materials,omitempty"`
	Stages                []Stage    `json:"stages"`
	Version               string
}

type Material struct {
	Type string `json:"type"`
	Attributes struct {
		Url             string      `json:"url"`
		Destination     string      `json:"destination,omitempty"`
		Filter          interface{} `json:"filter,omitempty"`
		InvertFilter    bool        `json:"invert_filter,omitempty"`
		Name            string      `json:"name,omitempty"`
		AutoUpdate      bool        `json:"auto_update,omitempty"`
		Branch          string      `json:"branch,omitempty"`
		SubmoduleFolder string      `json:"submodule_folder,omitempty"`
		ShallowClone    bool        `json:"shallow_clone,omitempty"`
	} `json:"attributes"`
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
