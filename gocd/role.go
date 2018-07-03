package gocd

import (
	"context"
	"fmt"
)

// RoleService describes Actions which can be performed on roles
type RoleService service

// Role represents a type of agent/actor who can access resources perform operations
type Role struct {
	Name       string              `json:"name"`
	Type       string              `json:"type"`
	Attributes *RoleAttributesGoCD `json:"attributes"`
	Version string `json:"version"`
}

// RoleAttributesGoCD are attributes describing a role, in this cae, which users are present in the role.
type RoleAttributesGoCD struct {
	Users        []string                   `json:"users,omitempty"`
	AuthConfigId *string                    `json:"auth_config_id,omitempty"`
	Properties   []*RoleAttributeProperties `json:"properties,omitempty"`
}

// RoleAttributeProperties
type RoleAttributeProperties struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// RoleListWrapper
type RoleListWrapper struct {
	Embedded struct {
		Roles []*Role `json:"roles"`
	} `json:"_embedded"`
}

// Create a role
func (rs *RoleService) Create(ctx context.Context, role *Role) (r *Role, resp *APIResponse, err error) {
	r = &Role{}
	_, resp, err = rs.client.postAction(ctx, &APIClientRequest{
		APIVersion:   apiV1,
		Path:         "admin/security/roles",
		RequestBody:  role,
		ResponseBody: r,
	})

	return
}

// List all roles
func (rs *RoleService) List(ctx context.Context) (r []*Role, resp *APIResponse, err error) {

	wrapper := RoleListWrapper{}

	_, resp, err = rs.client.getAction(ctx, &APIClientRequest{
		APIVersion:   apiV1,
		Path:         "admin/security/roles",
		ResponseBody: &wrapper,
	})

	return wrapper.Embedded.Roles, resp, err
}

// Delete a role by name
func (rs *RoleService) Delete(ctx context.Context, roleName string) (result bool, resp *APIResponse, err error) {
	var msg string
	msg, resp, err = rs.client.deleteAction(
		ctx,
		fmt.Sprintf("admin/security/roles/%s", roleName),
		apiV1,
	)

	expected := fmt.Sprintf("The role '%s' was deleted successfully.", roleName)
	result = (expected == msg)

	return
}
