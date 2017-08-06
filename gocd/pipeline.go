package gocd

import (
	"context"
	"fmt"
)

// PipelinesService describes the HAL _link resource for the api response object for a pipelineconfig
type PipelinesService service

// Pipeline describes a pipeline object
type Pipeline struct {
	Name                  string     `json:"name"`
	LabelTemplate         string     `json:"label_template,omitempty"`
	EnablePipelineLocking bool       `json:"enable_pipeline_locking,omitempty"`
	Template              string     `json:"template,omitempty"`
	Materials             []Material `json:"materials,omitempty"`
	Stages                []Stage    `json:"stages"`
	Version               string
}

// Material describes an artifact dependency for a pipeline object.
type Material struct {
	Type       string `json:"type"`
	Attributes struct {
		URL             string      `json:"url"`
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

// PipelineHistory describes the history of runs for a pipeline
type PipelineHistory struct {
	Pipelines []*PipelineInstance `json:"pipelines"`
}

// PipelineInstance describes a single pipeline run
type PipelineInstance struct {
}

// GetHistory returns a list of pipeline instanves describing the pipeline history.
func (pgs *PipelinesService) GetHistory(ctx context.Context, name string, offset int) (*PipelineHistory, *APIResponse, error) {
	stub := fmt.Sprintf("pipelines/%s/history", name)
	if offset > 0 {
		stub = fmt.Sprintf("%s/%d", stub, offset)
	}

	req, err := pgs.client.NewRequest("GET", stub, nil, apiV3)
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
