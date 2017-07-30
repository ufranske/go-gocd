package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/drewsonne/gocdsdk/gocd"
	"github.com/urfave/cli"
)

const (
	ListAgentsCommandName    = "list-agents"
	ListAgentsCommandUsage   = "List GoCD build agents."
	GetAgentCommandName      = "get-agent"
	GetAgentCommandUsage     = "Get Agent by UUID"
	UpdateAgentCommandName   = "update-agent"
	UpdateAgentCommandUsage  = "Update Agent"
	DeleteAgentCommandName   = "delete-agent"
	DeleteAgentCommandUsage  = "Delete Agent"
	UpdateAgentsCommandName  = "update-agents"
	UpdateAgentsCommandUsage = "Bulk Update Agents"
	DeleteAgentsCommandName  = "delete-agents"
	DeleteAgentsCommandUsage = "Bulk Delete Agents"
)

func ListAgentsAction(c *cli.Context) error {
	agents, r, err := cliAgent().Agents.List(context.Background())
	if err != nil {
		return handleOutput(nil, r, "ListAgents", err)
	}
	for _, agent := range agents {
		agent.RemoveLinks()
	}
	return handleOutput(agents, r, "ListAgents", err)
}

func GetAgentAction(c *cli.Context) error {
	agent, r, err := cliAgent().Agents.Get(context.Background(), c.String("uuid"))
	if r.Http.StatusCode != 404 {
		agent.RemoveLinks()
	}
	return handleOutput(agent, r, "GetAgent", err)
}

func UpdateAgentAction(c *cli.Context) error {

	if c.String("uuid") == "" {
		return handleOutput(nil, nil, "UpdateAgent", errors.New("'--uuid' is missing."))
	}

	if c.String("config") == "" {
		return handleOutput(nil, nil, "UpdateAgent", errors.New("'--config' is missing."))
	}

	a := gocd.AgentUpdate{}
	b := []byte(c.String("config"))
	if err := json.Unmarshal(b, &a); err != nil {
		return handleOutput(nil, nil, "UpdateAgent", err)
	}

	agent, r, err := cliAgent().Agents.Update(context.Background(), c.String("uuid"), a)
	if r.Http.StatusCode != 404 {
		agent.RemoveLinks()
	}
	return handleOutput(agent, r, "UpdateAgent", err)
}

func DeleteAgentAction(c *cli.Context) error {
	if c.String("uuid") == "" {
		return handleOutput(nil, nil, "DeleteAgent", errors.New("'--uuid' is missing."))
	}

	deleteResponse, r, err := cliAgent().Agents.Delete(context.Background(), c.String("uuid"))
	if r.Http.StatusCode == 406 {
		err = errors.New(deleteResponse)
	}
	return handleOutput(deleteResponse, r, "DeleteAgent", err)
}

func UpdateAgentsAction(c *cli.Context) error {

	u := gocd.AgentBulkUpdate{}
	if o := c.String("operations"); o != "" {
		b := []byte(o)
		op := gocd.AgentBulkOperationsUpdate{}
		if err := json.Unmarshal(b, &op); err == nil {
			return handleOutput(nil, nil, "BulkAgentUpdate", err)
		} else {
			u.Operations = &op
		}
	}

	if uuids := c.StringSlice("uuid"); len(uuids) == 0 {
		return handleOutput(nil, nil, "BulkAgentUpdate", errors.New("'--uuid' is missing."))
	} else {
		u.Uuids = uuids
	}

	if state := c.String("state"); state != "" {
		u.AgentConfigState = c.String("state")
	}

	updateResponse, r, err := cliAgent().Agents.BulkUpdate(context.Background(), u)
	if r.Http.StatusCode == 406 {
		err = errors.New(updateResponse)
	}
	return handleOutput(updateResponse, r, "BulkAgentUpdate", err)
}

func DeleteAgentsAction(c *cli.Context) error {
	return nil
}

func ListAgentsCommand() *cli.Command {
	return &cli.Command{
		Name:     ListAgentsCommandName,
		Usage:    ListAgentsCommandUsage,
		Action:   ListAgentsAction,
		Category: "Agents",
	}
}

func GetAgentCommand() *cli.Command {
	return &cli.Command{
		Name:     GetAgentCommandName,
		Usage:    GetAgentCommandUsage,
		Action:   GetAgentAction,
		Category: "Agents",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "uuid, u", Usage: "GoCD Agent UUID"},
		},
	}
}

func UpdateAgentCommand() *cli.Command {
	return &cli.Command{
		Name:     UpdateAgentCommandName,
		Usage:    UpdateAgentCommandUsage,
		Action:   UpdateAgentAction,
		Category: "Agents",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "uuid, u", Usage: "GoCD Agent UUID"},
			cli.StringFlag{Name: "config, c", Usage: "JSON encoded config for agent update."},
		},
	}
}

func DeleteAgentCommand() *cli.Command {
	return &cli.Command{
		Name:     DeleteAgentCommandName,
		Usage:    DeleteAgentCommandUsage,
		Action:   DeleteAgentAction,
		Category: "Agents",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "uuid, u", Usage: "GoCD Agent UUID"},
		},
	}
}

func UpdateAgentsCommand() *cli.Command {
	return &cli.Command{
		Name:     UpdateAgentsCommandName,
		Usage:    UpdateAgentsCommandUsage,
		Action:   UpdateAgentsAction,
		Category: "Agents",
		Flags: []cli.Flag{
			cli.StringSliceFlag{Name: "uuid", Usage: "GoCD Agent UUIDs"},
			cli.StringFlag{Name: "state", Usage: "Whether agents are enabled or disabled. Allowed values 'Enabled','Disabled'."},
			cli.StringFlag{Name: "operations", Usage: "JSON encoded config for bulk operation updates."},
		},
	}
}

func DeleteAgentsCommand() *cli.Command {
	return &cli.Command{
		Name:     DeleteAgentsCommandName,
		Usage:    DeleteAgentsCommandUsage,
		Action:   DeleteAgentsAction,
		Category: "Agents",
		Flags: []cli.Flag{
			cli.StringSliceFlag{Name: "uuid", Usage: "GoCD Agent UUIDs"},
		},
	}
}
