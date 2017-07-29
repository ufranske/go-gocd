package main

import (
	"github.com/urfave/cli"
	"context"
	"github.com/drewsonne/gocdsdk"
	"encoding/json"
	"errors"
)

const (
	ListAgentsCommandName   = "list-agents"
	ListAgentsCommandUsage  = "List GoCD build agents."
	GetAgentCommandName     = "get-agent"
	GetAgentCommandUsage    = "Get Agent by UUID"
	UpdateAgentCommandName  = "update-agent"
	UpdateAgentCommandUsage = "Update Agent"
)

func ListAgentsAction(c *cli.Context) error {
	agents, r, err := cliAgent().Agents.List(context.Background())
	for _, agent := range agents {
		agent.RemoveLinks()
	}
	return handleOutput(agents, r, "ListAgents", err)
}

func GetAgentAction(c *cli.Context) error {
	agent, r, err := cliAgent().Agents.Get(context.Background(), c.String("uuid"))
	if r.StatusCode != 404 {
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
	err := json.Unmarshal(b, &a)
	if err != nil {
		return handleOutput(nil, nil, "UpdateAgent", err)
	}

	agent, r, err := cliAgent().Agents.Update(context.Background(), c.String("uuid"), a)
	if r.StatusCode != 404 {
		agent.RemoveLinks()
	}
	return handleOutput(agent, r, "UpdateAgent", err)
}

func ListAgentsCommand() *cli.Command {
	return &cli.Command{
		Name:   ListAgentsCommandName,
		Usage:  ListAgentsCommandUsage,
		Action: ListAgentsAction,
	}
}

func GetAgentCommand() *cli.Command {
	return &cli.Command{
		Name:   GetAgentCommandName,
		Usage:  GetAgentCommandUsage,
		Action: GetAgentAction,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "uuid, u", Usage: "GoCD Agent UUID"},
		},
	}
}

func UpdateAgentCommand() *cli.Command {
	return &cli.Command{
		Name:   UpdateAgentCommandName,
		Usage:  UpdateAgentCommandUsage,
		Action: UpdateAgentAction,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "uuid, u", Usage: "GoCD Agent UUID"},
			cli.StringFlag{Name: "config, c", Usage: "JSON encoded config for agent update."},
		},
	}
}
