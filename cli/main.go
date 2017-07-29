// GoCD CLI tool

package main

import (
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
		*AgentsListCommand(),
	}
	app.Run(os.Args)
}
