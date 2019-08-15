package mysqlrouter

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func (c *Client) request(url string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Username, c.Password)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(errStatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}
