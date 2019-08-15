package mysqlrouter

import (
	"encoding/json"
	"fmt"
	"time"
)

var (
	metadataURL = "/metadata"
)

type MetadataResponse struct {
	Item []*Metadata `json:"items"`
}

type Metadata struct {
	Name string `json:"name"`
}

type MetadataConfig struct {
	ClusterName        string         `json:"clusterName"`
	TimeRefreshInMs    int            `json:"timeRefreshInMs"`
	GroupReplicationID string         `json:"groupReplicationId"`
	Nodes              []MetadataNode `json:"nodes"`
}

type MetadataNode struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
}

type MetadataStatus struct {
	RefreshFailed            int       `json:"refreshFailed"`
	RefreshSucceeded         int       `json:"refreshSucceeded"`
	TimeLastRefreshSucceeded time.Time `json:"timeLastRefreshSucceeded"`
	LastRefreshHostname      string    `json:"lastRefreshHostname"`
	LastRefreshPort          int       `json:"lastRefreshPort"`
}

func (c *Client) GetAllMetadata() ([]*Metadata, error) {
	b, err := c.request(c.URL + metadataURL)
	if err != nil {
		return nil, err
	}

	var r *MetadataResponse
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}

	return r.Item, nil
}

func (c *Client) GetMetadataConfig(cName string) (*MetadataConfig, error) {
	b, err := c.request(fmt.Sprintf("%s%s/%s/config", c.URL, metadataURL, cName))
	if err != nil {
		return nil, err
	}

	var r *MetadataConfig
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) GetMetadataStatus(cName string) (*MetadataStatus, error) {
	b, err := c.request(fmt.Sprintf("%s%s/%s/status", c.URL, metadataURL, cName))
	if err != nil {
		return nil, err
	}

	var r *MetadataStatus
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
