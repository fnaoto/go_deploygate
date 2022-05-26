package deploygate

import (
	"log"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func Test_ListOrganizationMembers(t *testing.T) {
	t.Parallel()

	c, err := NewClient("user_api_token")
	if err != nil {
		log.Fatal(err)
	}

	r, err := recorder.New("fixtures/apps/list_organization_members")
	if err != nil {
		log.Fatal(err)
	}

	c.httpClient.Transport = r
	defer r.Stop()

	resp, err := c.ListOrganizationMembers(&ListOrganizationMembersRequest{
		Organization: "test_organization",
	})
	if err != nil {
		log.Fatal(err)
	}

	if resp.Error {
		t.Error("response caused error")
	}
}
