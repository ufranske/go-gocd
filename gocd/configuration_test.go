package gocd

import (
	"testing"
	"github.com/h2non/gock"
)

type MockClient struct {
}

func TestConfigurationService(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution
	cs := ConfigurationService{}
}
