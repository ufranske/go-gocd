package main

import (
	"github.com/urfave/cli"
	"context"
)

const (
	CommandName  = ""
	CommandUsage = ""
)

func Action(c *cli.Context) error {
	return nil
}

func Command() *cli.Command {
	return &cli.Command{
		Name:   CommandName,
		Usage:  CommandUsage,
		Action: Action,
	}
}

