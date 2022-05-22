package deploygate

import (
	"log"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func Test_GetAppMembers(t *testing.T) {
	t.Parallel()

	c, err := NewClient("api_token")
	if err != nil {
		log.Fatal(err)
	}

	r, err := recorder.New("fixtures/apps/get_app_members")
	if err != nil {
		log.Fatal(err)
	}

	c.httpClient.Transport = r
	defer r.Stop()

	resp, err := c.GetAppMembers(&AppMemberConfig{
		Owner:    "f-naoto832",
		Platform: "android",
		AppId:    "com.deploygate.sample",
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error {
		t.Error("response caused error")
	}
}

func Test_AddAppMembers(t *testing.T) {
	t.Parallel()

	c, err := NewClient("api_token")
	if err != nil {
		log.Fatal(err)
	}

	r, err := recorder.New("fixtures/apps/add_app_members")
	if err != nil {
		log.Fatal(err)
	}

	c.httpClient.Transport = r
	defer r.Stop()

	resp, err := c.AddAppMembers(&AppMemberConfig{
		Owner:    "f-naoto832",
		Platform: "android",
		AppId:    "com.deploygate.sample",
		Users:    "naoto-fukuda-test",
		Role:     "2",
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error {
		t.Error("response caused error")
	}
}
