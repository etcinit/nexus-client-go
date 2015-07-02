package nexus

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/etcinit/nexus-client-go/requests"
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

	os.Setenv("NEXUS_SERVER", endpoint)
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

func Test_Ping(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check method is GET before going to check other features
		if r.Method != "POST" {
			t.Errorf("Expected method %q; got %q", "GET", r.Method)
		}
		if r.Header == nil {
			t.Errorf("Expected non-nil request Header")
		}
		switch r.URL.Path {
		default:
			t.Errorf("No testing for this case yet : %q", r.URL.Path)
		case "/v1/ping":
			t.Logf("case %v ", "/v1/ping OK")

			decoder := json.NewDecoder(r.Body)
			var request requests.PingRequest
			decoder.Decode(&request)

			assert.Equal(t, request.Name, "server1")
			assert.Equal(t, request.Message, "all systems go")

			w.WriteHeader(200)
		}
	}))

	defer ts.Close()

	client := NewClient(ts.URL, token)

	client.Ping("server1", "all systems go")
}

func Test_PingWithError(t *testing.T) {
	client := NewClient("http://127.0.0.1/fake", token)

	err := client.Ping("server1", "all systems go")

	assert.NotNil(t, err)
}

func Test_Log(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check method is GET before going to check other features
		if r.Method != "POST" {
			t.Errorf("Expected method %q; got %q", "GET", r.Method)
		}
		if r.Header == nil {
			t.Errorf("Expected non-nil request Header")
		}
		switch r.URL.Path {
		default:
			t.Errorf("No testing for this case yet : %q", r.URL.Path)
		case "/v1/logs":
			t.Logf("case %v ", "/v1/logs OK")

			decoder := json.NewDecoder(r.Body)
			var request requests.LogsRequest
			decoder.Decode(&request)

			assert.Equal(t, request.Name, "server1")
			assert.Equal(t, request.LogName, "system.log")
			assert.Len(t, request.Lines, 2)

			w.WriteHeader(200)
		}
	}))

	defer ts.Close()

	client := NewClient(ts.URL, token)

	err := client.Log("server1", "system.log", []string{"wow", "warning"})

	assert.Nil(t, err)
}

func Test_LogWithError(t *testing.T) {
	client := NewClient("http://127.0.0.1/fake", token)

	err := client.Log("server1", "system.log", []string{"wow", "warning"})

	assert.NotNil(t, err)
}
