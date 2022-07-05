package deploygate

import (
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func Test_DeleteDistributionsPage(t *testing.T) {
	t.Parallel()

	c, err := NewClient("user_api_token")
	if err != nil {
		t.Fatal(err)
	}

	r, err := recorder.New("fixtures/distribution/delete_distributions_page")
	if err != nil {
		t.Fatal(err)
	}

	c.HttpClient.Transport = r
	defer r.Stop()

	resp, err := c.DeleteDistributionsPage(&DeleteDistributionsPageRequest{
		Distribution: "test_distribution",
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error {
		t.Error("response caused error")
	}
}
