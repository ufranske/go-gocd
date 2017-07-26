package gocd

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestPipelineTemplateService_ListPipelineTemplates(t *testing.T) {

	setup()
	defer teardown()

	mux.HandleFunc("/api/admin/templates", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testAuth(t, r, mockAuthorization)
		fmt.Fprint(w, `{
  "_links": {
    "self": {
      "href": "https://ci.example.com/go/api/admin/templates"
    },
    "doc": {
      "href": "https://api.gocd.org/#template-config"
    },
    "find": {
      "href": "https://ci.example.com/go/api/admin/templates/:template_name"
    }
  },
  "_embedded": {
    "templates": [
      {
        "_links": {
          "self": {
            "href": "https://ci.example.com/go/api/admin/templates/template1"
          },
          "doc": {
            "href": "https://api.gocd.org/#template-config"
          },
          "find": {
            "href": "https://ci.example.com/go/api/admin/templates/:template_name"
          }
        },
        "name": "template1",
        "_embedded": {
          "pipelines": [
            {
              "_links": {
                "self": {
                  "href": "https://ci.example.com/go/api/admin/pipelines/up42"
                },
                "doc": {
                  "href": "https://api.gocd.org/#pipeline-config"
                },
                "find": {
                  "href": "https://ci.example.com/go/api/admin/pipelines/:pipeline_name"
                }
              },
              "name": "up42"
            }
          ]
        }
      }
    ]
  }
}
`)
	})

	templates, _, err := client.PipelineTemplates.ListPipelineTemplates(context.Background())

	if err != nil {
		t.Error(err)
	}

	if len((*templates)) != 1 {
		t.Error(
			"Expected template length: '1 ",
			"Got: '", (*templates)[0].Name, "'",
		)
	}

	if (*templates)[0].Name != "template1" {
		t.Error(
			"Expected: 'template1. ",
			"Got: '", (*templates)[0].Name, "'",
		)
	}

	if (*templates)[0].Pipelines[0] != "up42" {
		t.Error(
			"Expected: 'up42. ",
			"Got: '", (*templates)[0].Pipelines[0], "'",
		)
	}
}

func TestPipelineTemplateService_GetPipelineTemplates(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/admin/templates/template1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testAuth(t, r, mockAuthorization)
		fmt.Fprint(w, `{
  "_links": {
    "self": {
      "href": "https://ci.example.com/go/api/admin/templates/template.name"
    },
    "doc": {
      "href": "https://api.gocd.org/#template-config"
    },
    "find": {
      "href": "https://ci.example.com/go/api/admin/templates/:template_name"
    }
  },
  "name": "template1",
  "stages": [
    {
      "name": "up42_stage",
      "fetch_materials": true,
      "clean_working_directory": false,
      "never_cleanup_artifacts": false,
      "approval": {
        "type": "success",
        "authorization": {
          "roles": [

          ],
          "users": [

          ]
        }
      },
      "environment_variables": [

      ],
      "jobs": [
        {
          "name": "up42_job",
          "run_instance_count": null,
          "timeout": "never",
          "elastic_profile_id": "docker",
          "environment_variables": [

          ],
          "resources": [

          ],
          "tasks": [
            {
              "type": "exec",
              "attributes": {
                "run_if": [

                ],
                "on_cancel": null,
                "command": "ls",
                "working_directory": null
              }
            }
          ],
          "tabs": [

          ],
          "artifacts": [

          ],
          "properties": null
        }
      ]
    }
  ]
}`)

	})
	template, _, err := client.PipelineTemplates.GetPipelineTemplate(
		context.Background(),
		"template1",
	)

	if err != nil {
		t.Error(err)
	}

	if template.Name != "template1" {
		t.Error(fmt.Printf(
			"Expected template name: 'template1'. Got: '%s'.", template.Name,
		))
	}
}
