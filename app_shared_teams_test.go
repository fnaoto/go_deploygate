package deploygate

import (
	"log"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

// TODO: Test on enterprise account.
func Test_ListAppSharedTeams(t *testing.T) {
	t.Parallel()

	c, err := NewClient("api_key")
	if err != nil {
		log.Fatal(err)
	}

	r, err := recorder.New("fixtures/apps/list_app_shared_teams")
	if err != nil {
		log.Fatal(err)
	}

	c.httpClient.Transport = r
	defer r.Stop()

	resp, err := c.ListAppSharedTeams(&ListAppSharedTeamsRequest{
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
