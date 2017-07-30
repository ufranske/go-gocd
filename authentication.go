package gocd

import "context"

func (c *Client) Login(ctx context.Context) error {
	req, err := c.NewRequest("GET", "api/agents", nil, apiV2)
	if err != nil {
		return err
	}
	req.Http.SetBasicAuth(c.Auth.Username, c.Auth.Password)

	resp, err := c.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	c.cookie = resp.Http.Header["Set-Cookie"][0]

	return nil
}
