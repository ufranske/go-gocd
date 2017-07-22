package gocd

import (
	"context"
)

type PipelineTemplatesService service

type PipelineTemplateResponse struct {
	Embedded struct {
				 Templates []struct {
					 Name     string `json:"name"`
					 Embedded struct {
								  Pipelines []struct {
									  Name string `json:"name"`
								  }
							  }`json:"_embedded"`
				 } `json:"templates"`
			 } `json:"_embedded"`
}

type PipelineTemplate struct {
	Name      string
	Pipelines []string
}

func (s *PipelineTemplatesService) ListPipelines(ctx context.Context) (*[]PipelineTemplate, *Response, error) {
	u, err := addOptions("events", struct{}{})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	//var pipelineTemplates list.List[]
	var ptr PipelineTemplateResponse
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
