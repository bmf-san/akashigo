package client

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

// Client is an api client.
type Client struct {
	HTTPClient *http.Client
	APIToken   string
	CompanyID  string
}

// New creeats a new api client.
func New(apiToken string, companyID string) *Client {
	return &Client{
		HTTPClient: &http.Client{
			Timeout: time.Duration(100 * time.Second),
		},
		APIToken:  apiToken,
		CompanyID: companyID,
	}
}

// NewRequest returns a new request given a method and url.
func (c *Client) NewRequest(method string, header http.Header, body []byte, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	if header != nil {
		req.Header = header
	}

	return req, nil
}

// Do sends an HTTP request and returns an HTTP response.
func (c *Client) Do(req *http.Request) ([]byte, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
