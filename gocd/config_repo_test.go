package gocd

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestConfigRepo(t *testing.T) {
	setup()
	defer teardown()

	t.Run("Get", testConfigRepoGet)
	t.Run("List", testConfigRepoList)
	t.Run("Create", testConfigRepoCreate)
}

func testConfigRepoGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/admin/config_repos/repo1", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Unexpected HTTP method")
		j, _ := ioutil.ReadFile("test/resources/configrepos.0.json")
		w.Header().Set("Etag", "mock-etag")
		fmt.Fprint(w, string(j))
	})

	repo, _, err := client.ConfigRepos.Get(context.Background(), "repo1")

	assert.Nil(t, err)
	assert.Equal(t, "mock-etag", repo.Version)
	testConfigRepo(t, repo)
}

func testConfigRepoList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/admin/config_repos", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Unexpected HTTP method")
		j, _ := ioutil.ReadFile("test/resources/configrepos.1.json")
		fmt.Fprint(w, string(j))
	})

	repos, _, err := client.ConfigRepos.List(context.Background())

	assert.Nil(t, err)
	assert.Len(t, repos, 1)

	testConfigRepo(t, repos[0])
}

func testConfigRepoCreate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/admin/config_repos", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST", "Unexpected HTTP method")
		j, _ := ioutil.ReadFile("test/resources/configrepos.0.json")
		fmt.Fprint(w, string(j))
	})

	r := ConfigRepo{ID: "repo1", PluginID: "json.config.plugin", Material: Material{Type: "git", Attributes: &MaterialAttributesGit{URL: "https://github.com/config-repo/gocd-json-config-example.git", Branch: "master", AutoUpdate: true}}}
	repo, _, err := client.ConfigRepos.Create(context.Background(), &r)
	if err != nil {
		t.Error(t, err)
	}

	assert.NotNil(t, repo)
}

func testConfigRepo(t *testing.T, repo *ConfigRepo) {

	for _, attribute := range []EqualityTest{
		{repo.Links.Get("Self").URL.String(), "https://ci.example.com/go/api/admin/config_repos/repo1"},
		{repo.Links.Get("Doc").URL.String(), "https://api.gocd.org/#config-repos"},
		{repo.Links.Get("Find").URL.String(), "https://ci.example.com/go/api/admin/config_repos/:id"},
		{repo.ID, "repo1"},
		{repo.PluginID, "json.config.plugin"},
		{repo.Material.Type, "git"},
		{repo.Material.Attributes.(*MaterialAttributesGit).URL, "https://github.com/config-repo/gocd-json-config-example.git"},
		{repo.Material.Attributes.(*MaterialAttributesGit).Branch, "master"},
	} {
		assert.Equal(t, attribute.wanted, attribute.got)
	}
}
