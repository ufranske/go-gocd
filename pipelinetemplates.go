package gocd

import (
	"context"
	//"fmt"
	"fmt"
)

type PipelineTemplateResponse struct {
	Name     string `json:"name"`
	Embedded struct {
		Pipelines []struct {
			Name string `json:"name"`
		}
	} `json:"_embedded,omitempty"`
}

type PipelineTemplatesResponse struct {
	Embedded struct {
		Templates []PipelineTemplateResponse `json:"templates"`
	} `json:"_embedded"`
}

type PipelineTemplate struct {
	Name      string   `json:"name"`
	Pipelines []string `json:",omitempty"`
	Stages    []Stage  `json:"stages"`
}

type PipelineTemplatesService service

func (s *PipelineTemplatesService) GetPipelineTemplate(ctx context.Context, name string) (*PipelineTemplate, *APIResponse, error) {
	u, err := addOptions(fmt.Sprintf("admin/templates/%s", name))

	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil, apiV1)
	if err != nil {
		return nil, nil, err
	}

	var pt PipelineTemplate
	resp, err := s.client.Do(ctx, req, &pt)
	if err != nil {
		return nil, resp, err
	}

	return &pt, resp, nil

}

func (s *PipelineTemplatesService) ListPipelineTemplates(ctx context.Context) (*[]PipelineTemplate, *APIResponse, error) {
	u, err := addOptions("admin/templates")

	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil, apiV4)
	if err != nil {
		return nil, nil, err
	}

	var ptr PipelineTemplatesResponse
	resp, err := s.client.Do(ctx, req, &ptr)
	if err != nil {
		return nil, resp, err
	}

	t := []PipelineTemplate{}

	for _, template := range ptr.Embedded.Templates {
		pt := PipelineTemplate{}
		pt.Name = template.Name

		for _, pipeline := range template.Embedded.Pipelines {
			pt.Pipelines = append(pt.Pipelines, pipeline.Name)
		}
		t = append(t, pt)
	}

	return &t, resp, nil
}
