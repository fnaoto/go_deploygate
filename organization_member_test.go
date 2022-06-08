package deploygate

import (
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func Test_ListOrganizationMembers(t *testing.T) {
	t.Parallel()

	c, err := NewClient("user_api_token")
	if err != nil {
		t.Error(err)
	}

	r, err := recorder.New("fixtures/organization/list_organization_members")
	if err != nil {
		t.Error(err)
	}

	c.httpClient.Transport = r
	defer r.Stop()

	resp, err := c.ListOrganizationMembers(&ListOrganizationMembersRequest{
		Organization: "test_organization",
	})
	if err != nil {
		t.Error(err)
	}

	if resp.Error {
		t.Error("response caused error")
	}

	t.Logf("response: %+v", resp)
}

func Test_AddOrganizationMemberByUserName(t *testing.T) {
	t.Parallel()

	c, err := NewClient("user_api_token")
	if err != nil {
		t.Error(err)
	}

	r, err := recorder.New("fixtures/organization/add_organization_member_by_user_name")
	if err != nil {
		t.Error(err)
	}

	c.httpClient.Transport = r
	defer r.Stop()

	resp, err := c.AddOrganizationMemberByUserName(&AddOrganizationMemberByUserNameRequest{
		Organization: "test_organization",
		UserName:     "naoto-fukuda-test",
	})
	if err != nil {
		t.Error(err)
	}

	if resp.Error {
		t.Error("response caused error")
	}
}

func Test_AddOrganizationMemberByEmail(t *testing.T) {
	t.Parallel()

	c, err := NewClient("user_api_token")
	if err != nil {
		t.Error(err)
	}

	r, err := recorder.New("fixtures/organization/add_organization_member_by_email")
	if err != nil {
		t.Error(err)
	}

	c.httpClient.Transport = r
	defer r.Stop()

	resp, err := c.AddOrganizationMemberByEmail(&AddOrganizationMemberByEmailRequest{
		Organization: "test_organization",
		Email:        "dummy.email+test@dummy.email",
	})
	if err != nil {
		t.Error(err)
	}

	if resp.Error {
		t.Error("response caused error")
	}
}

func Test_RemoveOrganizationMemberByUserName(t *testing.T) {
	t.Parallel()

	c, err := NewClient("user_api_token")
	if err != nil {
		t.Error(err)
	}

	r, err := recorder.New("fixtures/organization/remove_organization_member_by_user_name")
	if err != nil {
		t.Error(err)
	}

	c.httpClient.Transport = r
	defer r.Stop()

	resp, err := c.RemoveOrganizationMemberByUserName(&RemoveOrganizationMemberByUserNameRequest{
		Organization: "test_organization",
		UserName:     "naoto-fukuda-test",
	})
	if err != nil {
		t.Error(err)
	}

	if resp.Error {
		t.Error("response caused error")
	}
}

func Test_RemoveOrganizationMemberByEmail(t *testing.T) {
	t.Parallel()

	c, err := NewClient("user_api_token")
	if err != nil {
		t.Error(err)
	}

	r, err := recorder.New("fixtures/organization/remove_organization_member_by_email")
	if err != nil {
		t.Error(err)
	}

	c.httpClient.Transport = r
	defer r.Stop()

	resp, err := c.RemoveOrganizationMemberByEmail(&RemoveOrganizationMemberByEmailRequest{
		Organization: "test_organization",
		Email:        "dummy.email+test@dummy.email",
	})
	if err != nil {
		t.Error(err)
	}

	if resp.Error {
		t.Error("response caused error")
	}
}
