package mysqlrouter

import (
	"encoding/json"
	"time"
)

var (
	appURL = "/router"
)

// Router general information of MySQL Router
type Router struct {
	ProcessID      int       `json:"processId"`
	ProductEdition string    `json:"productEdition"`
	TimeStarted    time.Time `json:"timeStarted"`
	Version        string    `json:"version"`
	Hostname       string    `json:"hostname"`
}

// GetRouterStatus returns information of MySQL Router
func (c *Client) GetRouterStatus() (*Router, error) {
	b, err := c.request(c.URL + appURL + "/status")
	if err != nil {
		return nil, err
	}

	var r *Router
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
