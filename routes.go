package mysqlrouter

import (
	"encoding/json"
	"fmt"
	"time"
)

var (
	routeURL = "/routes"
)

// RoutesResponse is the response of getting all routes endpoint.
type RoutesResponse struct {
	Item []*Routes `json:"items"`
}

// Routes is the response of getting all routes endpoint.
type Routes struct {
	Name string `json:"name"`
}

// RouteStatus is the state of the route.
type RouteStatus struct {
	ActiveConnections int `json:"activeConnections"`
	TotalConnections  int `json:"totalConnections"`
	BlockedHosts      int `json:"blockedHosts"`
}

// RouteHealth is the structure of health check of the route.
type RouteHealth struct {
	IsAlive bool `json:"isAlive"`
}

// RouteDestinationsResponse is the response of getting the information of the route destination endpoint.
type RouteDestinationsResponse struct {
	Item []*RouteDestinations `json:"items"`
}

// RouteDestinations is the structure of destination of the route.
type RouteDestinations struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

// RouteConnectionsResponse is the response of getting the information of route connection endpoint.
type RouteConnectionsResponse struct {
	Item []*RouteConnections `json:"items"`
}

// RouteConnections is the structure of the information of the connection of route.
type RouteConnections struct {
	BytesFromServer            int       `json:"bytesFromServer"`
	BytesToServer              int       `json:"bytesToServer"`
	SourceAddress              string    `json:"sourceAddress"`
	DestinationAddress         string    `json:"destinationAddress"`
	TimeStarted                time.Time `json:"timeStarted"`
	TimeConnectedToServer      time.Time `json:"timeConnectedToServer"`
	TimeLastSentToServer       time.Time `json:"timeLastSentToServer"`
	TimeLastReceivedFromServer time.Time `json:"timeLastReceivedFromServer"`
}

// GetAllRoutes returns all routes.
func (c *Client) GetAllRoutes() ([]*Routes, error) {
	b, err := c.request(c.URL + routeURL)
	if err != nil {
		return nil, err
	}

	var r *RoutesResponse
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}

	return r.Item, nil
}

// GetRouteStatus return the status of the route.
func (c *Client) GetRouteStatus(rName string) (*RouteStatus, error) {
	b, err := c.request(fmt.Sprintf("%s%s/%s/status", c.URL, routeURL, rName))
	if err != nil {
		return nil, err
	}

	var rStatus *RouteStatus
	err = json.Unmarshal(b, &rStatus)
	if err != nil {
		return nil, err
	}

	return rStatus, nil
}

// GetRouteHealth returns the health check of the route.
func (c *Client) GetRouteHealth(rName string) (*RouteHealth, error) {
	b, err := c.request(fmt.Sprintf("%s%s/%s/health", c.URL, routeURL, rName))
	if err != nil {
		return nil, err
	}

	var rHealth *RouteHealth
	err = json.Unmarshal(b, &rHealth)
	if err != nil {
		return nil, err
	}

	return rHealth, nil
}

// GetRouteDestinations returns destinations of the route.
func (c *Client) GetRouteDestinations(rName string) ([]*RouteDestinations, error) {
	b, err := c.request(fmt.Sprintf("%s%s/%s/destinations", c.URL, routeURL, rName))
	if err != nil {
		return nil, err
	}

	var r *RouteDestinationsResponse
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}

	return r.Item, nil
}

// GetRouteConnections return connections of the route.
func (c *Client) GetRouteConnections(rName string) ([]*RouteConnections, error) {
	b, err := c.request(fmt.Sprintf("%s%s/%s/connections", c.URL, routeURL, rName))
	if err != nil {
		return nil, err
	}

	var r *RouteConnectionsResponse
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}

	return r.Item, nil
}
