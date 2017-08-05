package gocd

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestJob_ValidateSuccess(t *testing.T) {
	t.Run("Exec", job_ValidateExecSuccess)
	t.Run("Ant", job_ValidateAntSuccess)
	t.Run("Nant", job_ValidateNantSuccess)
	t.Run("Rake", job_ValidateRakeSuccess)
	t.Run("Fetch", job_ValidateFetchSuccess)
	t.Run("PluggableTask", job_ValidatePluggableTaskSuccess)
}

func job_ValidateExecSuccess(t *testing.T) {
	err := TaskAttributes{}.ValidateExec()
	assert.Nil(t, err)
}

func job_ValidateAntSuccess(t *testing.T) {
	err := TaskAttributes{}.ValidateAnt()
	assert.Nil(t, err)

}

func job_ValidateNantSuccess(t *testing.T) {
	err := TaskAttributes{}.ValidateNant()
	assert.Nil(t, err)

}

func job_ValidateRakeSuccess(t *testing.T) {
	err := TaskAttributes{}.ValidateRake()
	assert.Nil(t, err)

}

func job_ValidateFetchSuccess(t *testing.T) {
	err := TaskAttributes{}.ValidateFetch()
	assert.Nil(t, err)

}

func job_ValidatePluggableTaskSuccess(t *testing.T) {
	err := TaskAttributes{}.ValidatePluggableTask()
	assert.Nil(t, err)

}
