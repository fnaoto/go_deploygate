package deploygate

import (
	"log"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func Test_GetAppTeams(t *testing.T) {
	t.Parallel()

	c, err := NewClient("api_token")
	if err != nil {
		log.Fatal(err)
	}

	r, err := recorder.New("fixtures/apps/get_app_teams")
	if err != nil {
		log.Fatal(err)
	}

	c.httpClient.Transport = r
	defer r.Stop()

	resp, err := c.GetAppTeams(&GetAppTeamsConfig{
		Organizations: "test_group_terraform",
		Platform:      "android",
		AppId:         "com.deploygate.sample",
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error {
		t.Errorf("response caused error")
	}
}

func Test_AddAppTeams(t *testing.T) {
	t.Parallel()

	c, err := NewClient("api_token")
	if err != nil {
		log.Fatal(err)
	}

	r, err := recorder.New("fixtures/apps/add_app_teams")
	if err != nil {
		log.Fatal(err)
	}

	c.httpClient.Transport = r
	defer r.Stop()

	// Success response is empty and couldn't decode.
	_, err = c.AddAppTeams(&AddAppTeamsConfig{
		Organizations: "test_group_terraform",
		Platform:      "android",
		AppId:         "com.deploygate.sample",
		Team:          "test_team",
	})
	if err != nil {
		t.Fatal(err)
	}
}
