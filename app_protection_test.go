package deploygate

import (
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func Test_EnableAppProtection(t *testing.T) {
	t.Parallel()

	c, err := NewClient("user_api_token")
	if err != nil {
		t.Fatal(err)
	}

	r, err := recorder.New("fixtures/app/enable_app_protection")
	if err != nil {
		t.Fatal(err)
	}

	c.HttpClient.Transport = r
	defer r.Stop()

	resp, err := c.EnableAppProtection(&EnableAppProtectionRequest{
		Owner:    "f-naoto832",
		Platform: "android",
		AppId:    "com.deploygate.sample",
		Revision: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error {
		t.Error("response caused error")
	}
}

func Test_DisableAppProtection(t *testing.T) {
	t.Parallel()

	c, err := NewClient("user_api_token")
	if err != nil {
		t.Fatal(err)
	}

	r, err := recorder.New("fixtures/app/disable_app_protection")
	if err != nil {
		t.Fatal(err)
	}

	c.HttpClient.Transport = r
	defer r.Stop()

	resp, err := c.DisableAppProtection(&DisableAppProtectionRequest{
		Owner:    "f-naoto832",
		Platform: "android",
		AppId:    "com.deploygate.sample",
		Revision: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error {
		t.Error("response caused error")
	}
}
