package deploygate

import (
	"testing"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
)

func Test_ListEnterpriseOrganizationMembers(t *testing.T) {
	t.Parallel()

	c, err := NewClient("user_api_token")
	if err != nil {
		t.Fatal(err)
	}

	r, err := recorder.New("fixtures/enterprise/list_enterprise_organization_members")
	if err != nil {
		t.Fatal(err)
	}

	r.AddFilter(func(i *cassette.Interaction) error {
		delete(i.Request.Headers, "Authorization")
		return nil
	})

	c.HttpClient.Transport = r
	defer r.Stop()

	resp, err := c.ListEnterpriseOrganizationMembers(&ListEnterpriseOrganizationMembersRequest{
		Enterprise:   "test-enterprise",
		Organization: "test_organization",
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error {
		t.Error("response caused error")
	}
}

func Test_AddEnterpriseOrganizationMember(t *testing.T) {
	t.Parallel()

	c, err := NewClient("user_api_token")
	if err != nil {
		t.Fatal(err)
	}

	r, err := recorder.New("fixtures/enterprise/add_enterprise_organization_member")
	if err != nil {
		t.Fatal(err)
	}

	r.AddFilter(func(i *cassette.Interaction) error {
		delete(i.Request.Headers, "Authorization")
		return nil
	})

	c.HttpClient.Transport = r
	defer r.Stop()

	// TODO: Success response is empty and couldn't decode.
	_, err = c.AddEnterpriseOrganizationMember(&AddEnterpriseOrganizationMemberRequest{
		Enterprise:   "test-enterprise",
		Organization: "test_organization",
		User:         "naoto-fukuda-test",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func Test_RemoveEnterpriseOrganizationMember(t *testing.T) {
	t.Parallel()

	c, err := NewClient("user_api_token")
	if err != nil {
		t.Fatal(err)
	}

	r, err := recorder.New("fixtures/enterprise/remove_enterprise_organization_member")
	if err != nil {
		t.Fatal(err)
	}

	r.AddFilter(func(i *cassette.Interaction) error {
		delete(i.Request.Headers, "Authorization")
		return nil
	})

	c.HttpClient.Transport = r
	defer r.Stop()

	resp, err := c.RemoveEnterpriseOrganizationMember(&RemoveEnterpriseOrganizationMemberRequest{
		Enterprise:   "test-enterprise",
		Organization: "test_organization",
		User:         "naoto-fukuda-test",
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error {
		t.Error("response caused error")
	}
}
