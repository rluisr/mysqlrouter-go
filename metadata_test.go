package mysqlrouter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	metadata string
)

func TestClient_GetAllMetadata(t *testing.T) {
	setupWithoutTLS()

	metadatas, err := client.GetAllMetadata()
	assert.NoError(t, err)
	assert.NotEmpty(t, metadatas)

	metadata = metadatas[0].Name
}

func TestClient_GetMetadataConfig(t *testing.T) {
	setupWithoutTLS()

	metadataConfig, err := client.GetMetadataConfig(metadata)
	assert.NoError(t, err)
	assert.NotEmpty(t, metadataConfig)
}

func TestClient_GetMetadataStatus(t *testing.T) {
	setupWithoutTLS()

	metadataStatus, err := client.GetMetadataStatus(metadata)
	assert.NoError(t, err)
	assert.NotEmpty(t, metadataStatus)
}
