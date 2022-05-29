package deploygate

import (
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func Test_ListEnterpriseSharedTeams(t *testing.T) {
	t.Parallel()

	c, err := NewClient("user_api_key")
	if err != nil {
		t.Fatal(err)
	}

	r, err := recorder.New("fixtures/apps/list_enterprise_shared_teams")
	if err != nil {
		t.Fatal(err)
	}

	c.httpClient.Transport = r
	defer r.Stop()

	resp, err := c.ListEnterpriseSharedTeams(&ListEnterpriseSharedTeamsRequest{
		Enterprise: "test-enterprise",
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error {
		t.Error("response caused error")
	}
}

func Test_AddEnterpriseSharedTeam(t *testing.T) {
	t.Parallel()

	c, err := NewClient("user_api_key")
	if err != nil {
		t.Fatal(err)
	}

	r, err := recorder.New("fixtures/apps/add_enterprise_shared_team")
	if err != nil {
		t.Fatal(err)
	}

	c.httpClient.Transport = r
	defer r.Stop()

	// TODO: Success response is empty and couldn't decode.
	_, err = c.AddEnterpriseSharedTeam(&AddEnterpriseSharedTeamRequest{
		Enterprise:  "test-enterprise",
		SharedTeam:  "test_shared_team",
		Description: "test_description",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func Test_RemoveEnterpriseSharedTeam(t *testing.T) {
	t.Parallel()

	c, err := NewClient("user_api_key")
	if err != nil {
		t.Fatal(err)
	}

	r, err := recorder.New("fixtures/apps/remove_enterprise_shared_team")
	if err != nil {
		t.Fatal(err)
	}

	c.httpClient.Transport = r
	defer r.Stop()

	resp, err := c.RemoveEnterpriseSharedTeam(&RemoveEnterpriseSharedTeamRequest{
		Enterprise: "test-enterprise",
		SharedTeam: "test_shared_team",
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error {
		t.Error("response caused error")
	}
}
