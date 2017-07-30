package gocd

import "errors"

type StagesService service

type Stage struct {
	Name                  string    `json:"name"`
	FetchMaterials        bool      `json:"fetch_materials"`
	CleanWorkingDirectory bool      `json:"clean_working_directory"`
	NeverCleanupArtifacts bool      `json:"never_cleanup_artifacts"`
	Approval              *Approval `json:"approval,omitempty"`
	EnvironmentVariables  []string  `json:"environment_variables,omitempty"`
	Jobs                  []*Job     `json:"jobs,omitempty"`
}

//func (s Stage) MarshalJSON() ([]byte, error) {
//	if s.Approval == nil {
//		return nil, nil
//	}
//	return nil, nil
//}

func (s *Stage) Validate() error {
	if s.Name == "" {
		return errors.New("`gocd.Stage.Name` is empty.")
	}

	if len(s.Jobs) == 0 {
		return errors.New("At least one `Job` must be spcified.")
	} else {
		for _, job := range s.Jobs {
			if err := job.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}

//type StageInstance struct {
//	Name                  *string `json:"name"`
//	CleanWorkingDirectory *bool `json:"clean_working_directory"`
//	NeverCleanupArtifacts *bool `json:"never_cleanup_artifacts"`
//	Approval              *Approval `json:"approval"`
//	ApprovedBy            *string `json:"approved_by"`
//	Jobs                  []Job `json:"jobs"`
//	PipelineCounter       *int `json:"pipeline_counter"`
//	PipelineName          *string `json:"pipeline_name"`
//	ApprovalType          *string `json:"approval_type"`
//	Result                *string `json:"result"`
//	Counter               *int `json:"counter"`
//	Id                    *int `json:"id"`
//	RerunOfCounter        *int `json:"rerun_of_counter"`
//	FetchMaterials        *bool `json:"fetch_materials"`
//	ArtifactsDeleted      *bool `json:"artifacts_deleted"`
//}
