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

// GetScheduledJobsAction gets a list of agents and return them.
func ListScheduledJobsAction(c *cli.Context) error {
	encryptedValue, r, err := cliAgent(c).Jobs.ListScheduled(context.Background())
	if err != nil {
		return handleOutput(nil, r, "ListScheduledJobs", err)
	}
	return handleOutput(encryptedValue, r, "ListScheduledJobs", err)
}

// EncryptCommand checks a template-name is provided and that the response is a 2xx response.
func ListScheduledJobsCommand() *cli.Command {
	return &cli.Command{
		Name:     ListScheduledJobsCommandName,
		Usage:    ListScheduledJobsCommandUsage,
		Action:   ListScheduledJobsAction,
		Category: "Jobs",
	}
}
