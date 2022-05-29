package deploygate

import (
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func Test_ListEnterpriseOrganizationMembers(t *testing.T) {
	t.Parallel()

	c, err := NewClient("user_api_key")
	if err != nil {
		t.Fatal(err)
	}

	r, err := recorder.New("fixtures/apps/list_enterprise_organization_members")
	if err != nil {
		t.Fatal(err)
	}

	c.httpClient.Transport = r
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

func Test_AddEnterpriseOrganizationMembers(t *testing.T) {
	t.Parallel()

	c, err := NewClient("user_api_key")
	if err != nil {
		t.Fatal(err)
	}

	r, err := recorder.New("fixtures/apps/add_enterprise_organization_members")
	if err != nil {
		t.Fatal(err)
	}

	c.httpClient.Transport = r
	defer r.Stop()

	// TODO: Success response is empty and couldn't decode.
	_, err = c.AddEnterpriseOrganizationMembers(&AddEnterpriseOrganizationMembersRequest{
		Enterprise:   "test-enterprise",
		Organization: "test_organization",
		User:         "naoto-fukuda-test",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func Test_RemoveEnterpriseOrganizationMembers(t *testing.T) {
	t.Parallel()

	c, err := NewClient("user_api_key")
	if err != nil {
		t.Fatal(err)
	}

	r, err := recorder.New("fixtures/apps/remove_enterprise_organization_members")
	if err != nil {
		t.Fatal(err)
	}

	c.httpClient.Transport = r
	defer r.Stop()

	resp, err := c.RemoveEnterpriseOrganizationMembers(&RemoveEnterpriseOrganizationMembersRequest{
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
