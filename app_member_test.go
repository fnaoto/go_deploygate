package deploygate

import (
	"log"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func Test_GetAppMembers(t *testing.T) {
	t.Parallel()

	r, err := recorder.New("fixtures/apps/get_app_members")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Stop()

	c, err := NewClient("api_token")
	if err != nil {
		log.Fatal(err)
	}

	c.httpClient.Transport = r

	resp, err := c.GetAppMembers(&AppMemberConfig{
		Owner:    "owner",
		Platform: "android",
		AppId:    "com.deploygate.sample",
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error {
		t.Error("response caused error")
	}
}