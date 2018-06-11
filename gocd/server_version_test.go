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
	t.Run("ServerVersion", testServerVersion)
	t.Run("BadServerVersion", testBadServerVersion)
}

func testServerVersion(t *testing.T) {
	setup()
	defer teardown()

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

func testBadServerVersion(t *testing.T) {
	for _, test := range []struct {
		name      string
		id        int
		errString string
	}{
		{name: "Major", id: 2, errString: "strconv.Atoi: parsing \"a\": invalid syntax"},
		{name: "Minor", id: 3, errString: "strconv.Atoi: parsing \"b\": invalid syntax"},
		{name: "Patch", id: 4, errString: "strconv.Atoi: parsing \"c\": invalid syntax"},
	} {
		t.Run(test.name, func(t *testing.T) { testBadServerVersionMajor(t, test.id, test.errString) })
	}
}

func testBadServerVersionMajor(t *testing.T, i int, errString string) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/version", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Unexpected HTTP method")
		assert.Equal(t, apiV1, r.Header.Get("Accept"))

		j, _ := ioutil.ReadFile(
			fmt.Sprintf("test/resources/server-version.v1.%d.json", i),
		)

		fmt.Fprint(w, string(j))
	})

	_, _, err := client.ServerVersion.Get(context.Background())

	assert.EqualError(t, err, errString)
}
