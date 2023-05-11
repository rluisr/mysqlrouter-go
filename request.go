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

	// We accept InternalServerError because MySQL Router REST API return 500 when route status does not alive. see: https://github.com/rluisr/mysqlrouter_exporter/issues/30#issue-1703518829
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusInternalServerError {
		return nil, fmt.Errorf("%s got %d", errStatusCode, resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}
