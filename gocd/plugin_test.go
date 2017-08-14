package gocd

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"context"
	"net/http"
	"fmt"
	"io/ioutil"
)

func TestPluginApi(t *testing.T) {
	setup()
	defer teardown()

	t.Run("List", testPluginApiList)
	t.Run("Get", testPluginApiGet)
}

func testPluginApiList(t *testing.T) {
	mux.HandleFunc("/api/admin/plugin_info", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Unexpected HTTP method")
		j, _ := ioutil.ReadFile("test/resources/plugin.0.json")
		fmt.Fprint(w, string(j))
	})

	plugins, _, err := client.Plugins.List(context.Background())
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, plugins)

	assert.NotNil(t, plugins.Links.Doc)
	assert.Equal(t, "https://api.gocd.org/#plugin-info", plugins.Links.Doc.String())
	assert.NotNil(t, plugins.Links.Self)
	assert.Equal(t, "https://ci.example.com/go/api/admin/plugin_info", plugins.Links.Self.String())

	assert.NotNil(t, plugins.Embedded)
	assert.NotNil(t, plugins.Embedded.PluginInfo)
	assert.Len(t, plugins.Embedded.PluginInfo, 1)

	pi := plugins.Embedded.PluginInfo[0]
	assert.NotNil(t, pi.Links)
	assert.Equal(t, "https://ci.example.com/go/api/admin/plugin_info/plugin_id", pi.Links.Self.String())
	assert.Equal(t, "https://api.gocd.org/#plugin-info", pi.Links.Doc.String())
	assert.Equal(t, "https://ci.example.com/go/api/admin/plugin_info/:id", pi.Links.Find.String())

	assert.Equal(t, "plugin_id", pi.ID, )
	assert.Equal(t, "SCM Plugin", pi.Name)
	assert.Equal(t, "SCM Plugin For HG", pi.DisplayName)
	assert.Equal(t, "1.2.3", pi.Version)
	assert.Equal(t, "scm", pi.Type)
}

func testPluginApiGet(t *testing.T) {
	assert.Nil(t, nil)
}
