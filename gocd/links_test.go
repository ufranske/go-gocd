package gocd

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestLinks(t *testing.T) {
	t.Run("MarshallJSON", testMarshallJSON)
	t.Run("Keys", testLinkKeys)
	t.Run("GetOk", testLinkGetOk)
}

func testLinkGetOk(t *testing.T) {
	u, _ := url.Parse("http://example.com")
	l := HALLinks{
		links: []*HALLink{
			{Name: "test-link", URL: u},
			{Name: "example", URL: u},
		},
	}

	l1, ok := l.GetOk("test-link")
	assert.True(t, ok)
	assert.NotNil(t, l1)

	l2, ok := l.GetOk("non-existance")
	assert.False(t, ok)
	assert.Nil(t, l2)

}

func testLinkKeys(t *testing.T) {
	u, _ := url.Parse("http://example.com")
	l := HALLinks{
		links: []*HALLink{
			{Name: "test-link", URL: u},
			{Name: "example", URL: u},
		},
	}

	assert.Equal(t, []string{"test-link", "example"}, l.Keys())

}

func testMarshallJSON(t *testing.T) {
	u, _ := url.Parse("http://example.com")
	l := HALLinks{
		links: []*HALLink{{Name: "test-link", URL: u}},
	}

	b, err := l.MarshallJSON()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "{\"test-link\":{\"href\":\"http://example.com\"}}", string(b))
}
