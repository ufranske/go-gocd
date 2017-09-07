package gocd

// GetStages from the pipeline
func (p *Pipeline) GetStages() []*Stage {
	return p.Stages
}

// GetStage from the pipeline template
func (p *Pipeline) GetStage(stageName string) *Stage {
	for _, stage := range p.Stages {
		if stage.Name == stageName {
			return stage
		}
	}
	return nil
}

// GetName of the pipeline
func (p *Pipeline) GetName() string {
	return p.Name
}

// SetStages overwrites any existing stages
func (p *Pipeline) SetStages(stages []*Stage) {
	p.Stages = stages
}

// AddStage appends a stage to this pipeline
func (p *Pipeline) AddStage(stage *Stage) {
	p.Stages = append(p.Stages, stage)
}
