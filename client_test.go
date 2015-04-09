package nexus

import (
	"net/http"
	"net/http/httptest"
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
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check method is GET before going to check other features
		if r.Method != "GET" {
			t.Errorf("Expected method %q; got %q", "GET", r.Method)
		}
		if r.Header == nil {
			t.Errorf("Expected non-nil request Header")
		}
		switch r.URL.Path {
		default:
			t.Errorf("No testing for this case yet : %q", r.URL.Path)
		case "/v1/fetch":
			t.Logf("case %v ", "/v1/fetch OK")
			w.Write([]byte("{\"application\":{\"id\":8,\"name\":\"Development\",\"description\":\"Client development\"},\"files\":{},\"status\":\"success\"}"))
		}
	}))

	defer ts.Close()

	client := NewClient(ts.URL, token)

	response, err := client.Fetch()

	assert.Nil(t, err)
	assert.True(t, response.Application.ID > 0)
}

func Test_FetchWithError(t *testing.T) {
	client := NewClient("http://127.0.0.1/fake", token)

	_, err := client.Fetch()

	assert.NotNil(t, err)
}
