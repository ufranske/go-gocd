package gocd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStages(t *testing.T) {
	t.Run("Cleans", testStagesClean)
}

func testStagesClean(t *testing.T) {
	s := Stage{Approval: &Approval{
		Type:          "success",
		Authorization: &Authorization{},
	}}

	assert.NotNil(t, s.Approval.Authorization)
	s.Clean()
	assert.Nil(t, s.Approval.Authorization)
}
