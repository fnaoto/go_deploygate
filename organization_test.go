package deploygate

import (
	"testing"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
)

func Test_ListOrganizations(t *testing.T) {
	t.Parallel()

	c, err := NewClient("api_key", "https://deploygate.com/api")
	if err != nil {
		t.Error(err)
	}

	r, err := recorder.New("fixtures/organization/list_organizations")
	if err != nil {
		t.Error(err)
	}

	r.AddFilter(func(i *cassette.Interaction) error {
		delete(i.Request.Headers, "Authorization")
		return nil
	})

	c.HttpClient.Transport = r
	defer r.Stop()

	resp, err := c.ListOrganizations()
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error {
		t.Errorf("response caused error: %s", resp.Message)
	}
}

func Test_CreateOrganization(t *testing.T) {
	t.Parallel()

	c, err := NewClient("api_key", "https://deploygate.com/api")
	if err != nil {
		t.Error(err)
	}

	r, err := recorder.New("fixtures/organization/create_organization")
	if err != nil {
		t.Error(err)
	}

	r.AddFilter(func(i *cassette.Interaction) error {
		delete(i.Request.Headers, "Authorization")
		return nil
	})

	c.HttpClient.Transport = r
	defer r.Stop()

	resp, err := c.CreateOrganization(&CreateOrganizationRequest{
		Name:        "test_organization",
		Description: "test_description",
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error {
		t.Errorf("response caused error: %s", resp.Message)
	}
}

func Test_GetOrganization(t *testing.T) {
	t.Parallel()

	c, err := NewClient("api_key", "https://deploygate.com/api")
	if err != nil {
		t.Error(err)
	}

	r, err := recorder.New("fixtures/organization/get_organization")
	if err != nil {
		t.Error(err)
	}

	r.AddFilter(func(i *cassette.Interaction) error {
		delete(i.Request.Headers, "Authorization")
		return nil
	})

	c.HttpClient.Transport = r
	defer r.Stop()

	resp, err := c.GetOrganization(&GetOrganizationRequest{
		Name: "test_organization",
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error {
		t.Errorf("response caused error: %s", resp.Message)
	}
}

func Test_UpdateOrganization(t *testing.T) {
	t.Parallel()

	c, err := NewClient("api_key", "https://deploygate.com/api")
	if err != nil {
		t.Error(err)
	}

	r, err := recorder.New("fixtures/organization/update_organization")
	if err != nil {
		t.Error(err)
	}

	r.AddFilter(func(i *cassette.Interaction) error {
		delete(i.Request.Headers, "Authorization")
		return nil
	})

	c.HttpClient.Transport = r
	defer r.Stop()

	resp, err := c.UpdateOrganization(&UpdateOrganizationRequest{
		Name:        "test_organization",
		Description: "test_new_description",
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error {
		t.Errorf("response caused error: %s", resp.Message)
	}
}

func Test_DeleteOrganization(t *testing.T) {
	t.Parallel()

	c, err := NewClient("api_key", "https://deploygate.com/api")
	if err != nil {
		t.Error(err)
	}

	r, err := recorder.New("fixtures/organization/delete_organization")
	if err != nil {
		t.Error(err)
	}

	r.AddFilter(func(i *cassette.Interaction) error {
		delete(i.Request.Headers, "Authorization")
		return nil
	})

	c.HttpClient.Transport = r
	defer r.Stop()

	resp, err := c.DeleteOrganization(&DeleteOrganizationRequest{
		Name: "test_organization",
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error {
		t.Errorf("response caused error: %s", resp.Message)
	}
}
