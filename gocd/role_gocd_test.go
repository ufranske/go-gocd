package gocd

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRole(t *testing.T) {
	t.Run("GoCD", testRoleGoCD)
}

func testRoleGoCD(t *testing.T) {

	intSetup()

	if runIntegrationTest() {

		ctx := context.Background()

		roles := []*Role{
			{
				Name: "spacetiger",
				Type: "gocd",
				Attributes: &RoleAttributesGoCD{
					Users: []string{"alice", "bob", "robin"},
				},
			},
			{
				Name: "my-mock-gocd-role",
				Type: "gocd",
				Attributes: &RoleAttributesGoCD{
					Users: []string{"user-one", "user-two"},
				},
			},
			// Currently there's no fixtures to test the plugin roles,
			// so until there is a way, we can not test plugin role types.
			//{
			//	Name: "blackbird",
			//	Type: "plugin",
			//	Attributes: &RoleAttributesGoCD{
			//		AuthConfigId: String("ldap"),
			//		Properties: []*RoleAttributeProperties{
			//			{
			//				Key:   "UserGroupMembershipAttribute",
			//				Value: "memberOf",
			//			},
			//			{
			//				Key:   "GroupIdentifiers",
			//				Value: "ou=admins,ou=groups,ou=system,dc=example,dc=com",
			//			},
			//		},
			//	},
			//},
		}

		// Test role creation
		for _, role := range roles {
			role_response, _, err := intClient.Roles.Create(ctx, role)
			assert.NoError(t, err)
			assert.Equal(t, role, role_response)
		}

		// Test role listing
		roles_response, _, err := intClient.Roles.List(ctx)
		assert.NoError(t, err)

		for i, role_response := range roles_response {
			assert.Equal(t, roles[i], role_response)
		}

		// Test role delete
		for _, role := range roles {
			result, _, err := intClient.Roles.Delete(ctx, role.Name)
			assert.True(t, result)
			assert.NoError(t, err)
		}
		roles_response, _, err = intClient.Roles.List(ctx)
		assert.NoError(t, err)
		assert.Empty(t, roles_response)

	} else {
		skipIntegrationtest(t)
	}
}
