package gocd

import (
	"github.com/google/uuid"
)

type JobsService service

type Job struct {
	AgentUUID *uuid.UUID `json:"agent_uuid"`
	Name      *string `json:"name"`
}
