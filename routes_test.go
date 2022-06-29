package mysqlrouter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	route string
)

func TestClient_GetAllRoutes(t *testing.T) {
	setupWithoutTLS()

	routes, err := client.GetAllRoutes()
	assert.NoError(t, err)
	assert.NotEmpty(t, routes)

	route = routes[1].Name
}

func TestClient_GetRouteStatus(t *testing.T) {
	setupWithoutTLS()

	_, err := client.GetRouteStatus(route)
	assert.NoError(t, err)
	// A newborn cluster has no route status.
	//assert.NotEmpty(t, routeStatus)
}

func TestClient_GetRouteHealth(t *testing.T) {
	setupWithoutTLS()

	routeHealth, err := client.GetRouteHealth(route)
	assert.NoError(t, err)
	assert.NotEmpty(t, routeHealth)
}

func TestClient_GetRouteDestinations(t *testing.T) {
	setupWithoutTLS()

	routeDestinations, err := client.GetRouteDestinations(route)
	assert.NoError(t, err)
	assert.NotEmpty(t, routeDestinations)
}

func TestClient_GetRouteConnections(t *testing.T) {
	setupWithoutTLS()

	_, err := client.GetRouteConnections(route)
	assert.NoError(t, err)
}
