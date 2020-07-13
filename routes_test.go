package mysqlrouter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	route string
)

func TestClient_GetAllRoutes(t *testing.T) {
	setup()

	routes, err := client.GetAllRoutes()
	assert.NoError(t, err)
	assert.NotEmpty(t, routes)

	route = routes[1].Name
}

func TestClient_GetRouteStatus(t *testing.T) {
	setup()

	routeStatus, err := client.GetRouteStatus(route)
	assert.NoError(t, err)
	assert.NotEmpty(t, routeStatus)
}

func TestClient_GetRouteHealth(t *testing.T) {
	setup()

	routeHealth, err := client.GetRouteHealth(route)
	assert.NoError(t, err)
	assert.NotEmpty(t, routeHealth)
}

func TestClient_GetRouteDestinations(t *testing.T) {
	setup()

	routeDestinations, err := client.GetRouteDestinations(route)
	assert.NoError(t, err)
	assert.NotEmpty(t, routeDestinations)
}

func TestClient_GetRouteConnections(t *testing.T) {
	setup()

	_, err := client.GetRouteConnections(route)
	assert.NoError(t, err)
}
