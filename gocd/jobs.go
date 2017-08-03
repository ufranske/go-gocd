package gocd

import "errors"

type JobsService service

type Job struct {
	AgentUUID            string                `json:"agent_uuid,omitempty"`
	Name                 string                `json:"name"`
	JobStateTransitions  []*JobStateTransition `json:"job_state_transitions,omitempty"`
	ScheduledDate        int64                 `json:"scheduled_date,omitempty"`
	OrginalJobId         string                `json:"orginal_job_id,omitempty"`
	PipelineCounter      int64                 `json:"pipeline_counter,omitempty"`
	Rerun                bool                  `json:"rerun,omitempty"`
	PipelineName         string                `json:"pipeline_name,omitempty"`
	Result               string                `json:"result,omitempty"`
	State                string                `json:"state,omitempty"`
	Id                   int64                 `json:"id,omitempty"`
	StageCounter         string                `json:"stage_counter,omitempty"`
	StageName            string                `json:"stage_name,omitempty"`
	RunInstanceCount     int64                 `json:"run_instance_count,omitempty"`
	Timeout              int64                 `json:"timeout,omitempty"`
	EnvironmentVariables []string              `json:"environment_variables,omitempty"`
	Resources            []string              `json:"resources,omitempty"`
	Tasks                []Task                `json:"tasks,omitempty"`
}

type PluginConfiguration struct {
	Id      string `json:"id"`
	Version string `json:"version"`
}

type PluginConfigurationKVPair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Task struct {
	Type       string         `json:"type"`
	Attributes TaskAttributes `json:"attributes"`
}

func (t *Task) Validate() error {
	switch t.Type {
	case "exec":
		return t.Attributes.ValidateExec()
	case "ant":
		return t.Attributes.ValidateAnt()
	default:
		return errors.New("Unexpected `gocd.Task.Attribute` types")
	}
}

type TaskAttributes struct {
	RunIf               []string                    `json:"run_if,omitempty"`
	Command             string                      `json:"command,omitempty"`
	WorkingDirectory    string                      `json:"working_directory,omitempty"`
	Arguments           []string                    `json:"arguments,omitempty"`
	BuildFile           string                      `json:"build_file,omitempty"`
	Target              string                      `json:"target,omitempty"`
	NantPath            string                      `json:"nant_path,omitempty"`
	Pipeline            string                      `json:"pipeline,omitempty"`
	Stage               string                      `json:"stage,omitempty"`
	Job                 string                      `json:"job,omitempty"`
	Source              string                      `json:"source,omitempty"`
	IsSourceAFile       string                      `json:"is_source_a_file,omitempty"`
	Destination         string                      `json:"destination,omitempty"`
	PluginConfiguration *PluginConfiguration         `json:"plugin_configuration,omitempty"`
	Configuration       []PluginConfigurationKVPair `json:"configuration,omitempty"`
}

func (t *TaskAttributes) ValidateExec() error {
	if len(t.RunIf) == 0 {
		return errors.New("'run_if' must not be empty.")
	}
	if t.Command == "" {
		return errors.New("'command' must not be empty")
	}
	if len(t.Arguments) == 0 {
		return errors.New("'arguments' must not be empty.")
	}
	if t.WorkingDirectory == "" {
		return errors.New("'working_directory' must not empty.")
	}

	return nil
}

func (t *TaskAttributes) ValidateAnt() error {
	if len(t.RunIf) == 0 {
		return errors.New("'run_if' must not be empty.")
	}
	if t.BuildFile == "" {
		return errors.New("'build_file' must not be empty")
	}
	if t.Target == "" {
		return errors.New("'target' must not be empty")
	}
	if t.WorkingDirectory == "" {
		return errors.New("'working_directory' must not empty.")
	}

	return nil
}

type JobStateTransition struct {
	StateChangeTime int64  `json:"state_change_time,omitempty"`
	Id              int64  `json:"id,omitempty"`
	State           string `json:"state,omitempty"`
}

type JobRunHistoryResponse struct {
	Jobs       []*Job              `json:"jobs,omitempty"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
}

func (j *Job) Validate() error {
	if j.Name == "" {
		return errors.New("`gocd.Jobs.Name` is empty.")
	}
	return nil
}
