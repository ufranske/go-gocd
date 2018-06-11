package gocd

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestServerVersion(t *testing.T) {
	setup()
	defer teardown()

	t.Run("ServerVersion", testServerVersion)
}

func testServerVersion(t *testing.T) {
	mux.HandleFunc("/api/version", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Unexpected HTTP method")
		assert.Equal(t, apiV1, r.Header.Get("Accept"))

		j, _ := ioutil.ReadFile("test/resources/server-version.v1.1.json")

		fmt.Fprint(w, string(j))
	})

	v, _, err := client.ServerVersion.Get(context.Background())

	assert.NoError(t, err)

	assert.Equal(t, &ServerVersion{
		Version:     "16.6.0",
		BuildNumber: "3348",
		GitSha:      "a7a5717cbd60c30006314fb8dd529796c93adaf0",
		FullVersion: "16.6.0 (3348-a7a5717cbd60c30006314fb8dd529796c93adaf0)",
		CommitURL:   "https://github.com/gocd/gocd/commits/a7a5717cbd60c30006314fb8dd529796c93adaf0",
		VersionParts: &ServerVersionParts{
			Major: 16,
			Minor: 6,
			Patch: 0,
		},
	}, v)

}
