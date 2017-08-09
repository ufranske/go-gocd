package gocd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTaskValidate(t *testing.T) {
	t.Run("Fail", taskValidateFail)
	t.Run("SuccessExec", taskValidateSuccessExec)
	t.Run("SuccessAnt", taskValidateSuccessAnt)
}

func taskValidateSuccessAnt(t *testing.T) {
	antTask := Task{
		Type: "ant",
	}
	assert.NotNil(t, antTask.Validate())

	antTask.Attributes.RunIf = []string{"one", "two"}
	assert.NotNil(t, antTask.Validate())

	antTask.Attributes.BuildFile = "build-file"
	assert.NotNil(t, antTask.Validate())

	antTask.Attributes.Target = "target"
	assert.NotNil(t, antTask.Validate())

	antTask.Attributes.WorkingDirectory = "working-directory"
	assert.Nil(t, antTask.Validate())
}

func taskValidateSuccessExec(t *testing.T) {
	execTask := Task{
		Type: "exec",
	}
	assert.NotNil(t, execTask.Validate())

	execTask.Attributes.RunIf = []string{"one", "two"}
	assert.NotNil(t, execTask.Validate())

	execTask.Attributes.Command = "command-one"
	assert.NotNil(t, execTask.Validate())

	execTask.Attributes.Arguments = []string{"one", "two"}
	assert.NotNil(t, execTask.Validate())

	execTask.Attributes.WorkingDirectory = "one-two-three"
	assert.Nil(t, execTask.Validate())

}

func taskValidateFail(t *testing.T) {
	task := Task{}
	assert.EqualError(t,
		task.Validate(), "Missing `gocd.TaskAttribute` type")

	task.Type = "invalid-task-type"
	assert.EqualError(t,
		task.Validate(), "Unexpected `gocd.Task.Attribute` types")

	task.Type = "exec"
	assert.NotNil(t, task.Validate())

	task.Type = "ant"
	assert.NotNil(t, task.Validate())
}

func TestJobValidate(t *testing.T) {
	t.Run("ValidateJob", jobValidateSuccess)
	t.Run("Exec", jobValidateExecSuccess)
	t.Run("Ant", jobValidateAntSuccess)
	//t.Run("Nant", job_ValidateNantSuccess)
	//t.Run("Rake", job_ValidateRakeSuccess)
	//t.Run("Fetch", job_ValidateFetchSuccess)
	//t.Run("PluggableTask", job_ValidatePluggableTaskSuccess)
}

func jobValidateSuccess(t *testing.T) {
	j := Job{}
	err := j.Validate()
	assert.NotNil(t, err)

	j.Name = "job-name"
	err = j.Validate()
	assert.Nil(t, err)
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
