package gocd

// PropertyRequest describes the parameters to be submitted when calling/creating properties.
type PropertyRequest struct {
	Pipeline        string
	PipelineCounter int
	Stage           string
	StageCounter    int
	Job             string
	LimitPipeline   string
	Limit           int
}
