package gocd

import (
	"context"
)

// RoleService describes Actions which can be performed on roles
type RoleService service

// Role represents a type of agent/actor who can access resources perform operations
type Role struct {
	Name       string              `json:"name"`
	Type       string              `json:"type"`
	Attributes *RoleAttributesGoCD `json:"attributes"`
}

// RoleAttributesGoCD are attributes describing a role, in this cae, which users are present in the role.
type RoleAttributesGoCD struct {
	Users []string `json:"users"`
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
