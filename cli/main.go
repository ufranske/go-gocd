// GoCD CLI tool

package main

import (
	"encoding/json"
	"fmt"
	"github.com/drewsonne/gocdsdk"
	"github.com/urfave/cli"
	"os"
)

const UtilityName = "gocd"
const UtilityUsageInstructions = "CLI Tool to interact with GoCD server"
const UtilityVersion = "1.0.0"

func main() {
	app := cli.NewApp()
	app.Name = UtilityName
	app.Usage = UtilityUsageInstructions
	app.Version = UtilityVersion
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		*ConfigureCommand(),
		*ListAgentsCommand(),
		*GetAgentCommand(),
		*UpdateAgentCommand(),
		*DeleteAgentCommand(),
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "server", EnvVar: EnvVarServer},
		cli.StringFlag{Name: "username", EnvVar: EnvVarUsername},
		cli.StringFlag{Name: "password", EnvVar: EnvVarPassword},
	}

	app.Run(os.Args)
}

func cliAgent() *gocd.Client {
	cfg, err := loadConfig()
	if err != nil {
		panic(err)
	}

	var auth *gocd.Auth
	if cfg.HasAuth() {
		auth = &gocd.Auth{
			Username: cfg.Username,
			Password: cfg.Password,
		}
	}

	return gocd.NewClient(cfg.Server, auth, nil, cfg.SslCheck)

}

func handleOutput(r interface{}, ar *gocd.APIResponse, reqType string, err error) error {
	var b []byte
	var o map[string]interface{}
	if err != nil {
		o = map[string]interface{}{
			"Error": err.Error(),
		}
	} else if ar.StatusCode >= 200 && ar.StatusCode < 300 {
		o = map[string]interface{}{
			fmt.Sprintf("%sResponse", reqType): r,
		}
	} else if ar.StatusCode == 404 {
		o = map[string]interface{}{
			"Error": fmt.Sprintf("Could not find resource for '%s' action.", reqType),
		}
	} else {
		o = map[string]interface{}{
			"Error":        "An error occured while retrieving the resource.",
			"RequestError": fmt.Sprintf("%s", ar),
		}
	}
	b, err = json.MarshalIndent(o, "", "    ")
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil

}
