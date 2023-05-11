package mysqlrouter

import (
	"fmt"
	"io"
	"net/http"
)

func (c *Client) request(url string) ([]byte, error) {
	client := &http.Client{}

	if c.Options != nil {
		if c.Options.Transport != nil {
			client.Transport = c.Options.Transport
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if c.Username != "" && c.Password != "" {
		req.SetBasicAuth(c.Username, c.Password)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s got %d", errStatusCode, resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}
