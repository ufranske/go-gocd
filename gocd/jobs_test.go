package gocd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJob_ValidateSuccess(t *testing.T) {
	t.Run("Exec", jobValidateExecSuccess)
	t.Run("Ant", jobValidateAntSuccess)
	//t.Run("Nant", job_ValidateNantSuccess)
	//t.Run("Rake", job_ValidateRakeSuccess)
	//t.Run("Fetch", job_ValidateFetchSuccess)
	//t.Run("PluggableTask", job_ValidatePluggableTaskSuccess)
}

func jobValidateExecSuccess(t *testing.T) {
	err := (&TaskAttributes{
		RunIf:            []string{"runif-exec"},
		Command:          "my-test-command",
		Arguments:        []string{"arg1", "arg2"},
		WorkingDirectory: "test-working-diretory",
	}).ValidateExec()
	assert.Nil(t, err)
}

func jobValidateAntSuccess(t *testing.T) {
	err := (&TaskAttributes{
		RunIf:            []string{"runif-ant"},
		BuildFile:        "test-build-file",
		Target:           "test-target",
		WorkingDirectory: "test-working-directory",
	}).ValidateAnt()
	assert.Nil(t, err)
}

//func job_ValidateNantSuccess(t *testing.T) {
//	err := (&TaskAttributes{}).ValidateNant()
//	assert.Nil(t, err)
//
//}
//
//func job_ValidateRakeSuccess(t *testing.T) {
//	err := (&TaskAttributes{}).ValidateRake()
//	assert.Nil(t, err)
//
//}
//
//func job_ValidateFetchSuccess(t *testing.T) {
//	err := (&TaskAttributes{}).ValidateFetch()
//	assert.Nil(t, err)
//
//}
//
//func job_ValidatePluggableTaskSuccess(t *testing.T) {
//	err := (&TaskAttributes{}).ValidatePluggableTask()
//	assert.Nil(t, err)
//
//}
