package cli

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/drewsonne/go-gocd/gocd"
	"github.com/urfave/cli"
)

// List of command name and descriptions
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

// ListAgentsAction gets a list of agents and return them.
func listAgentsAction(c *cli.Context) cli.ExitCoder {
	agents, r, err := cliAgent(c).Agents.List(context.Background())
	if err != nil {
		return NewCliError("ListAgents", r, err)
	}
	for _, agent := range agents {
		agent.RemoveLinks()
	}
	return handleOutput(agents, "ListAgents")
}

// GetAgentAction retrieves a single agent object.
func getAgentAction(c *cli.Context) cli.ExitCoder {
	agent, r, err := cliAgent(c).Agents.Get(context.Background(), c.String("uuid"))
	if r.HTTP.StatusCode != 404 {
		agent.RemoveLinks()
	}
	if err != nil {
		return NewCliError("GetAgent", r, err)
	}
	return handleOutput(agent, "GetAgent")
}

// UpdateAgentAction updates a single agent.
func updateAgentAction(c *cli.Context) cli.ExitCoder {

	if c.String("uuid") == "" {
		return NewCliError("UpdateAgent", nil, errors.New("'--uuid' is missing"))
	}

	if c.String("config") == "" {
		return NewCliError("UpdateAgent", nil, errors.New("'--config' is missing"))
	}

	a := &gocd.Agent{}
	b := []byte(c.String("config"))
	if err := json.Unmarshal(b, &a); err != nil {
		return NewCliError("UpdateAgent", nil, err)
	}

	agent, r, err := cliAgent(c).Agents.Update(context.Background(), c.String("uuid"), a)
	if r.HTTP.StatusCode != 404 {
		agent.RemoveLinks()
	} else if err != nil {
		return NewCliError("UpdateAgent", r, err)
	}
	return handleOutput(agent, "UpdateAgent")
}

// DeleteAgentAction delets an agent. Note: The agent must be disabled.
func deleteAgentAction(c *cli.Context) cli.ExitCoder {
	if c.String("uuid") == "" {
		return handleOutput(nil, "DeleteAgent")
	}

	deleteResponse, r, err := cliAgent(c).Agents.Delete(context.Background(), c.String("uuid"))
	if r.HTTP.StatusCode == 406 {
		err = errors.New(deleteResponse)
	}
	if err != nil {
		return NewCliError("DeleteAgent", r, err)
	}
	return handleOutput(deleteResponse, "DeleteAgent")
}

// UpdateAgentsAction updates a single agent.
func updateAgentsAction(c *cli.Context) cli.ExitCoder {

	u := gocd.AgentBulkUpdate{}
	if o := c.String("operations"); o != "" {
		b := []byte(o)
		op := gocd.AgentBulkOperationsUpdate{}
		if err := json.Unmarshal(b, &op); err == nil {
			return handleOutput(nil, "BulkAgentUpdate")
		}
		u.Operations = &op
	}

	var uuids []string
	if uuids = c.StringSlice("uuid"); len(uuids) == 0 {
		return handleOutput(nil, "BulkAgentUpdate")
	}
	u.Uuids = uuids

	if state := c.String("state"); state != "" {
		u.AgentConfigState = c.String("state")
	}

	updateResponse, r, err := cliAgent(c).Agents.BulkUpdate(context.Background(), u)
	if r.HTTP.StatusCode == 406 {
		err = errors.New(updateResponse)
	}
	if err != nil {
		return NewCliError("BulkAgentUpdate", r, err)
	}
	return handleOutput(updateResponse, "BulkAgentUpdate")
}

// DeleteAgentsAction must be implemented.
func deleteAgentsAction(c *cli.Context) cli.ExitCoder {
	return nil
}

// ListAgentsCommand checks a template-name is provided and that the response is a 2xx response.
func listAgentsCommand() *cli.Command {
	return &cli.Command{
		Name:     ListAgentsCommandName,
		Usage:    ListAgentsCommandUsage,
		Action:   listAgentsAction,
		Category: "Agents",
	}
}

// GetAgentCommand handles the interaction between the cli flags and the action handler for get-agent
func getAgentCommand() *cli.Command {
	return &cli.Command{
		Name:     GetAgentCommandName,
		Usage:    GetAgentCommandUsage,
		Action:   getAgentAction,
		Category: "Agents",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "uuid, u", Usage: "GoCD Agent UUID"},
		},
	}
}

// UpdateAgentCommand handles the interaction between the cli flags and the action handler for update-agent
func updateAgentCommand() *cli.Command {
	return &cli.Command{
		Name:     UpdateAgentCommandName,
		Usage:    UpdateAgentCommandUsage,
		Action:   updateAgentAction,
		Category: "Agents",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "uuid, u", Usage: "GoCD Agent UUID"},
			cli.StringFlag{Name: "config, c", Usage: "JSON encoded config for agent update."},
		},
	}
}

// DeleteAgentCommand handles the interaction between the cli flags and the action handler for delete-agent
func deleteAgentCommand() *cli.Command {
	return &cli.Command{
		Name:     DeleteAgentCommandName,
		Usage:    DeleteAgentCommandUsage,
		Action:   deleteAgentAction,
		Category: "Agents",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "uuid, u", Usage: "GoCD Agent UUID"},
		},
	}
}

// UpdateAgentsCommand handles the interaction between the cli flags and the action handler for update-agents
func updateAgentsCommand() *cli.Command {
	return &cli.Command{
		Name:     UpdateAgentsCommandName,
		Usage:    UpdateAgentsCommandUsage,
		Action:   updateAgentsAction,
		Category: "Agents",
		Flags: []cli.Flag{
			cli.StringSliceFlag{Name: "uuid", Usage: "GoCD Agent UUIDs"},
			cli.StringFlag{Name: "state", Usage: "Whether agents are enabled or disabled. Allowed values 'Enabled','Disabled'."},
			cli.StringFlag{Name: "operations", Usage: "JSON encoded config for bulk operation updates."},
		},
	}
}

// DeleteAgentsCommand handles the interaction between the cli flags and the action handler for delete-agents
func deleteAgentsCommand() *cli.Command {
	return &cli.Command{
		Name:     DeleteAgentsCommandName,
		Usage:    DeleteAgentsCommandUsage,
		Action:   deleteAgentsAction,
		Category: "Agents",
		Flags: []cli.Flag{
			cli.StringSliceFlag{Name: "uuid", Usage: "GoCD Agent UUIDs"},
		},
	}
}
