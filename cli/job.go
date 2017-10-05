package cli

import (
	"context"
	"github.com/urfave/cli"
)

// List of command name and descriptions
const (
	ListScheduledJobsCommandName  = "list-scheduled-jobs"
	ListScheduledJobsCommandUsage = "List Scheduled Jobs"
)

// ListScheduledJobsAction gets a list of agents and return them.
func listScheduledJobsAction(c *cli.Context) cli.ExitCoder {
	encryptedValue, r, err := cliAgent(c).Jobs.ListScheduled(context.Background())
	if err != nil {
		return NewCliError("ListScheduledJobs", r, err)
	}
	return handleOutput(encryptedValue, "ListScheduledJobs")
}

// ListScheduledJobsCommand provides interface between handler and action
func listScheduledJobsCommand() *cli.Command {
	return &cli.Command{
		Name:     ListScheduledJobsCommandName,
		Usage:    ListScheduledJobsCommandUsage,
		Action:   listScheduledJobsAction,
		Category: "Jobs",
	}
}
