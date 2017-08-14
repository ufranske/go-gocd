package gocd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfiguration(t *testing.T) {
	t.Run("HasAuth", testConfigurationHasAuth)
	t.Run("New", testConfigurationNew)
	t.Run("SanitizeURL", testConfigurationSantizieURL)
}

func testConfigurationSantizieURL(t *testing.T) {
	u := sanitizeURL(nil)
	assert.Nil(t, u)
}

func testConfigurationNew(t *testing.T) {
	c := Configuration{}
	client := c.Client()
	assert.NotNil(t, client)
}

func testConfigurationHasAuth(t *testing.T) {
	c := Configuration{}

	c.Username = "user"
	c.Password = "pass"
	assert.True(t, c.HasAuth())

	c.Username = "user"
	c.Password = ""
	assert.False(t, c.HasAuth())

	c.Username = ""
	c.Password = "pass"
	assert.False(t, c.HasAuth())

	c.Username = ""
	c.Password = ""
	assert.False(t, c.HasAuth())
}
