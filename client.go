package mysqlrouter

import (
	"errors"
)

type Client struct {
	URL      string
	Username string
	Password string
}

func newClient(url, user, pass string) *Client {
	return &Client{
		URL:      url + "/api/20190715",
		Username: user,
		Password: pass,
	}
}

func New(url, user, pass string) (*Client, error) {
	if url == "" || user == "" || pass == "" {
		return nil, errors.New(errEmptyClientInformation)
	}

	client := newClient(url, user, pass)

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
