package main

import (
	"github.com/urfave/cli"
	"context"
	"errors"
)

const (
	ListPipelineTemplatesCommandName  = "list-pipeline-templates"
	ListPipelineTemplatesCommandUsage = "List Pipeline Templates"
	GetPipelineTemplateCommandName    = "get-pipeline-template"
	GetPipelineTemplateCommandUsage   = "Get Pipeline Templates"
)

func ListPipelineTemplatesAction(c *cli.Context) error {
	ts, r, err := cliAgent().PipelineTemplates.List(context.Background())
	if err != nil {
		return handleOutput(nil, r, "ListPipelineTemplates", err)
	}

	type p struct {
		Name string `json:"name"`
	}

	type ptr struct {
		Name      string `json:"name"`
		Pipelines []*p `json:"pipelines"`
	}
	responses := []ptr{}
	for _, pt := range ts {
		pt.RemoveLinks()
		ps := []*p{}
		for _, pipe := range pt.Pipelines() {
			ps = append(ps, &p{pipe.Name})
		}

		responses = append(responses, ptr{
			Name:      pt.Name,
			Pipelines: ps,
		})
	}
	return handleOutput(responses, r, "ListPipelineTemplates", err)
}

func GetPipelineTemplateAction(c *cli.Context) error {
	if c.String("template-name") == "" {
		return handleOutput(nil, nil, "GetPipelineTemplate", errors.New("'--template-name' is missing."))
	}

	pt, r, err := cliAgent().PipelineTemplates.Get(context.Background(), c.String("template-name"))
	if r.StatusCode != 404 {
		pt.RemoveLinks()
	}
	return handleOutput(pt, r, "GetPipelineTemplate", err)
}

func ListPipelineTemplatesCommand() *cli.Command {
	return &cli.Command{
		Name:   ListPipelineTemplatesCommandName,
		Usage:  ListPipelineTemplatesCommandUsage,
		Action: ListPipelineTemplatesAction,
	}
}

func GetPipelineTemplateCommand() *cli.Command {
	return &cli.Command{
		Name:   GetPipelineTemplateCommandName,
		Usage:  GetPipelineTemplateCommandUsage,
		Action: GetPipelineTemplateAction,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "template-name", Usage: "Name of the Pipeline Template configuration."},
		},
	}
}
