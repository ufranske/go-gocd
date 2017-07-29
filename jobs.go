package gocd

import "context"

type JobsService service

type Job struct {
	AgentUUID           string                `json:"agent_uuid"`
	Name                string                `json:"name"`
	JobStateTransitions []*JobStateTransition `json:"job_state_transitions"`
	ScheduledDate       int64                 `json:"scheduled_date"`
	OrginalJobId        string                `json:"orginal_job_id"`
	PipelineCounter     int64                 `json:"pipeline_counter"`
	Rerun               bool                  `json:"rerun"`
	PipelineName        string                `json:"pipeline_name"`
	Result              string                `json:"result"`
	State               string                `json:"state"`
	Id                  int64                 `json:"id"`
	StageCounter        string                `json:"stage_counter"`
	StageName           string                `json:"stage_name"`
}

type JobStateTransition struct {
	StateChangeTime int64  `json:"state_change_time"`
	Id              int64  `json:"id"`
	State           string `json:"state"`
}

type JobRunHistoryResponse struct {
	Jobs       []*Job             `json:"jobs"`
	Pagination PaginationResponse `json:"pagination,omitempty"`
}

func (a *Agent) JobRunHistory(ctx context.Context) ([]*Job, error) {
	a.client.Agents.JobRunHistory(ctx, a.Uuid)
	return nil, nil
}
