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
	Jobs                  []*Job    `json:"jobs,omitempty"`
}

func (s *Stage) Validate() error {
	if s.Name == "" {
		return errors.New("`gocd.Stage.Name` is empty.")
	}

	if len(s.Jobs) == 0 {
		return errors.New("At least one `Job` must be specified.")
	} else {
		for _, job := range s.Jobs {
			if err := job.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}
