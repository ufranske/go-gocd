package gocd

type PipelineGroupsService service

type PipelineGroup struct {
	Name   string `json:"name"`
	Stages []*Stage `json:"stages"`
}

type PipelineGroups struct {
	Collection []*PipelineGroup
}
