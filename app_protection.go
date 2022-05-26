package deploygate

import (
	"fmt"
)

func (c *Client) EnableAppProtection(req *EnableAppProtectionRequest) (*EnableAppProtectionResponse, error) {
	path := fmt.Sprintf("users/%s/platforms/%s/apps/%s/binaries/%d/protect", req.Owner, req.Platform, req.AppId, req.Revision)

	resp, err := c.Post(&HttpRequest{
		path: path,
	})
	if err != nil {
		return nil, err
	}

	var r *EnableAppProtectionResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
