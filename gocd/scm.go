package gocd

import (
	"context"
	"fmt"
)

// SCMsService exposes calls for interacting with SCM objects in the GoCD API.
type SCMsService service

type SCM struct {
	ID             string              `json:"id"`
	Name           string              `json:"name"`
	AutoUpdate     bool                `json:"auto_update"`
	PluginMetadata *SCMMetadata        `json:"plugin_metadata"`
	Configuration  []*SCMConfiguration `json:"configuration"`

	Version string    `json:"version"`
	Links   *HALLinks `json:"links"`
}

type SCMMetadata struct {
	ID      string `json:"id"`
	Version string `json:"version"`
}

type SCMConfiguration struct {
	Key            string `json:"key"`
	Value          string `json:"value"`
	EncryptedValue string `json:"encrypted_value"`
}

func (scms *SCMsService) Create(ctx context.Context, newSCM *SCM) (scm *SCM, resp *APIResponse, err error) {
	var ver string
	scm = &SCM{}
	if ver, err = scms.client.getAPIVersion(ctx, "/api/admin/scms/"); err == nil {
		_, resp, err = scms.client.postAction(ctx, &APIClientRequest{
			Path:         "admin/scms",
			APIVersion:   ver,
			RequestBody:  newSCM,
			ResponseBody: scm,
		})
	}
	return
}

func (scms *SCMsService) Get(ctx context.Context, name string) (scm *SCM, resp *APIResponse, err error) {
	var ver string
	scm = &SCM{}
	if ver, err = scms.client.getAPIVersion(ctx, "/api/admin/scms/:material_name"); err == nil {
		_, resp, err = scms.client.getAction(ctx, &APIClientRequest{
			Path:         fmt.Sprintf("admin/scms/%s", name),
			APIVersion:   ver,
			ResponseBody: scm,
		})
	}
	return
}

func (scms *SCMsService) Update(ctx context.Context, name string, newSCM *SCM) (scm *SCM, resp *APIResponse, err error) {
	var ver string
	scm = &SCM{}
	if ver, err = scms.client.getAPIVersion(ctx, "/api/admin/scms/:material_name"); err == nil {
		_, resp, err = scms.client.putAction(ctx, &APIClientRequest{
			Path:         fmt.Sprintf("admin/scms/%s", name),
			APIVersion:   ver,
			RequestBody:  newSCM,
			ResponseBody: scm,
		})
	}
	return
}

func (scms *SCMsService) List(ctx context.Context) (scmSlice []*SCM, resp *APIResponse, err error) {
	var ver string
	scmSlice = make([]*SCM, 0)
	if ver, err = scms.client.getAPIVersion(ctx, "/api/admin/scms/"); err == nil {
		_, resp, err = scms.client.getAction(ctx, &APIClientRequest{
			Path:         "admin/scms",
			APIVersion:   ver,
			ResponseBody: scmSlice,
		})
	}
	return
}
