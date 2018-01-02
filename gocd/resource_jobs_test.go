package gocd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResourceJobs(t *testing.T) {
	for _, test := range []struct {
		expected string
		input    int
	}{
		{expected: "5", input: 5},
	} {
		tf := TimeoutField(test.input)

		b, err := tf.MarshalJSON()

		assert.Equal(t, test.expected, string(b))
		assert.NoError(t, err)
	}
}
