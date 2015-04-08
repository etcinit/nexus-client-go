package nexus

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const endpoint = "http://127.0.0.1"
const token = "Vcqda7FKQKpCCG3qof4KFnsMivKpaz2WNBT7d7iVqyRv7UoWo17hY0N0x5BZb42w"

func Test_NewClient(t *testing.T) {
	NewClient(endpoint, token)
}

func Test_NewClientFromEnv(t *testing.T) {
	_, err := NewClientFromEnv()

	assert.NotNil(t, err)

	os.Setenv("NEXUS_ENDPOINT", endpoint)
	os.Setenv("NEXUS_APIKEY", token)

	client, err2 := NewClientFromEnv()

	assert.Nil(t, err2)
	assert.Equal(t, endpoint, client.endpoint)
	assert.Equal(t, token, client.token)
}

func Test_Fetch(t *testing.T) {
	client := NewClient(endpoint, token)

	response, err := client.Fetch()

	assert.Nil(t, err)
	assert.True(t, response.Application.ID > 0)
}
