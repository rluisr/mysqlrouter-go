package mysqlrouter

import (
	"errors"
)

const apiVer = "20190715"

// Client holds the configuration for 20190715 version API client.
type Client struct {
	URL      string
	Username string
	Password string
	SkipTLSVerify bool
}

func newClient(url, user, pass string, skipTLSVerify bool) *Client {
	return &Client{
		URL:      url + "/api/" + apiVer,
		Username: user,
		Password: pass,
		SkipTLSVerify: skipTLSVerify,
	}
}

// New creates a new API client.
func New(url, user, pass string, skipTLSVerify bool) (*Client, error) {
	if url == ""  {
		return nil, errors.New(errEmptyClientInformation)
	}

	client := newClient(url, user, pass, skipTLSVerify)

	err := client.verifyConnection()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) verifyConnection() error {
	_, err := c.request(c.URL + "/swagger.json")
	return err
}
