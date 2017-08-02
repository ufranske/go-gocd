package gocd

import (
	"context"
)

func (c *Client) genericHeadAction(ctx context.Context, path string, apiversion string) (bool, *APIResponse, error) {
	u, err := addOptions(path)
	if err != nil {
		return false, nil, err
	}

	req, err := c.NewRequest("HEAD", u, nil, apiversion)
	if err != nil {
		return false, nil, err
	}

	resp, err := c.Do(ctx, req, nil)
	if err != nil {
		return false, resp, err
	}

	exists := resp.Http.StatusCode >= 300 || resp.Http.StatusCode < 200

	return exists, resp, nil

}

func (c *Client) genericDeleteAction(ctx context.Context, path string, apiversion string) (string, *APIResponse, error) {
	u, err := addOptions(path)
	if err != nil {
		return "", nil, err
	}

	req, err := c.NewRequest("DELETE", u, nil, apiversion)
	if err != nil {
		return "", nil, err
	}

	a := StringResponse{}
	resp, err := c.Do(ctx, req, &a)
	if err != nil {
		return "", resp, err
	}

	return a.Message, resp, nil

}
