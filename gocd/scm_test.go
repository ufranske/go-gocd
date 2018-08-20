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
		scmGet        string
		scmGetWant    *SCM
		scmUpdate     *SCM
		scmUpdateWant *SCm
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
			scmGet: "mock-name",
			scmGetWant: &SCM{
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
			scmUpdate: &SCM{
				ID:             "mock-id",
				Name:           "updated-mock-name",
				AutoUpdate:     true,
				PluginMetadata: &SCMMetadata{},
				Configuration: []*SCMConfiguration{
					{Key: "username", Value: "admin"},
					{Key: "password", EncryptedValue: "1f3rrs9uhn63hd"},
					{Key: "url", Value: "https://github.com/sample/example.git"},
				},
			},
			scmUpdateWant: &SCM{
				ID:             "mock-id",
				Name:           "updated-mock-name",
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
				
				t.Run("create", func(t *testing.T) {
					scmCreateGot, _, err := intClient.SCM.Create(ctx, test.scmCreate)
					assert.NoError(t, err)
					assert.Equal(t, test.scmCreateWant, scmCreateGot)
				})

				t.Run("get", func(t *testing.T) {
					scmGetGot, _, err := intClient.SCM.Get(ctx, test.scmGet)
					assert.NoError(t, err)
					assert.Equal(t, test.scmGetWant, scmGetGot)
				})

				t.Run("update", func(t *testing.T) {
					scmUpdateGot, _, err := intClient.SCM.Update(ctx, test.scmUpdate.Name, test.scmUpdate)
					assert.NoError(t, err)
					assert.Equal(t, test.scmUpdateWant, scmUpdateGot)
				})

				t.Run("list", func(t *testing.T) {
					scmListGot, _, err := intClient.SCM.List(ctx)
					assert.NoError(t, err)
					assert.Equal(t, test.scmListWant, scmListGot)
				})
			}
		})
	}
}
