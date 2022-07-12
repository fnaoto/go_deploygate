package deploygate

import (
	"testing"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
)

func Test_ListOrganizationTeamMembers(t *testing.T) {
	t.Parallel()

	c, err := NewClient("user_api_token")
	if err != nil {
		t.Error(err)
	}

	r, err := recorder.New("fixtures/organization/list_organization_team_members")
	if err != nil {
		t.Error(err)
	}

	r.AddFilter(func(i *cassette.Interaction) error {
		delete(i.Request.Headers, "Authorization")
		return nil
	})

	c.HttpClient.Transport = r
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

func Test_AddOrganizationTeamMember(t *testing.T) {
	t.Parallel()

	c, err := NewClient("user_api_token")
	if err != nil {
		t.Error(err)
	}

	r, err := recorder.New("fixtures/organization/add_organization_team_member")
	if err != nil {
		t.Error(err)
	}

	r.AddFilter(func(i *cassette.Interaction) error {
		delete(i.Request.Headers, "Authorization")
		return nil
	})

	c.HttpClient.Transport = r
	defer r.Stop()

	// TODO: Success response is empty and couldn't decode.
	_, err = c.AddOrganizationTeamMember(&AddOrganizationTeamMemberRequest{
		Organization: "test_organization",
		Team:         "test_team",
		User:         "naoto-fukuda-test",
	})
	if err != nil {
		t.Error(err)
	}
}

func Test_RemoveOrganizationTeamMember(t *testing.T) {
	t.Parallel()

	c, err := NewClient("user_api_token")
	if err != nil {
		t.Error(err)
	}

	r, err := recorder.New("fixtures/organization/remove_organization_team_member")
	if err != nil {
		t.Error(err)
	}

	r.AddFilter(func(i *cassette.Interaction) error {
		delete(i.Request.Headers, "Authorization")
		return nil
	})

	c.HttpClient.Transport = r
	defer r.Stop()

	resp, err := c.RemoveOrganizationTeamMember(&RemoveOrganizationTeamMemberRequest{
		Organization: "test_organization",
		Team:         "test_team",
		User:         "naoto-fukuda-test",
	})
	if err != nil {
		t.Error(err)
	}

	if resp.Error {
		t.Error("response caused error")
	}
}
