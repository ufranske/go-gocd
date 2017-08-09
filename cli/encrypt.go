package cli

import (
	"github.com/urfave/cli"
	"context"
)

// List of command name and descriptions
const (
	EncryptCommandName  = "encrypt"
	EncryptCommandUsage = "Encrypt a value for use in GoCD configurations"
)

// EncryptAction gets a list of agents and return them.
func EncryptAction(c *cli.Context) error {
	encryptedValue, r, err := cliAgent(c).Encryption.Encrypt(context.Background())
	if err != nil {
		return handleOutput(nil, r, "Encrypt", err)
	}
	return handleOutput(encryptedValue, r, "Encrypt", err)
}

// EncryptCommand checks a template-name is provided and that the response is a 2xx response.
func EncryptCommand() *cli.Command {
	return &cli.Command{
		Name:     EncryptCommandName,
		Usage:    EncryptCommandUsage,
		Action:   EncryptAction,
		Category: "Encryption",
	}
}
