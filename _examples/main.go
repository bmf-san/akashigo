package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bmf-san/akashigo/client"
	"github.com/bmf-san/akashigo/gettoken"
	"github.com/bmf-san/akashigo/stamp"
	"github.com/bmf-san/akashigo/types"
)

func reqGetToken(client *client.Client) {
	model := &gettoken.GetToken{
		Client: client,
	}
	gt, err := model.GetToken(gettoken.GetTokenParams{
		Token: client.APIToken,
	})
	if err != nil {
		log.Fatal("%s", err)
	}
	fmt.Printf("%v", gt)
}

func reqStamp(client *client.Client) {
	model := &stamp.Stamp{
		Client: client,
	}
	s, err := model.Stamp(stamp.StampParams{
		Token:    client.APIToken,
		Type:     types.Attendance,
		Timezone: "+09:00",
	})
	if err != nil {
		log.Fatal("%s", err)
	}
	fmt.Printf("%v", s)
}

func main() {
	apiToken := os.Getenv("AKASHI_API_TOKEN")
	companyID := os.Getenv("AKASHI_COMPANY_ID")

	client := client.New(apiToken, companyID)

	// reqGetToken(client)
	reqStamp(client)
}
