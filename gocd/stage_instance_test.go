package gocd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStageInstance(t *testing.T) {
	t.Run("Validate", testStageInstanceValidate)
	t.Run("JSONStringFail", testStageInstanceJSONStringFail)
	t.Run("JSONString", testStageInstanceJSONString)
}

func testStageInstanceJSONStringFail(t *testing.T) {
	s := StageInstance{
		ApprovedBy: "admin",
		ID:         13,
	}
	_, err := s.JSONString()
	assert.EqualError(t, err, "`gocd.Stage.Name` is empty")
}

func testStageInstanceJSONString(t *testing.T) {
	s := StageInstance{
		Name:       "stage1",
		ApprovedBy: "admin",
		Jobs: []*Job{
			{
				Name:          "job1",
				Result:        "Failed",
				State:         "Completed",
				ID:            13,
				ScheduledDate: 1436172201081,
			},
		},
		CanRun:            true,
		Result:            "Failed",
		ApprovalType:      "success",
		Counter:           "1",
		ID:                13,
		OperatePermission: true,
		Scheduled:         true,
	}
	j, err := s.JSONString()
	if err != nil {
		assert.Nil(t, err)
	}
	assert.Equal(
		t, `{
  "name": "stage1",
  "id": 13,
  "jobs": [
    {
      "name": "job1",
      "scheduled_date": 1436172201081,
      "result": "Failed",
      "state": "Completed",
      "id": 13
    }
  ],
  "can_run": true,
  "scheduled": true,
  "approval_type": "success",
  "approved_by": "admin",
  "counter": "1",
  "operate_permission": true,
  "result": "Failed"
}`, j)
}

/*
expected: "{\n  \"name\": \"stage1\",\n  \"id\": 13,\n  \"jobs\": [\n    {\n      \"name\": \"job1\",\n      \"scheduled_date\": 1436172201081,\n      \"result\": \"Failed\",\n      \"state\": \"Completed\",\n      \"id\": 13\n    }\n  ],\n  \"can_run\": true,\n  \"scheduled\": true,\n  \"approval_type\": \"success\",\n  \"approved_by\": \"admin\",\n  \"counter\": 1,\n  \"operate_permission\": true,\n  \"result\": \"Failed\"\n}"
  actual: "{\n  \"name\": \"stage1\",\n  \"id\": 13,\n  \"jobs\": [\n    {\n      \"name\": \"job1\",\n      \"scheduled_date\": 1436172201081,\n      \"result\": \"Failed\",\n      \"state\": \"Completed\",\n      \"id\": 13\n    }\n  ],\n  \"can_run\": true,\n  \"scheduled\": true,\n  \"approval_type\": \"success\",\n  \"approval_by\": \"admin\",\n  \"counter\": 1,\n  \"operate_permission\": true,\n  \"result\": \"Failed\"\n}"

*/
// "{\n  \"name\": \"stage1\",\n  \"id\": 13,\n  \"jobs\": [\n    {\n      \"name\": \"job1\",\n      \"scheduled_date\": 1436172201081,\n      \"result\": \"Failed\",\n      \"state\": \"Completed\",\n      \"id\": 13\n    }\n  ],\n  \"can_run\": true,\n  \"approval_type\": \"success\",\n  \"approved_by\": \"admin\",\n  \"counter\": \"1\",\n  \"operate_permission\": true,\n  \"scheduled\": true\n  \"result\": \"Failed\",\n}"
// "{\n  \"name\": \"stage1\",\n  \"id\": 13,\n  \"jobs\": [\n    {\n      \"name\": \"job1\",\n      \"scheduled_date\": 1436172201081,\n      \"result\": \"Failed\",\n      \"state\": \"Completed\",\n      \"id\": 13\n    }\n  ],\n  \"can_run\": true,\n  \"scheduled\": true,\n  \"approval_type\": \"success\",\n  \"approval_by\": \"admin\",\n  \"counter\": 1,\n  \"operate_permission\": true,\n  \"result\": \"Failed\"\n}"

// {\n  \"name\": \"stage1\",\n  \"id\": 13,\n  \"jobs\": [\n    {\n      \"name\": \"job1\",\n      \"scheduled_date\": 1436172201081,\n      \"result\": \"Failed\",\n      \"state\": \"Completed\",\n      \"id\": 13\n    }\n  ],\n  \"can_run\": true,\n  \"approval_type\": \"success\",\n  \"approved_by\": \"admin\",\n  \"counter\": \"1\",\n  \"operate_permission\": true,\n  \"scheduled\": true,\n  \"result\": \"Failed\"\n}
// {\n  \"name\": \"stage1\",\n  \"id\": 13,\n  \"jobs\": [\n    {\n      \"name\": \"job1\",\n      \"scheduled_date\": 1436172201081,\n      \"result\": \"Failed\",\n      \"state\": \"Completed\",\n      \"id\": 13\n    }\n  ],\n  \"can_run\": true,\n  \"scheduled\": true,\n  \"approval_type\": \"success\",\n  \"approval_by\": \"admin\",\n  \"counter\": 1,\n  \"operate_permission\": true,\n  \"result\": \"Failed\"\n}"

func testStageInstanceValidate(t *testing.T) {
	s := Stage{}

	err := s.Validate()
	assert.EqualError(t, err, "`gocd.Stage.Name` is empty")

	s.Name = "test-stage"
	err = s.Validate()
	assert.EqualError(t, err, "At least one `Job` must be specified")

	s.Jobs = []*Job{{}}
	err = s.Validate()
	assert.EqualError(t, err, "`gocd.Jobs.Name` is empty")

	s.Jobs[0].Name = "test-job"
	err = s.Validate()
	assert.Nil(t, err)
}
