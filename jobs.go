package gocd

import "errors"

type JobsService service

type Job struct {
	AgentUUID           string                `json:"agent_uuid,omitempty"`
	Name                string                `json:"name"`
	JobStateTransitions []*JobStateTransition `json:"job_state_transitions,omitempty"`
	ScheduledDate       int64                 `json:"scheduled_date,omitempty"`
	OrginalJobId        string                `json:"orginal_job_id,omitempty"`
	PipelineCounter     int64                 `json:"pipeline_counter,omitempty"`
	Rerun               bool                  `json:"rerun,omitempty"`
	PipelineName        string                `json:"pipeline_name,omitempty"`
	Result              string                `json:"result,omitempty"`
	State               string                `json:"state,omitempty"`
	Id                  int64                 `json:"id,omitempty"`
	StageCounter        string                `json:"stage_counter,omitempty"`
	StageName           string                `json:"stage_name,omitempty"`
}

type JobStateTransition struct {
	StateChangeTime int64  `json:"state_change_time"`
	Id              int64  `json:"id"`
	State           string `json:"state"`
}

type JobRunHistoryResponse struct {
	Jobs       []*Job              `json:"jobs"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
}

func (j *Job) Validate() error {
	if j.Name == "" {
		return errors.New("`gocd.Jobs.Name` is empty.")
	}
	return nil
}
