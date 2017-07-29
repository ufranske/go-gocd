package main

import (
	"github.com/urfave/cli"
	"os/user"
	"strings"
)

const (
	ConfigureCommandName  = "configure"
	ConfigureCommandUsage = "Generate configuration file ~/.gocd.conf"
	ConfigDirectoryPath   = "~/.gocd.conf"
)

func ConfigureAction() error {
	return nil
}

func configFilePath() string {
	// @TODO Make it work for windows. Maybe...
	usr, _ := user.Current()
	return strings.Replace(ConfigDirectoryPath, "~", usr.HomeDir, 1)
}

func ConfigureCommand() *cli.Command {
	return &cli.Command{
		Name:   ConfigureCommandName,
		Usage:  ConfigureCommandUsage,
		Action: ConfigureAction,
	}
}
