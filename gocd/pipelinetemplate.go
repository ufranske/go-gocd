package gocd

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

// PipelineTemplatesService describes the HAL _link resource for the api response object for a pipeline configuration objects.
type PipelineTemplatesService service

// PipelineTemplatesLinks describes a single pipeline template config object HAL links
//go:generate gocd-response-links-generator -type=PipelineTemplatesLinks,PipelineTemplateLinks
type PipelineTemplatesLinks struct {
	Self *url.URL `json:"self"`
	Doc  *url.URL `json:"doc"`
	Find *url.URL `json:"find"`
}

// PipelineTemplateLinks describes multiple pipeline template config object HAL links
type PipelineTemplateLinks struct {
	Self *url.URL `json:"self"`
	Doc  *url.URL `json:"doc"`
	Find *url.URL `json:"find"`
}

// PipelineTemplateRequest describes a PipelineTemplate
type PipelineTemplateRequest struct {
	Name   string   `json:"name"`
	Stages []*Stage `json:"stages"`
}

// PipelineTemplateResponse describes an api response for a single pipeline templates
type PipelineTemplateResponse struct {
	Name     string `json:"name"`
	Embedded *struct {
		Pipelines []*struct {
			Name string `json:"name"`
		}
	} `json:"_embedded,omitempty"`
}

// PipelineTemplatesResponse describes an api response for multiple pipeline templates
type PipelineTemplatesResponse struct {
	Links    PipelineTemplatesLinks `json:"_links,omitempty"`
	Embedded *struct {
		Templates []*PipelineTemplate `json:"templates"`
	} `json:"_embedded,omitempty"`
}

// PipelineTemplate describes a response from the API for a pipeline template object.
type PipelineTemplate struct {
	Links    *PipelineTemplateLinks `json:"_links,omitempty"`
	Name     string                 `json:"name"`
	Embedded *struct {
		Pipelines []*Pipeline `json:"pipelines,omitempty"`
	} `json:"_embedded,omitempty"`
	Version string   `json:"template_version"`
	Stages  []*Stage `json:"stages,omitempty"`
}

// RemoveLinks gets the PipelineTemplate ready to be submitted to the GoCD API.
func (pt *PipelineTemplate) RemoveLinks() {
	pt.Links = nil
}

// Pipelines returns a list of Pipelines attached to this PipelineTemplate object.
func (pt *PipelineTemplate) Pipelines() []*Pipeline {
	return pt.Embedded.Pipelines
}

// Exists ensures whether or not a PipelineTemplate is present in the API.
func (pts *PipelineTemplatesService) Exists(ctx context.Context, name string) (bool, *APIResponse, error) {
	return pts.client.genericHeadAction(ctx, fmt.Sprintf("admin/templates/%s", name), apiV3)
}

// Get a single PipelineTemplate object in the GoCD API.
func (pts *PipelineTemplatesService) Get(ctx context.Context, name string) (*PipelineTemplate, *APIResponse, error) {

	req, err := pts.client.NewRequest("GET", "admin/templates/"+name, nil, apiV3)
	if err != nil {
		return nil, nil, err
	}

	pt := PipelineTemplate{}
	resp, err := pts.client.Do(ctx, req, &pt, responseTypeJSON)
	if err != nil {
		return nil, resp, err
	}

	pt.Version = strings.Replace(resp.HTTP.Header.Get("Etag"), "\"", "", -1)

	return &pt, resp, nil

}

// List all PipelineTemplate objects in the GoCD API.
func (pts *PipelineTemplatesService) List(ctx context.Context) ([]*PipelineTemplate, *APIResponse, error) {

	req, err := pts.client.NewRequest("GET", "admin/templates", nil, apiV3)
	if err != nil {
		return nil, nil, err
	}

	ptr := PipelineTemplatesResponse{}
	resp, err := pts.client.Do(ctx, req, &ptr, responseTypeJSON)
	if err != nil {
		return nil, resp, err
	}

	return ptr.Embedded.Templates, resp, nil
}

// Create a new PipelineTemplate object in the GoCD API.
func (pts *PipelineTemplatesService) Create(ctx context.Context, name string, st []*Stage) (*PipelineTemplate, *APIResponse, error) {

	pt := PipelineTemplateRequest{
		Name:   name,
		Stages: st,
	}

	req, err := pts.client.NewRequest("POST", "admin/templates", pt, apiV3)
	if err != nil {
		return nil, nil, err
	}

	ptr := PipelineTemplate{}
	resp, err := pts.client.Do(ctx, req, &ptr, responseTypeJSON)
	if err != nil {
		return nil, resp, err
	}

	ptr.Version = strings.Replace(resp.HTTP.Header.Get("Etag"), "\"", "", -1)

	return &ptr, resp, nil

}

// Update an PipelineTemplate object in the GoCD API.
func (pts *PipelineTemplatesService) Update(ctx context.Context, name string, version string, st []*Stage) (*PipelineTemplate, *APIResponse, error) {
	pt := &PipelineTemplateRequest{
		Name:   name,
		Stages: st,
	}

	req, err := pts.client.NewRequest("PUT", "admin/templates/"+name, pt, apiV3)
	if err != nil {
		return nil, nil, err
	}

	req.HTTP.Header.Set("If-Match", fmt.Sprintf("\"%s\"", version))

	ptr := PipelineTemplate{}
	resp, err := pts.client.Do(ctx, req, &ptr, responseTypeJSON)
	if err != nil {
		return nil, resp, err
	}

	ptr.Version = strings.Replace(resp.HTTP.Header.Get("Etag"), "\"", "", -1)

	return &ptr, resp, nil

}

// Delete a PipelineTemplate from the GoCD API.
func (pts *PipelineTemplatesService) Delete(ctx context.Context, uuid string) (string, *APIResponse, error) {
	return pts.client.genericDeleteAction(ctx, fmt.Sprintf("admin/templates/%s", uuid), apiV3)
}
