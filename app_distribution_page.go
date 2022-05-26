package deploygate

import (
	"bytes"
	"fmt"
)

func (c *Client) DeleteAppDistributionsPage(req *DeleteAppDistributionsPageRequest) (*DeleteAppDistributionsPageResponse, error) {
	path := fmt.Sprintf("/users/%s/platforms/%s/apps/%s/distributions", req.Owner, req.Platform, req.AppId)
	body := fmt.Sprintf("distribution_name=%s", req.DistributionName)

	resp, err := c.Delete(&HttpRequest{
		path: path,
		body: bytes.NewBufferString(body),
	})
	if err != nil {
		return nil, err
	}

	var r *DeleteAppDistributionsPageResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
