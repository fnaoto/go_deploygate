package deploygate

import (
	"bytes"
	"fmt"
)

func (c *Client) GetAppMembers(cfg *GetAppMembersRequest) (*GetAppMembersResponse, error) {
	path := fmt.Sprintf("/users/%s/platforms/%s/apps/%s/members", cfg.Owner, cfg.Platform, cfg.AppId)

	resp, err := c.Get(&HttpRequest{
		spath: path,
	})
	if err != nil {
		return nil, err
	}

	var r *GetAppMembersResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) AddAppMembers(cfg *AddAppMembersRequest) (*AddAppMembersResponse, error) {
	path := fmt.Sprintf("/users/%s/platforms/%s/apps/%s/members", cfg.Owner, cfg.Platform, cfg.AppId)
	body := fmt.Sprintf("users=%s&role=%s", cfg.Users, cfg.Role)

	resp, err := c.Post(&HttpRequest{
		spath: path,
		body:  bytes.NewBufferString(body),
	})
	if err != nil {
		return nil, err
	}

	var r *AddAppMembersResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) DeleteAppMembers(cfg *DeleteAppMembersRequest) (*DeleteAppMembersResponse, error) {
	path := fmt.Sprintf("/users/%s/platforms/%s/apps/%s/members", cfg.Owner, cfg.Platform, cfg.AppId)
	body := fmt.Sprintf("users=%s", cfg.Users)

	resp, err := c.Delete(&HttpRequest{
		spath: path,
		body:  bytes.NewBufferString(body),
	})
	if err != nil {
		return nil, err
	}

	var r *DeleteAppMembersResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
