package gocd

type StagesService service

type Stage struct {
	Name                  *string `json:"name"`
	CleanWorkingDirectory *bool `json:"clean_working_directory"`
	ApprovedBy            *string `json:"approved_by"`
	Jobs                  []Job `json:"jobs"`
	PipelineCounter       *int `json:"pipeline_counter"`
	PipelineName          *string `json:"pipeline_name"`
	ApprovalType          *string `json:"approval_type"`
	Result                *string `json:"result"`
	Counter               *int `json:"counter"`
	Id                    *int `json:"id"`
	RerunOfCounter        *int `json:"rerun_of_counter"`
	FetchMaterials        *bool `json:"fetch_materials"`
	ArtifactsDeleted      *bool `json:"artifacts_deleted"`
}
