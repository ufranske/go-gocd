package gocd

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestElasticAgentProfile(t *testing.T) {
	for _, test := range []struct {
		name          string
		eapCreate     *ElasticAgentProfile
		eapCreateWant *ElasticAgentProfile
		eapGet        string
		eapGetWant    *ElasticAgentProfile
		eapUpdate     *ElasticAgentProfile
		eapUpdateWant *SCm
		eapListWant   []*ElasticAgentProfile
	}{
		{
			name: "basic",
			eapCreate: &ElasticAgentProfile{
				ID:       "mock-id",
				PluginID: "mock-name",
				Properties: []*ConfigProperty{
					{Key: "Image", Value: "alpine:latest"},
					{Key: "Command", Value: ""},
					{Key: "Environment", Value: ""},
					{Key: "MaxMemory", Value: "200M"},
					{Key: "ReservedMemory", Value: "150M"},
				},
			},
			eapCreateWant: &ElasticAgentProfile{
				ID:       "mock-id",
				PluginID: "mock-name",
				Properties: []*ConfigProperty{
					{Key: "Image", Value: "alpine:latest"},
					{Key: "Command", Value: ""},
					{Key: "Environment", Value: ""},
					{Key: "MaxMemory", Value: "200M"},
					{Key: "ReservedMemory", Value: "150M"},
				},
			},
			eapGet: "mock-name",
			eapGetWant: &ElasticAgentProfile{
				ID:       "mock-id",
				PluginID: "mock-name",
				Properties: []*ConfigProperty{
					{Key: "Image", Value: "alpine:latest"},
					{Key: "Command", Value: ""},
					{Key: "Environment", Value: ""},
					{Key: "MaxMemory", Value: "200M"},
					{Key: "ReservedMemory", Value: "150M"},
				},
			},
			eapUpdate: &ElasticAgentProfile{
				ID:       "mock-id",
				PluginID: "mock-name",
				Properties: []*ConfigProperty{
					{Key: "Image", Value: "alpine:latest"},
					{Key: "Command", Value: ""},
					{Key: "Environment", Value: ""},
					{Key: "MaxMemory", Value: "2G"},
					{Key: "ReservedMemory", Value: "1G"},
				},
			},
			eapUpdateWant: &ElasticAgentProfile{
				ID:       "mock-id",
				PluginID: "mock-name",
				Properties: []*ConfigProperty{
					{Key: "Image", Value: "alpine:latest"},
					{Key: "Command", Value: ""},
					{Key: "Environment", Value: ""},
					{Key: "MaxMemory", Value: "2G"},
					{Key: "ReservedMemory", Value: "1G"},
				},
			},
			eapListWant: []*ElasticAgentProfile{
				{
					ID:       "mock-id",
					PluginID: "mock-name",
					Properties: []*ConfigProperty{
						{Key: "Image", Value: "alpine:latest"},
						{Key: "Command", Value: ""},
						{Key: "Environment", Value: ""},
						{Key: "MaxMemory", Value: "2G"},
						{Key: "ReservedMemory", Value: "1G"},
					},
				},
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			if runIntegrationTest(t) {
				ctx := context.Background()

				t.Run("create", func(t *testing.T) {
					eapCreateGot, _, err := intClient.ElasticAgentProfiles.Create(ctx, test.eapCreate)
					assert.NoError(t, err)
					assert.Equal(t, test.eapCreateWant, eapCreateGot)
				})

				t.Run("get", func(t *testing.T) {
					eapGetGot, _, err := intClient.ElasticAgentProfiles.Get(ctx, test.eapGet)
					assert.NoError(t, err)
					assert.Equal(t, test.eapGetWant, eapGetGot)
				})

				t.Run("update", func(t *testing.T) {
					eapUpdateGot, _, err := intClient.ElasticAgentProfiles.Update(ctx, test.eapUpdate.Name, test.eapUpdate)
					assert.NoError(t, err)
					assert.Equal(t, test.eapUpdateWant, eapUpdateGot)
				})

				t.Run("list", func(t *testing.T) {
					eapListGot, _, err := intClient.ElasticAgentProfiles.List(ctx)
					assert.NoError(t, err)
					assert.Equal(t, test.eapListWant, eapListGot)
				})
			}
		})
	}
}
