package gocd

import (
	"github.com/h2non/gock"
	"testing"
)

type MockClient struct {
}

func TestConfigurationService(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution
	cs := ConfigurationService{}
}
