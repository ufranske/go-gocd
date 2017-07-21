package gocd

type PipelineGroupsService service

type PipelineGroup struct {
    Name    *string `json:"name"`
}

type PipelineGroups struct {
    Groups []PipelineGroup
}