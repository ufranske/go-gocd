package cli

import (
	"github.com/urfave/cli"
	"github.com/alecthomas/jsonschema"
	"fmt"
	"github.com/drewsonne/go-gocd/gocd"
)

const (
	GenerateJSONSchemaCommandName  = "generate-json-schema"
	GenerateJSONSchemaCommandUsage = "Generates a JSON schema based on the structs in this library"
)

func GenerateJSONSchemaAction(c *cli.Context) error {
	fmt.Println(jsonschema.Reflect(&gocd.TestUser{}))

	return nil
}

func GenerateJSONSchemaCommand() *cli.Command {
	return &cli.Command{
		Name:   GenerateJSONSchemaCommandName,
		Usage:  GenerateJSONSchemaCommandUsage,
		Category: "Schema",
		Action: GenerateJSONSchemaAction,
	}
}