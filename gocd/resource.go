package gocd

// StageContainer describes structs which contain stages, eg Pipelines and PipelineTemplates
type StageContainer interface {
	GetName() string
	SetStage(stage *Stage)
	GetStage(string) *Stage
	SetStages(stages []*Stage)
	GetStages() []*Stage
	AddStage(stage *Stage)
}

// Versioned describes resources which can get and set versions
type Versioned interface {
	GetVersion() string
	SetVersion(version string)
}
