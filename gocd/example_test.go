package gocd_test

import (
	"context"
	"fmt"
	"github.com/beamly/go-gocd/gocd"
)

func ExampleAgentsService_List() {
	cfg := gocd.Configuration{
		Server:   "https://my_gocd/go/", // don't forget the "/go/" at the end of the url to avoid issues!
		Username: "ApiUser",
		Password: "MySecretPassword",
	}

	c := cfg.Client()

	// list all agents in use by the GoCD Server
	var a []*gocd.Agent
	var err error
	var r *gocd.APIResponse
	if a, r, err = c.Agents.List(context.Background()); err != nil {
		if r.HTTP.StatusCode == 404 {
			fmt.Println("Couldn't find agent")
		} else {
			panic(err)
		}
	}

	fmt.Println(a)
}

func ExampleConfigRepoService_List() {
	cfg := gocd.Configuration{
		Server:   "https://my_gocd/go/", // don't forget the "/go/" at the end of the url to avoid issues!
		Username: "ApiUser",
		Password: "MySecretPassword",
	}

	c := cfg.Client()

	l, _, err := c.ConfigRepos.List(context.Background())
	if err != nil {
		panic(err)
	}
	// Loops through the list of repositories to display some basic informations
	for _, r := range l {
		fmt.Printf("Pipeline: %s\n\tMaterial type: %s\n", r.ID, r.Material.Type)
		if r.Material.Type == "git" {
			fmt.Printf("\tMaterial url: %s\n", r.Material.Attributes.(*gocd.MaterialAttributesGit).URL)
		}
		fmt.Printf("\tNumber of configuration parameters: %d\n\n", len(r.Configuration))
	}
}

func ExampleConfigRepoService_Get() {
	cfg := gocd.Configuration{
		Server:   "https://my_gocd/go/", // don't forget the "/go/" at the end of the url to avoid issues!
		Username: "ApiUser",
		Password: "MySecretPassword",
	}

	c := cfg.Client()

	r, _, err := c.ConfigRepos.Get(context.Background(), "my_repo_config_id")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Pipeline: %s\n\tMaterial type: %s\n", r.ID, r.Material.Type)
	if r.Material.Type == "git" {
		fmt.Printf("\tMaterial url: %s\n", r.Material.Attributes.(*gocd.MaterialAttributesGit).URL)
	}
	fmt.Printf("\tNumber of configuration parameters: %d\n\n", len(r.Configuration))
}
