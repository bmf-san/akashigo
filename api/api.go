package api

import "fmt"

const (
	// Endpoint is an akashi publid api endopoint.
	Endpoint = "https://atnd.ak4.jp/api/cooperation"

	// StampEndpoint is a /[AKASHI企業ID]/stamps
	StampEndpoint = Endpoint + StampPath
	// GetTokenEndpoint is a /token/reissue/[ログイン企業ID]
	GetTokenEndpoint = Endpoint + GetTokenPath

	StampPath    = "/%s/stamps"
	GetTokenPath = "/token/reissue/%s"
)

// URL returns a url by endpoint and arguments.
func URL(endpoint string, query string, args ...interface{}) string {
	if query != "" {
		endpoint = endpoint + query
	}

	return fmt.Sprintf(endpoint, args...)
}
