package deploygate

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

func (c *Client) UploadApps(req *UploadAppsRequest) (*UploadAppsResponse, error) {
	path := fmt.Sprintf("/users/%s/platforms/%s/apps", req.Owner, req.Platform)

	fieldname := "file"
	filename := req.File
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	body := &bytes.Buffer{}
	mw := multipart.NewWriter(body)
	fw, err := mw.CreateFormFile(fieldname, filename)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(fw, file)
	if err != nil {
		return nil, err
	}
	contentType := mw.FormDataContentType()
	err = mw.Close()
	if err != nil {
		return nil, err
	}

	resp, err := c.Post(&HttpRequest{
		path: path,
		body: body,
		header: &Header{
			accept:      "application/json",
			contentType: contentType,
		},
	})
	if err != nil {
		return nil, err
	}

	var r *UploadAppsResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
