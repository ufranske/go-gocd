package gocd

import (
	"context"
	//"fmt"
	"fmt"
	"net/url"
)

type PipelineTemplatesService service

//go:generate gocd-response-links -type=PipelineTemplatesLinks,PipelineTemplateLinks
type PipelineTemplatesLinks struct {
	Self *url.URL `json:"self"`
	Doc  *url.URL `json:"doc"`
	Find *url.URL `json:"find"`
}

type PipelineTemplateLinks struct {
	Self *url.URL `json:"self"`
	Doc  *url.URL `json:"doc"`
	Find *url.URL `json:"find"`
}
type PipelineTemplateResponse struct {
	Name     string `json:"name"`
	Embedded struct {
		Pipelines []struct {
			Name string `json:"name"`
		}
	} `json:"_embedded,omitempty"`
}

type PipelineTemplatesResponse struct {
	Links    PipelineTemplatesLinks `json:"_links,omitempty"`
	Embedded struct {
		Templates []*PipelineTemplate `json:"agents"`
	} `json:"_embedded"`
}

type PipelineTemplate struct {
	Name      string   `json:"name"`
	Pipelines []string `json:",omitempty"`
	Stages    []Stage  `json:"stages"`
}

func (s *PipelineTemplatesService) GetPipelineTemplate(ctx context.Context, name string) (*PipelineTemplate, *APIResponse, error) {
	u, err := addOptions(fmt.Sprintf("admin/templates/%s", name))

	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil, apiV3)
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

//func (s *PipelineTemplatesService) List(ctx context.Context) ([]*PipelineTemplate, *APIResponse, error) {
//	u, err := addOptions("admin/templates")
//
//	if err != nil {
//		return nil, nil, err
//	}
//
//	req, err := s.client.NewRequest("GET", u, nil, apiV3)
//	if err != nil {
//		return nil, nil, err
//	}
//
//	var ptr PipelineTemplatesResponse
//	resp, err := s.client.Do(ctx, req, &ptr)
//	if err != nil {
//		return nil, resp, err
//	}
//
//	t := []PipelineTemplate{}
//
//	for _, template := range ptr.Embedded.Templates {
//		pt := PipelineTemplate{}
//		pt.Name = template.Name
//
//		for _, pipeline := range template.Embedded.Pipelines {
//			pt.Pipelines = append(pt.Pipelines, pipeline.Name)
//		}
//		t = append(t, pt)
//	}
//
//	return &t, resp, nil
//}
