package gocd

type JobsService service

type Job struct {
	//AgentUUID *uuid.UUID `json:"agent_uuid"`
	Name string `json:"name"`
}
