package gocd

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	defaultHome := os.Getenv("HOME")
	os.Setenv("HOME", "/mock/home")
	path, err := ConfigFilePath()

	assert.Nil(t, err)

	assert.Equal(t, path, "")

	os.Setenv("HOME", defaultHome)
}
