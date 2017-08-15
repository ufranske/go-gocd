package gocd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnmarshallLinkFieldFail(t *testing.T) {
	d := linkField{}
	err := unmarshallLinkField(d, "missing-field", nil)
	if err != nil {
		t.Error(err)
	}
	assert.EqualError(t, err, "'missing-field' was not present in `map[]`")
}
