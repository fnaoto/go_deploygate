package deploygate

import (
	"bytes"
	"fmt"
)

func (c *Client) ListAppTeams(cfg *ListAppTeamsRequest) (*ListAppTeamsResponse, error) {
	path := fmt.Sprintf("/organizations/%s/platforms/%s/apps/%s/teams", cfg.Organization, cfg.Platform, cfg.AppId)

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

func (c *Client) AddAppTeams(cfg *AddAppTeamsRequest) (*AddAppTeamsResponse, error) {
	path := fmt.Sprintf("/organizations/%s/platforms/%s/apps/%s/teams", cfg.Organization, cfg.Platform, cfg.AppId)
	body := fmt.Sprintf("team=%s", cfg.Team)

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

func (c *Client) RemoveAppTeams(cfg *RemoveAppTeamsRequest) (*RemoveAppTeamsResponse, error) {
	path := fmt.Sprintf("/organizations/%s/platforms/%s/apps/%s/teams/%s", cfg.Organization, cfg.Platform, cfg.AppId, cfg.Team)

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
