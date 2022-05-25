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

	resp, err := c.GetAppMembers(&GetAppMembersRequest{
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

	resp, err := c.AddAppMembers(&AddAppMembersRequest{
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

func Test_RemoveAppMembers(t *testing.T) {
	t.Parallel()

	c, err := NewClient("api_token")
	if err != nil {
		log.Fatal(err)
	}

	r, err := recorder.New("fixtures/apps/delete_app_members")
	if err != nil {
		log.Fatal(err)
	}

	c.httpClient.Transport = r
	defer r.Stop()

	resp, err := c.RemoveAppMembers(&RemoveAppMembersRequest{
		Owner:    "f-naoto832",
		Platform: "android",
		AppId:    "com.deploygate.sample",
		Users:    "naoto-fukuda-test",
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error {
		t.Error("response caused error")
	}
}
