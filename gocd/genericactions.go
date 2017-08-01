package gocd

import (
	"context"
)

func (c *Client) genericDeleteAction(ctx context.Context, path string) (string, *APIResponse, error) {
	u, err := addOptions(path)
	if err != nil {
		return "", nil, err
	}

	req, err := c.NewRequest("DELETE", u, nil, apiV4)
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
