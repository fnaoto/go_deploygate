package deploygate

import (
	"bytes"
	"fmt"
)

func (c *Client) ListAppSharedTeams(cfg *ListAppSharedTeamsRequest) (*ListAppSharedTeamsResponse, error) {
	path := fmt.Sprintf("organizations/%s/platforms/%s/apps/%s/shared_teams", cfg.Organization, cfg.Platform, cfg.AppId)

	resp, err := c.Get(&HttpRequest{
		path: path,
	})
	if err != nil {
		return nil, err
	}

	var r *ListAppSharedTeamsResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) AddAppSharedTeams(cfg *AddAppSharedTeamsRequest) (*AddAppSharedTeamsResponse, error) {
	path := fmt.Sprintf("organizations/%s/platforms/%s/apps/%s/shared_teams", cfg.Organization, cfg.Platform, cfg.AppId)
	body := fmt.Sprintf("team=%s", cfg.Team)

	resp, err := c.Post(&HttpRequest{
		path: path,
		body: bytes.NewBufferString(body),
	})
	if err != nil {
		return nil, err
	}

	var r *AddAppSharedTeamsResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) RemoveAppSharedTeams(cfg *RemoveAppSharedTeamsRequest) (*RemoveAppSharedTeamsResponse, error) {
	path := fmt.Sprintf("organizations/%s/platforms/%s/apps/%s/shared_teams/%s", cfg.Organization, cfg.Platform, cfg.AppId, cfg.Team)

	resp, err := c.Delete(&HttpRequest{
		path: path,
	})
	if err != nil {
		return nil, err
	}

	var r *RemoveAppSharedTeamsResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
