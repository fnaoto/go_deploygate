package deploygate

import (
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func Test_ListOrganizationTeamMembers(t *testing.T) {
	t.Parallel()

	c, err := NewClient("user_api_token")
	if err != nil {
		t.Error(err)
	}

	r, err := recorder.New("fixtures/apps/list_organization_team_members")
	if err != nil {
		t.Error(err)
	}

	c.httpClient.Transport = r
	defer r.Stop()

	resp, err := c.ListOrganizationTeamMembers(&ListOrganizationTeamMembersRequest{
		Organization: "test_organization",
		Team:         "test_team",
	})
	if err != nil {
		t.Error(err)
	}

	if resp.Error {
		t.Error("response caused error")
	}
}
