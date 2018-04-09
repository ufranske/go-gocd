package gocd_test

import (
	"context"
	"fmt"
	"github.com/beamly/go-gocd/gocd"
)

func ExampleAgent_List() {
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
