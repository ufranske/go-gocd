package gocd

// GetGroupByPipelineName filters pipeline groups by their contained pipelines
func (pg *PipelineGroups) GetGroupByPipelineName(pipelineName string) *PipelineGroup {
	for _, pipelineGroup := range *pg {
		for _, pipeline := range pipelineGroup.Pipelines {
			if pipeline.Name == pipelineName {
				return pipelineGroup
			}
		}
	}
	return nil
}
