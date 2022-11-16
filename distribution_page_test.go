package deploygate

import (
	"testing"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
)

func Test_DeleteDistributionsPage(t *testing.T) {
	t.Parallel()

	c, err := NewClient("api_key", "https://deploygate.com/api")
	if err != nil {
		t.Fatal(err)
	}

	r, err := recorder.New("fixtures/distribution/delete_distributions_page")
	if err != nil {
		t.Fatal(err)
	}

	r.AddFilter(func(i *cassette.Interaction) error {
		delete(i.Request.Headers, "Authorization")
		return nil
	})

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
