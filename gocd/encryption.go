package gocd

import (
	"context"
	"net/url"
)

// ConfigurationService describes the HAL _link resource for the api response object for a pipelineconfig
type EncryptionService service

type CipherText struct {
	EncryptedValue string       `json:"encrypted_value"`
	Links          EncryptLinks `json:"_links"`
}

// AgentsLinks describes the HAL _link resource for the api response object for a collection of agent objects.
//go:generate gocd-response-links-generator -type=EncryptLinks
type EncryptLinks struct {
	Self *url.URL `json:"self"`
	Doc  *url.URL `json:"doc"`
}

func (e *EncryptionService) Encrypt(ctx context.Context, value string) (*CipherText, *APIResponse, error) {
	type plaintext struct {
		Value string `json:"value"`
	}
	v := &plaintext{Value: value}
	req, err := e.client.NewRequest("POST", "admin/encrypt", v, apiV1)
	if err != nil {
		return nil, nil, err
	}

	c := CipherText{}
	resp, err := e.client.Do(ctx, req, &c, responseTypeJSON)
	if err != nil {
		return nil, resp, err
	}

	return &c, resp, nil
}
