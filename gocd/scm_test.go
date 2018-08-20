package gocd

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSCM(t *testing.T) {
	for _, test := range []struct {
		name          string
		scmCreate     *SCM
		scmCreateWant *SCM
		scmListWant   []*SCM
	}{
		{
			name: "basic",
			scmCreate: &SCM{
				ID:             "mock-id",
				Name:           "mock-name",
				AutoUpdate:     true,
				PluginMetadata: &SCMMetadata{},
				Configuration: []*SCMConfiguration{
					{Key: "username", Value: "admin"},
					{Key: "password", EncryptedValue: "1f3rrs9uhn63hd"},
					{Key: "url", Value: "https://github.com/sample/example.git"},
				},
			},
			scmCreateWant: &SCM{
				ID:             "mock-id",
				Name:           "mock-name",
				AutoUpdate:     true,
				PluginMetadata: &SCMMetadata{},
				Configuration: []*SCMConfiguration{
					{Key: "username", Value: "admin"},
					{Key: "password", EncryptedValue: "1f3rrs9uhn63hd"},
					{Key: "url", Value: "https://github.com/sample/example.git"},
				},
			},
			scmListWant: []*SCM{
				{
					ID:             "mock-id",
					Name:           "mock-name",
					AutoUpdate:     true,
					PluginMetadata: &SCMMetadata{},
					Configuration: []*SCMConfiguration{
						{Key: "username", Value: "admin"},
						{Key: "password", EncryptedValue: "1f3rrs9uhn63hd"},
						{Key: "url", Value: "https://github.com/sample/example.git"},
					},
				},
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			if runIntegrationTest(t) {
				ctx := context.Background()
				scmCreateGot, _, err := intClient.SCM.Create(ctx, test.scmCreate)
				assert.NoError(t, err)

				assert.Equal(t, test.scmCreateWant, scmCreateGot)

				scmListGot, _, err := intClient.SCM.List(ctx)
				assert.NoError(t, err)
				assert.Equal(t, test.scmListWant, scmListGot)
			}
		})
	}
}
