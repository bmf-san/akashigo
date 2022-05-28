package akashigo

import "github.com/bmf-san/akashigo/client"

func New(apiToken string, companyID string) *client.Client {
	return client.New(apiToken, companyID)
}
