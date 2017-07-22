package gocd

type PipelineTemplatesService service

type PipelineTemplate struct {
	Name   *string `json:"name"`
	Stages []Stage `json:"stages"`
}
