package gocd

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

type PipelineTemplatesService service

//go:generate gocd-response-links-generator -type=PipelineTemplatesLinks,PipelineTemplateLinks
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

type PipelineTemplateRequest struct {
	Name   string   `json:"name"`
	Stages []*Stage `json:"stages"`
}

type PipelineTemplateResponse struct {
	Name     string `json:"name"`
	Embedded *struct {
		Pipelines []*struct {
			Name string `json:"name"`
		}
	} `json:"_embedded,omitempty"`
}

type PipelineTemplatesResponse struct {
	Links    PipelineTemplatesLinks `json:"_links,omitempty"`
	Embedded *struct {
		Templates []*PipelineTemplate `json:"templates"`
	} `json:"_embedded,omitempty"`
}

type PipelineTemplate struct {
	Links    *PipelineTemplateLinks `json:"_links,omitempty"`
	Name     string                 `json:"name"`
	Embedded *struct {
		Pipelines []*Pipeline `json:"pipelines,omitempty"`
	} `json:"_embedded,omitempty"`
	Version string   `json:"template_version"`
	Stages  []*Stage `json:"stages,omitempty"`
}

func (pt *PipelineTemplate) RemoveLinks() {
	pt.Links = nil
}

func (pt *PipelineTemplate) Pipelines() []*Pipeline {
	return pt.Embedded.Pipelines
}

func (s *PipelineTemplatesService) Get(ctx context.Context, name string) (*PipelineTemplate, *APIResponse, error) {
	u, err := addOptions(fmt.Sprintf("admin/templates/%s", name))

	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil, apiV3)
	if err != nil {
		return nil, nil, err
	}

	pt := PipelineTemplate{}
	resp, err := s.client.Do(ctx, req, &pt)
	if err != nil {
		return nil, resp, err
	}

	pt.Version = strings.Replace(resp.Http.Header.Get("Etag"), "\"", "", -1)

	return &pt, resp, nil

}

func (s *PipelineTemplatesService) List(ctx context.Context) ([]*PipelineTemplate, *APIResponse, error) {
	u, err := addOptions("admin/templates")

	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil, apiV3)
	if err != nil {
		return nil, nil, err
	}

	ptr := PipelineTemplatesResponse{}
	resp, err := s.client.Do(ctx, req, &ptr)
	if err != nil {
		return nil, resp, err
	}

	return ptr.Embedded.Templates, resp, nil
}

func (s *PipelineTemplatesService) Create(ctx context.Context, name string, st []*Stage) (*PipelineTemplate, *APIResponse, error) {
	u, err := addOptions("admin/templates")
	if err != nil {
		return nil, nil, err
	}

	pt := PipelineTemplateRequest{
		Name:   name,
		Stages: st,
	}

	req, err := s.client.NewRequest("POST", u, pt, apiV3)
	if err != nil {
		return nil, nil, err
	}

	ptr := PipelineTemplate{}
	resp, err := s.client.Do(ctx, req, &ptr)
	if err != nil {
		return nil, resp, err
	}

	ptr.Version = strings.Replace(resp.Http.Header.Get("Etag"), "\"", "", -1)

	return &ptr, resp, nil

}

func (s *PipelineTemplatesService) Update(ctx context.Context, name string, version string, st []*Stage) (*PipelineTemplate, *APIResponse, error) {
	u, err := addOptions(fmt.Sprintf("admin/templates/%s", name))
	if err != nil {
		return nil, nil, err
	}

	pt := &PipelineTemplateRequest{
		Name:   name,
		Stages: st,
	}

	req, err := s.client.NewRequest("PUT", u, pt, apiV3)
	if err != nil {
		return nil, nil, err
	}

	req.Http.Header.Set("If-Match", fmt.Sprintf("\"%s\"", version))

	ptr := PipelineTemplate{}
	resp, err := s.client.Do(ctx, req, &ptr)
	if err != nil {
		return nil, resp, err
	}

	ptr.Version = strings.Replace(resp.Http.Header.Get("Etag"), "\"", "", -1)

	return &ptr, resp, nil

}

func (s *PipelineTemplatesService) Delete(ctx context.Context, uuid string) (string, *APIResponse, error) {
	u, err := addOptions(fmt.Sprintf("admin/templates/%s", uuid))
	if err != nil {
		return "", nil, err
	}

	req, err := s.client.NewRequest("DELETE", u, nil, apiV4)
	if err != nil {
		return "", nil, err
	}

	a := StringResponse{}
	resp, err := s.client.Do(ctx, req, &a)
	if err != nil {
		return "", resp, err
	}

	return a.Message, resp, nil
}
