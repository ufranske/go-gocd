package gocd

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestApproval(t *testing.T) {
	approval := &Approval{
		Type: "success",
		Authorization: &Authorization{
			Roles: []string{"one"},
		},
	}
	assert.NotNil(t, approval.Authorization)
	approval.Clean()
	assert.Nil(t, approval.Authorization)
}
