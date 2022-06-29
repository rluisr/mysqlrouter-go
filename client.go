package mysqlrouter

import (
	"errors"
	"net/http"
)

const apiVer = "20190715"

// Client holds the configuration for 20190715 version API client.
type Client struct {
	URL      string
	Username string
	Password string
	Options  *Options
}

type Options struct {
	Transport *http.Transport
}

func newClient(url, user, pass string, Options *Options) *Client {
	return &Client{
		URL:      url + "/api/" + apiVer,
		Username: user,
		Password: pass,
		Options:  Options,
	}
}

// New creates a new API client.
func New(url, user, pass string, Options *Options) (*Client, error) {
	if url == "" {
		return nil, errors.New(errEmptyClientInformation)
	}

	client := newClient(url, user, pass, Options)

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
