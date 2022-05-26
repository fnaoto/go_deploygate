package deploygate

import (
	"log"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func Test_DeleteAppDistributionPage(t *testing.T) {
	t.Parallel()

	c, err := NewClient("api_token")
	if err != nil {
		log.Fatal(err)
	}

	r, err := recorder.New("fixtures/apps/delete_app_distributions_page")
	if err != nil {
		log.Fatal(err)
	}

	c.httpClient.Transport = r
	defer r.Stop()

	resp, err := c.DeleteAppDistributionsPage(&DeleteAppDistributionsPageRequest{
		Owner:            "f-naoto832",
		Platform:         "android",
		AppId:            "com.deploygate.sample",
		DistributionName: "Sample (2022/05/26-14:02:08)",
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error {
		t.Error("response caused error")
	}
}
