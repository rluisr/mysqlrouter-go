package mysqlrouter

import (
	"encoding/json"
	"fmt"
	"time"
)

var (
	routeURL = "/routes"
)

type RoutesResponse struct {
	Item []*Routes `json:"items"`
}

type Routes struct {
	Name string `json:"name"`
}

type RouteStatus struct {
	ActiveConnections int `json:"activeConnections"`
	TotalConnections  int `json:"totalConnections"`
	BlockedHosts      int `json:"blockedHosts"`
}

type RouteHealth struct {
	IsAlive bool `json:"isAlive"`
}

type RouteDestinationsResponse struct {
	Item []*RouteDestinations `json:"items"`
}

type RouteDestinations struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

type RouteConnectionsResponse struct {
	Item []*RouteConnections `json:"items"`
}

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
