package deploygate

import (
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func Test_DeleteAppDistributionPage(t *testing.T) {
	t.Parallel()

	c, err := NewClient("user_api_token")
	if err != nil {
		t.Fatal(err)
	}

	r, err := recorder.New("fixtures/app/delete_app_distributions_page")
	if err != nil {
		t.Fatal(err)
	}

	c.HttpClient.Transport = r
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
