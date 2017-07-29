package main

import (
	"github.com/urfave/cli"
)

func AgentsListCommand() *cli.Command {
	return &cli.Command{
		Name:   "configure",
		Usage:  "List GoCD build agents.",
		Action: AgentsListRun,
	}
}

func AgentsListRun() error {
	return nil
}
