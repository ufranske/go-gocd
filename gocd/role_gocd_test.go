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

	setup()
	defer teardown()

	mux.HandleFunc("/api/admin/security/roles", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST", "Unexpected HTTP method")
		assert.Equal(t, apiV1, r.Header.Get("Accept"))

		j, _ := ioutil.ReadFile("test/resources/role.1.json")

		fmt.Fprint(w, string(j))
	})

	r, _, err := client.Roles.Create(context.Background(),
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
	}, r)

}

func testListGoCDRoles(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/admin/security/roles", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Unexpected HTTP method")
		assert.Equal(t, apiV1, r.Header.Get("Accept"))

		j, _ := ioutil.ReadFile("test/resources/roles.0.json")

		fmt.Fprint(w, string(j))
	})

	r, _, err := client.Roles.List(context.Background())

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
				AutoConfigId: String("ldap"),
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
	}, r)
}
