package gocd

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestRole(t *testing.T) {
	t.Run("GoCD", testRoleGoCD)
}

func testRoleGoCD(t *testing.T) {
	t.Run("Create", testCreateGoCDRole)
	t.Run("List", testListGoCDRoles)
}

func testCreateGoCDRole(t *testing.T) {

	if runIntegrationTest() {

		role, _, err := client.Roles.Create(context.Background(),
			&Role{
				Name: "my-mock-gocd-role",
				Type: "gocd",
				Attributes: &RoleAttributesGoCD{
					Users: []string{"user-one", "user-two"},
				},
			},
		)

		assert.NoError(t, err)

		assert.Equal(t, &Role{
			Name: "my-mock-gocd-role",
			Type: "gocd",
			Attributes: &RoleAttributesGoCD{
				Users: []string{"user-one", "user-two"},
			},
		}, role)

		roles, _, err := client.Roles.List(context.Background())

		assert.NoError(t, err)

		assert.Equal(t, []*Role{
			{
				Name: "spacetiger",
				Type: "gocd",
				Attributes: &RoleAttributesGoCD{
					Users: []string{"alice", "bob", "robin"},
				},
			},
			{
				Name: "blackbird",
				Type: "plugin",
				Attributes: &RoleAttributesGoCD{
					AuthConfigId: String("ldap"),
					Properties: []*RoleAttributeProperties{
						{
							Key:   "UserGroupMembershipAttribute",
							Value: "memberOf",
						},
						{
							Key:   "GroupIdentifiers",
							Value: "ou=admins,ou=groups,ou=system,dc=example,dc=com",
						},
					},
				},
			},
		}, roles)
	} else {
		t.Skip("'GOCD_ACC=1' must be set to run integration tests")
	}
})
