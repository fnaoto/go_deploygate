package deploygate

import (
	"bytes"
	"fmt"
)

func (c *Client) GetAppMembers(req *GetAppMembersRequest) (*GetAppMembersResponse, error) {
	path := fmt.Sprintf("/users/%s/platforms/%s/apps/%s/members", req.Owner, req.Platform, req.AppId)

	resp, err := c.Get(&HttpRequest{
		path: path,
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

func (c *Client) AddAppMembers(req *AddAppMembersRequest) (*AddAppMembersResponse, error) {
	path := fmt.Sprintf("/users/%s/platforms/%s/apps/%s/members", req.Owner, req.Platform, req.AppId)
	body := fmt.Sprintf("users=%s&role=%s", req.Users, req.Role)

	resp, err := c.Post(&HttpRequest{
		path: path,
		body: bytes.NewBufferString(body),
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

func (c *Client) RemoveAppMembers(req *RemoveAppMembersRequest) (*RemoveAppMembersResponse, error) {
	path := fmt.Sprintf("/users/%s/platforms/%s/apps/%s/members", req.Owner, req.Platform, req.AppId)
	body := fmt.Sprintf("users=%s", req.Users)

	resp, err := c.Delete(&HttpRequest{
		path: path,
		body: bytes.NewBufferString(body),
	})
	if err != nil {
		return nil, err
	}

	var r *RemoveAppMembersResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
