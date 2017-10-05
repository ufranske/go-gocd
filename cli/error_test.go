package cli

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
	"testing"
)

func TestError(t *testing.T) {
	t.Run("Basic", testErrorBasic)
	t.Run("Type", testErrorType)
}

func testErrorType(t *testing.T) {
	var err cli.ExitCoder
	err = JSONCliError{}
	_, ok := err.(cli.ExitCoder)
	assert.True(t, ok)
}

func testErrorBasic(t *testing.T) {
	err := NewCliError("TestReqType", nil, errors.New("test-error"))
	assert.Equal(t, `{
  "Error": "test-error",
  "Request": "TestReqType"
}`, err.Error())
}
