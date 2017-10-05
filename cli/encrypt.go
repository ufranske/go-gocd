package cli

import (
	"context"
	"errors"
	"github.com/urfave/cli"
)

// List of command name and descriptions
const (
	EncryptCommandName  = "encrypt"
	EncryptCommandUsage = "Encrypt a value for use in GoCD configurations"
)

// EncryptAction gets a list of agents and return them.
func encryptAction(c *cli.Context) cli.ExitCoder {
	var value string
	if value = c.String("value"); value == "" {
		return NewCliError("Encrypt", nil, errors.New("'--value' is missing"))
	}

	encryptedValue, r, err := cliAgent(c).Encryption.Encrypt(context.Background(), value)
	if err != nil {
		return NewCliError("Encrypt", r, err)
	}
	return handleOutput(encryptedValue, "Encrypt")
}

// EncryptCommand checks a template-name is provided and that the response is a 2xx response.
func encryptCommand() *cli.Command {
	return &cli.Command{
		Name:     EncryptCommandName,
		Usage:    EncryptCommandUsage,
		Action:   encryptAction,
		Category: "Encryption",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "value"},
		},
	}
}
