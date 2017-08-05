package gocd

import "errors"

// JobsService describes the HAL _link resource for the api response object for a job
type JobsService service

// Job describes a job object.
type Job struct {
	AgentUUID            string                 `json:"agent_uuid,omitempty"`
	Name                 string                 `json:"name"`
	JobStateTransitions  []*JobStateTransition  `json:"job_state_transitions,omitempty"`
	ScheduledDate        int64                  `json:"scheduled_date,omitempty"`
	OrginalJobID         string                 `json:"orginal_job_id,omitempty"`
	PipelineCounter      int64                  `json:"pipeline_counter,omitempty"`
	Rerun                bool                   `json:"rerun,omitempty"`
	PipelineName         string                 `json:"pipeline_name,omitempty"`
	Result               string                 `json:"result,omitempty"`
	State                string                 `json:"state,omitempty"`
	ID                   int64                  `json:"id,omitempty"`
	StageCounter         string                 `json:"stage_counter,omitempty"`
	StageName            string                 `json:"stage_name,omitempty"`
	RunInstanceCount     int64                  `json:"run_instance_count,omitempty"`
	Timeout              int64                  `json:"timeout,omitempty"`
	EnvironmentVariables []*EnvironmentVariable `json:"environment_variables,omitempty"`
	Properties           []*JobProperty         `json:"properties,omitempty"`
	Resources            []string               `json:"resources,omitempty"`
	Tasks                []Task                 `json:"tasks,omitempty"`
	Tabs                 []string               `json:"tabs,omitempty"`
	Artifacts            []string               `json:"artifacts,omitempty"`
}

// JobProperty describes the property for a job
type JobProperty struct {
	Name   string `json:"name"`
	Source string `json:"source"`
	XPath  string `json:"xpath"`
}

// EnvironmentVariable describes an environment variable key/pair.
type EnvironmentVariable struct {
	Name           string `json:"name"`
	Value          string `json:"value,omitempty"`
	EncryptedValue string `json:"encrypted_value,omitempty"`
	Secure         bool   `json:"secure"`
}

// PluginConfiguration describes how to reference a plugin.
type PluginConfiguration struct {
	ID      string `json:"id"`
	Version string `json:"version"`
}

// PluginConfigurationKVPair describes a key/value pair of plugin configurations.
type PluginConfigurationKVPair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Task Describes a Task object in the GoCD api.
type Task struct {
	Type       string         `json:"type"`
	Attributes TaskAttributes `json:"attributes"`
}

// Validate each of the possible task types.
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

// TaskAttributes describes all the properties for a Task.
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
	PluginConfiguration *PluginConfiguration        `json:"plugin_configuration,omitempty"`
	Configuration       []PluginConfigurationKVPair `json:"configuration,omitempty"`
}

// ValidateExec checks that the specified values for the Task struct are correct for a cli exec task
func (t *TaskAttributes) ValidateExec() error {
	if len(t.RunIf) == 0 {
		return errors.New("'run_if' must not be empty")
	}
	if t.Command == "" {
		return errors.New("'command' must not be empty")
	}
	if len(t.Arguments) == 0 {
		return errors.New("'arguments' must not be empty")
	}
	if t.WorkingDirectory == "" {
		return errors.New("'working_directory' must not empty")
	}

	return nil
}

// ValidateAnt checks that the specified values for the Task struct are correct for a an Ant task
func (t *TaskAttributes) ValidateAnt() error {
	if len(t.RunIf) == 0 {
		return errors.New("'run_if' must not be empty")
	}
	if t.BuildFile == "" {
		return errors.New("'build_file' must not be empty")
	}
	if t.Target == "" {
		return errors.New("'target' must not be empty")
	}
	if t.WorkingDirectory == "" {
		return errors.New("'working_directory' must not empty")
	}

	return nil
}

// JobStateTransition describes a State Transition object in a GoCD api response
type JobStateTransition struct {
	StateChangeTime int64  `json:"state_change_time,omitempty"`
	ID              int64  `json:"id,omitempty"`
	State           string `json:"state,omitempty"`
}

// JobRunHistoryResponse describes the api response from
type JobRunHistoryResponse struct {
	Jobs       []*Job              `json:"jobs,omitempty"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
}

// Validate a job structure has non-nil values on correct attributes
func (j *Job) Validate() error {
	if j.Name == "" {
		return errors.New("`gocd.Jobs.Name` is empty")
	}
	return nil
}
