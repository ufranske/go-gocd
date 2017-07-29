package main

import (
	"flag"
	"fmt"
	"strings"
	"os"
)

const command_template = `package main

import "github.com/urfave/cli"

const (
	%[1]sCommandName  = "%[2]s"
	%[1]sCommandUsage = "%[3]s"
)

func %[1]sAction() error {
	return nil
}

func %[1]sCommand() *cli.Command {
	return &cli.Command{
		Name:   %[1]sCommandName,
		Usage:  %[1]sCommandUsage,
		Action: %[1]sAction,
	}
}

`

func main() {
	var cn string
	flag.StringVar(&cn, "command", "", "Name of the command to generate")
	var dsc string
	flag.StringVar(&dsc, "description", "", "Description for the command")
	var stdout bool
	flag.BoolVar(&stdout, "stdout", false, "If true, print to stdout.")
	flag.Parse()

	if stdout {
		fmt.Printf(command_template, strings.Title(cn), strings.ToLower(cn), dsc)
	} else {
		f, err := os.Create(fmt.Sprintf("./%s.go", cn))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		_, err = f.WriteString(fmt.Sprintf(command_template, strings.Title(cn), strings.ToLower(cn), dsc))
		if err != nil {
			panic(err)
		}
	}
}
