package deploygate

import (
	"bytes"
	"fmt"
)

func (c *Client) ListAppSharedTeams(req *ListAppSharedTeamsRequest) (*ListAppSharedTeamsResponse, error) {
	path := fmt.Sprintf("organizations/%s/platforms/%s/apps/%s/shared_teams", req.Organization, req.Platform, req.AppId)

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

func (c *Client) AddAppSharedTeam(req *AddAppSharedTeamRequest) (*AddAppSharedTeamResponse, error) {
	path := fmt.Sprintf("organizations/%s/platforms/%s/apps/%s/shared_teams", req.Organization, req.Platform, req.AppId)
	body := fmt.Sprintf("team=%s", req.Team)

	resp, err := c.Post(&HttpRequest{
		path: path,
		body: bytes.NewBufferString(body),
	})
	if err != nil {
		return nil, err
	}

	var r *AddAppSharedTeamResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) RemoveAppSharedTeam(req *RemoveAppSharedTeamRequest) (*RemoveAppSharedTeamResponse, error) {
	path := fmt.Sprintf("organizations/%s/platforms/%s/apps/%s/shared_teams/%s", req.Organization, req.Platform, req.AppId, req.Team)

	resp, err := c.Delete(&HttpRequest{
		path: path,
	})
	if err != nil {
		return nil, err
	}

	var r *RemoveAppSharedTeamResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
