package deploygate

import "fmt"

func (c *Client) GetAppTeams(cfg *AppTeamConfig) (*GetAppTeamsResponse, error) {
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
