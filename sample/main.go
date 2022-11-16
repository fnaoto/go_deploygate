package main

import (
	"encoding/json"
	"fmt"
	"os"

	deploygate "github.com/fnaoto/go_deploygate"
)

func main() {
	dg_api_key := readTextFromTerminal("deploygate api key")
	// Or use export DEPLOYGATE_API_KEY = "<API KEY FROM deploygate.com>"
	os.Setenv("DEPLOYGATE_API_KEY", dg_api_key)

	dg_app_id := readTextFromTerminal("deploygate app id (etc. com.sample.test)")
	dg_owner := readTextFromTerminal("deploygate app owner")
	dg_app_planform := readTextFromTerminal("app platform (android/ios)")

	client, err := deploygate.DefaultClient()

	if err != nil {
		fmt.Printf("Got client error: %s", err)
		return
	}

	getAppMembersReq := &deploygate.GetAppMembersRequest{
		Owner:    dg_owner,        // app owner name in deploygate
		Platform: dg_app_planform, // app platform (etc. android/ios)
		AppId:    dg_app_id,       // app id (etc. com.sample.test)
	}

	resp, err := client.GetAppMembers(getAppMembersReq)

	if err != nil {
		fmt.Printf("Got get app members error: %s", err)
		return
	}

	json_resp, err := json.Marshal(resp)

	if err != nil {
		fmt.Printf("Got json marshal error: %s", err)
		return
	}

	fmt.Printf("Got result: %+v\n", string(json_resp))
}

func readTextFromTerminal(print_text string) string {
	fmt.Printf("Enter %s: ", print_text)
	var read_text string
	fmt.Scanln(&read_text)
	return read_text
}
