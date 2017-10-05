package cli

import (
	"encoding/json"
	"fmt"
	"github.com/drewsonne/go-gocd/gocd"
	"github.com/urfave/cli"
)

// GetCliCommands returns a list of all CLI Command structs
func GetCliCommands() []cli.Command {
	return []cli.Command{
		*configureCommand(),
		*listAgentsCommand(),
		*listPipelineTemplatesCommand(),
		*getAgentCommand(),
		*getPipelineTemplateCommand(),
		*createPipelineTemplateCommand(),
		*updateAgentCommand(),
		*updateAgentsCommand(),
		*updatePipelineConfigCommand(),
		*updatePipelineTemplateCommand(),
		*deleteAgentCommand(),
		*deleteAgentsCommand(),
		*deletePipelineTemplateCommand(),
		*deletePipelineConfigCommand(),
		*listPipelineGroupsCommand(),
		*getPipelineHistoryCommand(),
		*getPipelineCommand(),
		*createPipelineConfigCommand(),
		*generateJSONSchemaCommand(),
		*getPipelineStatusCommand(),
		*pausePipelineCommand(),
		*unpausePipelineCommand(),
		*releasePipelineLockCommand(),
		*getConfigurationCommand(),
		*encryptCommand(),
		*getVersionCommand(),
		*listPluginsCommand(),
		*getPluginCommand(),
		*listScheduledJobsCommand(),
		*getPipelineConfigCommand(),
		*listEnvironmentsCommand(),
		*getEnvironmentCommand(),
		*addPipelinesToEnvironmentCommand(),
		*removePipelinesFromEnvironmentCommand(),
		*listPropertiesCommand(),
		*createPropertyCommand(),
	}
}

// NewCliClient
func cliAgent(c *cli.Context) *gocd.Client {
	var profile string
	if profile = c.Parent().String("profile"); profile == "" {
		profile = "default"
	}

	cfg, err := gocd.LoadConfigByName(profile)
	if err != nil {
		panic(err)
	}

	if server := c.String("server"); server != "" {
		cfg.Server = server
	}

	if username := c.String("username"); username != "" {
		cfg.Username = username
	}

	if password := c.String("password"); password != "" {
		cfg.Password = password
	}

	cfg.SkipSslCheck = cfg.SkipSslCheck || c.Bool("skip_ssl_check")

	return cfg.Client()
}

func handleOutput(r interface{}, reqType string) cli.ExitCoder {
	o := map[string]interface{}{
		fmt.Sprintf("%sResponse", reqType): r,
	}
	b, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
	return nil
}
