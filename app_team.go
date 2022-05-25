package deploygate

import (
	"bytes"
	"fmt"
)

func (c *Client) ListAppTeams(req *ListAppTeamsRequest) (*ListAppTeamsResponse, error) {
	path := fmt.Sprintf("/organizations/%s/platforms/%s/apps/%s/teams", req.Organization, req.Platform, req.AppId)

	resp, err := c.Get(&HttpRequest{
		path: path,
	})
	if err != nil {
		return nil, err
	}

	var r *ListAppTeamsResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) AddAppTeams(req *AddAppTeamsRequest) (*AddAppTeamsResponse, error) {
	path := fmt.Sprintf("/organizations/%s/platforms/%s/apps/%s/teams", req.Organization, req.Platform, req.AppId)
	body := fmt.Sprintf("team=%s", req.Team)

	resp, err := c.Post(&HttpRequest{
		path: path,
		body: bytes.NewBufferString(body),
	})
	if err != nil {
		return nil, err
	}

	var r *AddAppTeamsResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) RemoveAppTeams(req *RemoveAppTeamsRequest) (*RemoveAppTeamsResponse, error) {
	path := fmt.Sprintf("/organizations/%s/platforms/%s/apps/%s/teams/%s", req.Organization, req.Platform, req.AppId, req.Team)

	resp, err := c.Delete(&HttpRequest{
		path: path,
	})
	if err != nil {
		return nil, err
	}

	var r *RemoveAppTeamsResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
