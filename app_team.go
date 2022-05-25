package deploygate

import (
	"bytes"
	"fmt"
)

func (c *Client) GetAppTeams(cfg *GetAppTeamsConfig) (*GetAppTeamsResponse, error) {
	path := fmt.Sprintf("/organizations/%s/platforms/%s/apps/%s/teams", cfg.Organizations, cfg.Platform, cfg.AppId)

	resp, err := c.Get(&HttpRequest{
		spath: path,
	})
	if err != nil {
		return nil, err
	}

	var r *GetAppTeamsResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) AddAppTeams(cfg *AddAppTeamsConfig) (*AddAppTeamsResponse, error) {
	path := fmt.Sprintf("/organizations/%s/platforms/%s/apps/%s/teams", cfg.Organizations, cfg.Platform, cfg.AppId)
	body := fmt.Sprintf("team=%s", cfg.Team)

	resp, err := c.Post(&HttpRequest{
		spath: path,
		body:  bytes.NewBufferString(body),
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
