package mysqlrouter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetRouterStatus(t *testing.T) {
	setupWithoutTLS()

	routerStatus, err := client.GetRouterStatus()
	assert.NoError(t, err)
	assert.NotEmpty(t, routerStatus)
}
