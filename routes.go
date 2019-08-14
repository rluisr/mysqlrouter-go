package mysqlrouter

import "encoding/json"

var (
	routeURL = "/routes"
)

type Routes struct {
	Item []Route `json:"items"`
}

type Route struct {
	Name string `json:"name"`
}

func (c *Client) GetAllRoutes() (*Routes, error) {
	b, err := c.request(c.URL + routeURL)
	if err != nil {
		return nil, err
	}

	var routes *Routes
	err = json.Unmarshal(b, &routes)
	if err != nil {
		return nil, err
	}

	return routes, nil
}
