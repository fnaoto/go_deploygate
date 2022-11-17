package deploygate

import (
	"net/http"
	"testing"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
)

func Test_UploadApps(t *testing.T) {
	t.Parallel()

	conf := setUpConfigForTest()

	c, err := NewClient(conf)
	if err != nil {
		t.Fatal(err)
	}

	r, err := recorder.New("fixtures/app/upload_app")
	if err != nil {
		t.Fatal(err)
	}

	r.AddSaveFilter(func(i *cassette.Interaction) error {
		delete(i.Request.Headers, "Authorization")
		i.Response.Headers = make(http.Header)
		return nil
	})

	c.HttpClient.Transport = r
	defer r.Stop()

	resp, err := c.UploadApps(&UploadAppsRequest{
		Owner:    "f-naoto832",
		Platform: "android",
		AppId:    "com.deploygate.sample",
		File:     "files/DeployGateSample.apk",
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error {
		t.Errorf("response caused error: %s", resp.Message)
	}
}
