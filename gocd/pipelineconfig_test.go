package gocd

import (
	"context"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestPipelineConfig(t *testing.T) {
	for _, tt := range []struct {
		name               string
		pipelineCreate     *Pipeline
		pipelineCreateWant *Pipeline
		pipelineGet        string
		pipelineGetWant    *Pipeline
		pipelineUpdate     *Pipeline
		pipelineUpdateWant *Pipeline
		delete             string
		deleteWant         string
	}{
		{
			name: "basic",
			pipelineCreate: &Pipeline{
				Name: "new_pipeline",
				Materials: []Material{{
					Type: "git",
					Attributes: MaterialAttributesGit{
						URL:         "git@github.com:sample_repo/example.git",
						Destination: "dest",
						Branch:      "master",
					},
				}},
				Stages: buildMockPipelineStages(),
			},
			pipelineCreateWant: &Pipeline{
				Group:                "test-group",
				Name:                 "new_pipeline",
				LabelTemplate:        "${COUNT}",
				Parameters:           make([]*Parameter, 0),
				EnvironmentVariables: make([]*EnvironmentVariable, 0),
				Materials: []Material{{
					Type: "git",
					Attributes: &MaterialAttributesGit{
						URL:         "git@github.com:sample_repo/example.git",
						Destination: "dest",
						Branch:      "master",
						AutoUpdate:  true,
					},
				}},
				Stages: buildMockPipelineStages(),
			},
			pipelineGet: "new_pipeline",
			pipelineGetWant: &Pipeline{
				Name:                 "new_pipeline",
				LabelTemplate:        "${COUNT}",
				Parameters:           make([]*Parameter, 0),
				EnvironmentVariables: make([]*EnvironmentVariable, 0),
				Materials: []Material{{
					Type: "git",
					Attributes: &MaterialAttributesGit{
						URL:         "git@github.com:sample_repo/example.git",
						Destination: "dest",
						Branch:      "master",
						AutoUpdate:  true,
					},
				}},
				Stages: []*Stage{{
					Name: "defaultStage",
					Approval: &Approval{
						Type: "success",
						Authorization: &Authorization{
							Users: []string{},
							Roles: []string{},
						},
					},
					Jobs: []*Job{{
						Name:                 "defaultJob",
						EnvironmentVariables: []*EnvironmentVariable{},
						Resources:            []string{},
						Tasks: []*Task{{
							Type: "exec",
							Attributes: TaskAttributes{
								RunIf:   []string{"passed"},
								Command: "ls",
							},
						}},
						Tabs:      []*Tab{},
						Artifacts: []*Artifact{},
					}},
					EnvironmentVariables: []*EnvironmentVariable{},
				}},
			},
			pipelineUpdateWant: &Pipeline{
				Group:                "test-group",
				Name:                 "new_pipeline",
				LabelTemplate:        "Update ${COUNT}",
				Parameters:           make([]*Parameter, 0),
				EnvironmentVariables: make([]*EnvironmentVariable, 0),
				Materials: []Material{{
					Type: "git",
					Attributes: &MaterialAttributesGit{
						URL:         "git@github.com:sample_repo/example.git",
						Destination: "dest",
						Branch:      "master",
						AutoUpdate:  true,
					},
				}},
				Stages: buildMockPipelineStages(),
			},
			delete:     "new_pipeline",
			deleteWant: "The pipeline 'new_pipeline' was deleted successfully.",
		},
	} {
		t.Run("basic", func(t *testing.T) {
			if runIntegrationTest(t) {

				ctx := context.Background()

				var getPipeline *Pipeline
				t.Run("create", func(t *testing.T) {
					createPipeline, _, err := intClient.PipelineConfigs.Create(ctx, "test-group", tt.pipelineCreate)
					assert.NoError(t, err)
					assert.Regexp(t, regexp.MustCompile("^[a-f0-9]{32}--gzip$"), createPipeline.Version)

					createPipeline.RemoveLinks()
					assert.Equal(t, tt.pipelineCreateWant, createPipeline)
				})

				t.Run("get", func(t *testing.T) {
					getPipeline, _, err := intClient.PipelineConfigs.Get(ctx, tt.pipelineGet)
					assert.NoError(t, err)

					getPipeline.RemoveLinks()
					assert.Equal(t, tt.pipelineGetWant, getPipeline)
				})

				t.Run("update", func(t *testing.T) {

					updatePipeline, _, err := intClient.PipelineConfigs.Update(context.Background(), p.Name, p)
					assert.NoError(t, err)
					assert.NotEqual(t, getPipeline.Version, updatePipeline.Version)
					updatePipeline.Version = getPipeline.Version

					updatePipeline.RemoveLinks()
					assert.Equal(t, tt.pipelineUpdateWant, updatePipeline)
				})

				t.Run("delete", func(t *testing.T) {
					message, _, err := intClient.PipelineConfigs.Delete(ctx, tt.delete)
					assert.NoError(t, err)
					assert.Equal(t, tt.deleteWant, message)
				})
			}
		})
	}
}

func buildMockPipelineStages() []*Stage {
	return []*Stage{{
		Name: "defaultStage",
		Jobs: []*Job{{
			Name: "defaultJob",
			Tasks: []*Task{{
				Type: "exec",
				Attributes: TaskAttributes{
					RunIf:   []string{"passed"},
					Command: "ls",
				},
			}},
			Tabs:                 make([]*Tab, 0),
			Artifacts:            make([]*Artifact, 0),
			EnvironmentVariables: make([]*EnvironmentVariable, 0),
			Resources:            []string{},
		}},
		Approval: &Approval{
			Type: "success",
			Authorization: &Authorization{
				Users: make([]string, 0),
				Roles: make([]string, 0),
			},
		},
		EnvironmentVariables: make([]*EnvironmentVariable, 0),
	}}
}
