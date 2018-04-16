

# gocd
`import "github.com/beamly/go-gocd/gocd"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>
Package gocd provides a client for using the GoCD Server API.

Usage:


	import "github.com/beamly/go-gocd/gocd"

Construct a new GoCD client and supply the URL to your GoCD server and if required, username and password. Then use the
various services on the client to access different parts of the GoCD API.
For example:


	package main
	import (
		"github.com/beamly/go-gocd/gocd"
		"context"
		"fmt"
	)
	
	func main() {
		cfg := gocd.Configuration{
			Server: "<a href="https://my_gocd/go/">https://my_gocd/go/</a>",
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

If you wish to use your own http client, you can use the following idiom


	package main
	
	import (
		"github.com/beamly/go-gocd/gocd"
		"net/http"
		"context"
	)
	
	func main() {
		client := gocd.NewClient(
			&gocd.Configuration{},
			&http.Client{},
		)
		client.Login(context.Background())
	}

The services of a client divide the API into logical chunks and correspond to
the structure of the GoCD API documentation at
<a href="https://api.gocd.org/current/">https://api.gocd.org/current/</a>.




## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func CheckResponse(response *APIResponse) (err error)](#CheckResponse)
* [func ConfigFilePath() (configPath string, err error)](#ConfigFilePath)
* [func LoadConfigByName(name string, cfg *Configuration) (err error)](#LoadConfigByName)
* [func LoadConfigFromFile() (cfgs map[string]*Configuration, err error)](#LoadConfigFromFile)
* [func SetupLogging(log *logrus.Logger)](#SetupLogging)
* [type APIClientRequest](#APIClientRequest)
* [type APIRequest](#APIRequest)
* [type APIResponse](#APIResponse)
* [type Agent](#Agent)
  * [func (a *Agent) GetLinks() *HALLinks](#Agent.GetLinks)
  * [func (a *Agent) RemoveLinks()](#Agent.RemoveLinks)
* [type AgentBulkOperationUpdate](#AgentBulkOperationUpdate)
* [type AgentBulkOperationsUpdate](#AgentBulkOperationsUpdate)
* [type AgentBulkUpdate](#AgentBulkUpdate)
* [type AgentsResponse](#AgentsResponse)
* [type AgentsService](#AgentsService)
  * [func (s *AgentsService) BulkUpdate(ctx context.Context, agents AgentBulkUpdate) (message string, resp *APIResponse, err error)](#AgentsService.BulkUpdate)
  * [func (s *AgentsService) Delete(ctx context.Context, uuid string) (string, *APIResponse, error)](#AgentsService.Delete)
  * [func (s *AgentsService) Get(ctx context.Context, uuid string) (*Agent, *APIResponse, error)](#AgentsService.Get)
  * [func (s *AgentsService) JobRunHistory(ctx context.Context, uuid string) (jobs []*Job, resp *APIResponse, err error)](#AgentsService.JobRunHistory)
  * [func (s *AgentsService) List(ctx context.Context) (agents []*Agent, resp *APIResponse, err error)](#AgentsService.List)
  * [func (s *AgentsService) Update(ctx context.Context, uuid string, agent *Agent) (*Agent, *APIResponse, error)](#AgentsService.Update)
* [type Approval](#Approval)
  * [func (a *Approval) Clean()](#Approval.Clean)
* [type Artifact](#Artifact)
* [type Auth](#Auth)
* [type Authorization](#Authorization)
* [type BuildCause](#BuildCause)
* [type BuildDetails](#BuildDetails)
* [type CipherText](#CipherText)
* [type Client](#Client)
  * [func NewClient(cfg *Configuration, httpClient *http.Client) *Client](#NewClient)
  * [func (c *Client) Do(ctx context.Context, req *APIRequest, v interface{}, responseType string) (*APIResponse, error)](#Client.Do)
  * [func (c *Client) Lock()](#Client.Lock)
  * [func (c *Client) Login(ctx context.Context) (err error)](#Client.Login)
  * [func (c *Client) NewRequest(method, urlStr string, body interface{}, apiVersion string) (req *APIRequest, err error)](#Client.NewRequest)
  * [func (c *Client) Unlock()](#Client.Unlock)
* [type ConfigApproval](#ConfigApproval)
* [type ConfigArtifact](#ConfigArtifact)
* [type ConfigAuthConfig](#ConfigAuthConfig)
* [type ConfigElastic](#ConfigElastic)
* [type ConfigElasticProfile](#ConfigElasticProfile)
* [type ConfigEnvironmentVariable](#ConfigEnvironmentVariable)
* [type ConfigFilter](#ConfigFilter)
* [type ConfigJob](#ConfigJob)
* [type ConfigMaterialRepository](#ConfigMaterialRepository)
* [type ConfigPackage](#ConfigPackage)
* [type ConfigParam](#ConfigParam)
* [type ConfigPipeline](#ConfigPipeline)
* [type ConfigPipelineGroup](#ConfigPipelineGroup)
* [type ConfigPluginConfiguration](#ConfigPluginConfiguration)
* [type ConfigProperty](#ConfigProperty)
* [type ConfigRepo](#ConfigRepo)
  * [func (c *ConfigRepo) GetVersion() (version string)](#ConfigRepo.GetVersion)
  * [func (c *ConfigRepo) SetVersion(version string)](#ConfigRepo.SetVersion)
* [type ConfigRepoProperty](#ConfigRepoProperty)
* [type ConfigRepoService](#ConfigRepoService)
  * [func (crs *ConfigRepoService) Create(ctx context.Context, cr *ConfigRepo) (out *ConfigRepo, resp *APIResponse, err error)](#ConfigRepoService.Create)
  * [func (crs *ConfigRepoService) Delete(ctx context.Context, id string) (string, *APIResponse, error)](#ConfigRepoService.Delete)
  * [func (crs *ConfigRepoService) Get(ctx context.Context, id string) (out *ConfigRepo, resp *APIResponse, err error)](#ConfigRepoService.Get)
  * [func (crs *ConfigRepoService) List(ctx context.Context) (repos []*ConfigRepo, resp *APIResponse, err error)](#ConfigRepoService.List)
  * [func (crs *ConfigRepoService) Update(ctx context.Context, id string, cr *ConfigRepo) (out *ConfigRepo, resp *APIResponse, err error)](#ConfigRepoService.Update)
* [type ConfigReposListResponse](#ConfigReposListResponse)
* [type ConfigRepository](#ConfigRepository)
* [type ConfigRepositoryGit](#ConfigRepositoryGit)
* [type ConfigRole](#ConfigRole)
* [type ConfigSCM](#ConfigSCM)
* [type ConfigSecurity](#ConfigSecurity)
* [type ConfigServer](#ConfigServer)
* [type ConfigStage](#ConfigStage)
* [type ConfigTask](#ConfigTask)
* [type ConfigTaskRunIf](#ConfigTaskRunIf)
* [type ConfigTasks](#ConfigTasks)
* [type ConfigXML](#ConfigXML)
* [type Configuration](#Configuration)
  * [func (c *Configuration) Client() *Client](#Configuration.Client)
  * [func (c *Configuration) HasAuth() bool](#Configuration.HasAuth)
* [type ConfigurationService](#ConfigurationService)
  * [func (cs *ConfigurationService) Get(ctx context.Context) (cx *ConfigXML, resp *APIResponse, err error)](#ConfigurationService.Get)
  * [func (cs *ConfigurationService) GetVersion(ctx context.Context) (v *Version, resp *APIResponse, err error)](#ConfigurationService.GetVersion)
* [type EmbeddedEnvironments](#EmbeddedEnvironments)
* [type EncryptionService](#EncryptionService)
  * [func (es *EncryptionService) Encrypt(ctx context.Context, plaintext string) (c *CipherText, resp *APIResponse, err error)](#EncryptionService.Encrypt)
* [type Environment](#Environment)
  * [func (env *Environment) GetLinks() *HALLinks](#Environment.GetLinks)
  * [func (env *Environment) GetVersion() (version string)](#Environment.GetVersion)
  * [func (env *Environment) RemoveLinks()](#Environment.RemoveLinks)
  * [func (env *Environment) SetVersion(version string)](#Environment.SetVersion)
* [type EnvironmentPatchRequest](#EnvironmentPatchRequest)
* [type EnvironmentVariable](#EnvironmentVariable)
* [type EnvironmentVariablesAction](#EnvironmentVariablesAction)
* [type EnvironmentsResponse](#EnvironmentsResponse)
  * [func (er *EnvironmentsResponse) GetLinks() *HALLinks](#EnvironmentsResponse.GetLinks)
  * [func (er *EnvironmentsResponse) RemoveLinks()](#EnvironmentsResponse.RemoveLinks)
* [type EnvironmentsService](#EnvironmentsService)
  * [func (es *EnvironmentsService) Create(ctx context.Context, name string) (e *Environment, resp *APIResponse, err error)](#EnvironmentsService.Create)
  * [func (es *EnvironmentsService) Delete(ctx context.Context, name string) (string, *APIResponse, error)](#EnvironmentsService.Delete)
  * [func (es *EnvironmentsService) Get(ctx context.Context, name string) (e *Environment, resp *APIResponse, err error)](#EnvironmentsService.Get)
  * [func (es *EnvironmentsService) List(ctx context.Context) (e *EnvironmentsResponse, resp *APIResponse, err error)](#EnvironmentsService.List)
  * [func (es *EnvironmentsService) Patch(ctx context.Context, name string, patch *EnvironmentPatchRequest) (e *Environment, resp *APIResponse, err error)](#EnvironmentsService.Patch)
* [type GitRepositoryMaterial](#GitRepositoryMaterial)
* [type HALContainer](#HALContainer)
* [type HALLink](#HALLink)
* [type HALLinks](#HALLinks)
  * [func (al *HALLinks) Add(link *HALLink)](#HALLinks.Add)
  * [func (al HALLinks) Get(name string) (link *HALLink)](#HALLinks.Get)
  * [func (al HALLinks) GetOk(name string) (link *HALLink, ok bool)](#HALLinks.GetOk)
  * [func (al HALLinks) Keys() (keys []string)](#HALLinks.Keys)
  * [func (al HALLinks) MarshallJSON() ([]byte, error)](#HALLinks.MarshallJSON)
  * [func (al *HALLinks) UnmarshalJSON(j []byte) (err error)](#HALLinks.UnmarshalJSON)
* [type Job](#Job)
  * [func (j *Job) JSONString() (body string, err error)](#Job.JSONString)
  * [func (j *Job) Validate() (err error)](#Job.Validate)
* [type JobProperty](#JobProperty)
* [type JobRunHistoryResponse](#JobRunHistoryResponse)
* [type JobSchedule](#JobSchedule)
* [type JobScheduleEnvVar](#JobScheduleEnvVar)
* [type JobScheduleLink](#JobScheduleLink)
* [type JobScheduleResponse](#JobScheduleResponse)
* [type JobStateTransition](#JobStateTransition)
* [type JobsService](#JobsService)
  * [func (js *JobsService) ListScheduled(ctx context.Context) (jobs []*JobSchedule, resp *APIResponse, err error)](#JobsService.ListScheduled)
* [type MailHost](#MailHost)
* [type Material](#Material)
  * [func (m Material) Equal(a *Material) (isEqual bool, err error)](#Material.Equal)
  * [func (m *Material) Ingest(payload map[string]interface{}) (err error)](#Material.Ingest)
  * [func (m *Material) IngestAttributeGenerics(i interface{}) (err error)](#Material.IngestAttributeGenerics)
  * [func (m *Material) IngestAttributes(rawAttributes map[string]interface{}) (err error)](#Material.IngestAttributes)
  * [func (m *Material) IngestType(payload map[string]interface{})](#Material.IngestType)
  * [func (m *Material) UnmarshalJSON(b []byte) (err error)](#Material.UnmarshalJSON)
* [type MaterialAttribute](#MaterialAttribute)
* [type MaterialAttributesDependency](#MaterialAttributesDependency)
  * [func (mad MaterialAttributesDependency) GenerateGeneric() (ma map[string]interface{})](#MaterialAttributesDependency.GenerateGeneric)
  * [func (mad MaterialAttributesDependency) GetFilter() *MaterialFilter](#MaterialAttributesDependency.GetFilter)
  * [func (mad MaterialAttributesDependency) HasFilter() bool](#MaterialAttributesDependency.HasFilter)
* [type MaterialAttributesGit](#MaterialAttributesGit)
  * [func (mag MaterialAttributesGit) GenerateGeneric() (ma map[string]interface{})](#MaterialAttributesGit.GenerateGeneric)
  * [func (mag MaterialAttributesGit) GetFilter() *MaterialFilter](#MaterialAttributesGit.GetFilter)
  * [func (mag MaterialAttributesGit) HasFilter() bool](#MaterialAttributesGit.HasFilter)
* [type MaterialAttributesHg](#MaterialAttributesHg)
  * [func (mhg MaterialAttributesHg) GenerateGeneric() (ma map[string]interface{})](#MaterialAttributesHg.GenerateGeneric)
  * [func (mhg MaterialAttributesHg) GetFilter() *MaterialFilter](#MaterialAttributesHg.GetFilter)
  * [func (mhg MaterialAttributesHg) HasFilter() bool](#MaterialAttributesHg.HasFilter)
* [type MaterialAttributesP4](#MaterialAttributesP4)
  * [func (mp4 MaterialAttributesP4) GenerateGeneric() (ma map[string]interface{})](#MaterialAttributesP4.GenerateGeneric)
  * [func (mp4 MaterialAttributesP4) GetFilter() *MaterialFilter](#MaterialAttributesP4.GetFilter)
  * [func (mp4 MaterialAttributesP4) HasFilter() bool](#MaterialAttributesP4.HasFilter)
* [type MaterialAttributesPackage](#MaterialAttributesPackage)
  * [func (mapk MaterialAttributesPackage) GenerateGeneric() (ma map[string]interface{})](#MaterialAttributesPackage.GenerateGeneric)
  * [func (mapk MaterialAttributesPackage) GetFilter() *MaterialFilter](#MaterialAttributesPackage.GetFilter)
  * [func (mapk MaterialAttributesPackage) HasFilter() bool](#MaterialAttributesPackage.HasFilter)
* [type MaterialAttributesPlugin](#MaterialAttributesPlugin)
  * [func (mapp MaterialAttributesPlugin) GenerateGeneric() (ma map[string]interface{})](#MaterialAttributesPlugin.GenerateGeneric)
  * [func (mapp MaterialAttributesPlugin) GetFilter() *MaterialFilter](#MaterialAttributesPlugin.GetFilter)
  * [func (mapp MaterialAttributesPlugin) HasFilter() bool](#MaterialAttributesPlugin.HasFilter)
* [type MaterialAttributesSvn](#MaterialAttributesSvn)
  * [func (mas MaterialAttributesSvn) GenerateGeneric() (ma map[string]interface{})](#MaterialAttributesSvn.GenerateGeneric)
  * [func (mas MaterialAttributesSvn) GetFilter() *MaterialFilter](#MaterialAttributesSvn.GetFilter)
  * [func (mas MaterialAttributesSvn) HasFilter() bool](#MaterialAttributesSvn.HasFilter)
* [type MaterialAttributesTfs](#MaterialAttributesTfs)
  * [func (mtfs MaterialAttributesTfs) GenerateGeneric() (ma map[string]interface{})](#MaterialAttributesTfs.GenerateGeneric)
  * [func (mtfs MaterialAttributesTfs) GetFilter() *MaterialFilter](#MaterialAttributesTfs.GetFilter)
  * [func (mtfs MaterialAttributesTfs) HasFilter() bool](#MaterialAttributesTfs.HasFilter)
* [type MaterialFilter](#MaterialFilter)
  * [func (mf *MaterialFilter) GenerateGeneric() (g map[string]interface{})](#MaterialFilter.GenerateGeneric)
* [type MaterialRevision](#MaterialRevision)
* [type Modification](#Modification)
* [type PaginationResponse](#PaginationResponse)
* [type Parameter](#Parameter)
* [type PasswordFilePath](#PasswordFilePath)
* [type PatchStringAction](#PatchStringAction)
* [type Pipeline](#Pipeline)
  * [func (p *Pipeline) AddStage(stage *Stage)](#Pipeline.AddStage)
  * [func (p *Pipeline) GetLinks() *HALLinks](#Pipeline.GetLinks)
  * [func (p *Pipeline) GetName() string](#Pipeline.GetName)
  * [func (p *Pipeline) GetStage(stageName string) (stage *Stage)](#Pipeline.GetStage)
  * [func (p *Pipeline) GetStages() []*Stage](#Pipeline.GetStages)
  * [func (p *Pipeline) GetVersion() (version string)](#Pipeline.GetVersion)
  * [func (p *Pipeline) RemoveLinks()](#Pipeline.RemoveLinks)
  * [func (p *Pipeline) SetStage(newStage *Stage)](#Pipeline.SetStage)
  * [func (p *Pipeline) SetStages(stages []*Stage)](#Pipeline.SetStages)
  * [func (p *Pipeline) SetVersion(version string)](#Pipeline.SetVersion)
* [type PipelineConfigOrigin](#PipelineConfigOrigin)
* [type PipelineConfigRequest](#PipelineConfigRequest)
  * [func (pr *PipelineConfigRequest) GetVersion() (version string)](#PipelineConfigRequest.GetVersion)
  * [func (pr *PipelineConfigRequest) SetVersion(version string)](#PipelineConfigRequest.SetVersion)
* [type PipelineConfigsService](#PipelineConfigsService)
  * [func (pcs *PipelineConfigsService) Create(ctx context.Context, group string, p *Pipeline) (pr *Pipeline, resp *APIResponse, err error)](#PipelineConfigsService.Create)
  * [func (pcs *PipelineConfigsService) Delete(ctx context.Context, name string) (string, *APIResponse, error)](#PipelineConfigsService.Delete)
  * [func (pcs *PipelineConfigsService) Get(ctx context.Context, name string) (p *Pipeline, resp *APIResponse, err error)](#PipelineConfigsService.Get)
  * [func (pcs *PipelineConfigsService) Update(ctx context.Context, name string, p *Pipeline) (pr *Pipeline, resp *APIResponse, err error)](#PipelineConfigsService.Update)
* [type PipelineGroup](#PipelineGroup)
* [type PipelineGroups](#PipelineGroups)
  * [func (pg *PipelineGroups) GetGroupByPipeline(pipeline *Pipeline) *PipelineGroup](#PipelineGroups.GetGroupByPipeline)
  * [func (pg *PipelineGroups) GetGroupByPipelineName(pipelineName string) *PipelineGroup](#PipelineGroups.GetGroupByPipelineName)
* [type PipelineGroupsService](#PipelineGroupsService)
  * [func (pgs *PipelineGroupsService) List(ctx context.Context, name string) (*PipelineGroups, *APIResponse, error)](#PipelineGroupsService.List)
* [type PipelineHistory](#PipelineHistory)
* [type PipelineInstance](#PipelineInstance)
* [type PipelineMaterial](#PipelineMaterial)
* [type PipelineRequest](#PipelineRequest)
* [type PipelineStatus](#PipelineStatus)
* [type PipelineTemplate](#PipelineTemplate)
  * [func (pt *PipelineTemplate) AddStage(stage *Stage)](#PipelineTemplate.AddStage)
  * [func (pt PipelineTemplate) GetName() string](#PipelineTemplate.GetName)
  * [func (pt PipelineTemplate) GetStage(stageName string) *Stage](#PipelineTemplate.GetStage)
  * [func (pt PipelineTemplate) GetStages() []*Stage](#PipelineTemplate.GetStages)
  * [func (pt PipelineTemplate) GetVersion() (version string)](#PipelineTemplate.GetVersion)
  * [func (pt PipelineTemplate) Pipelines() []*Pipeline](#PipelineTemplate.Pipelines)
  * [func (pt *PipelineTemplate) RemoveLinks()](#PipelineTemplate.RemoveLinks)
  * [func (pt *PipelineTemplate) SetStage(newStage *Stage)](#PipelineTemplate.SetStage)
  * [func (pt *PipelineTemplate) SetStages(stages []*Stage)](#PipelineTemplate.SetStages)
  * [func (pt *PipelineTemplate) SetVersion(version string)](#PipelineTemplate.SetVersion)
* [type PipelineTemplateRequest](#PipelineTemplateRequest)
  * [func (pt PipelineTemplateRequest) GetVersion() (version string)](#PipelineTemplateRequest.GetVersion)
  * [func (pt *PipelineTemplateRequest) SetVersion(version string)](#PipelineTemplateRequest.SetVersion)
* [type PipelineTemplateResponse](#PipelineTemplateResponse)
* [type PipelineTemplatesResponse](#PipelineTemplatesResponse)
* [type PipelineTemplatesService](#PipelineTemplatesService)
  * [func (pts *PipelineTemplatesService) Create(ctx context.Context, name string, st []*Stage) (ptr *PipelineTemplate, resp *APIResponse, err error)](#PipelineTemplatesService.Create)
  * [func (pts *PipelineTemplatesService) Delete(ctx context.Context, name string) (string, *APIResponse, error)](#PipelineTemplatesService.Delete)
  * [func (pts *PipelineTemplatesService) Get(ctx context.Context, name string) (pt *PipelineTemplate, resp *APIResponse, err error)](#PipelineTemplatesService.Get)
  * [func (pts *PipelineTemplatesService) List(ctx context.Context) (pt []*PipelineTemplate, resp *APIResponse, err error)](#PipelineTemplatesService.List)
  * [func (pts *PipelineTemplatesService) Update(ctx context.Context, name string, template *PipelineTemplate) (ptr *PipelineTemplate, resp *APIResponse, err error)](#PipelineTemplatesService.Update)
* [type PipelinesService](#PipelinesService)
  * [func (pgs *PipelinesService) GetHistory(ctx context.Context, name string, offset int) (pt *PipelineHistory, resp *APIResponse, err error)](#PipelinesService.GetHistory)
  * [func (pgs *PipelinesService) GetInstance(ctx context.Context, name string, offset int) (pt *PipelineInstance, resp *APIResponse, err error)](#PipelinesService.GetInstance)
  * [func (pgs *PipelinesService) GetStatus(ctx context.Context, name string, offset int) (ps *PipelineStatus, resp *APIResponse, err error)](#PipelinesService.GetStatus)
  * [func (pgs *PipelinesService) Pause(ctx context.Context, name string) (bool, *APIResponse, error)](#PipelinesService.Pause)
  * [func (pgs *PipelinesService) ReleaseLock(ctx context.Context, name string) (bool, *APIResponse, error)](#PipelinesService.ReleaseLock)
  * [func (pgs *PipelinesService) Unpause(ctx context.Context, name string) (bool, *APIResponse, error)](#PipelinesService.Unpause)
* [type PluggableInstanceSettings](#PluggableInstanceSettings)
* [type Plugin](#Plugin)
* [type PluginConfiguration](#PluginConfiguration)
* [type PluginConfigurationKVPair](#PluginConfigurationKVPair)
* [type PluginConfigurationMetadata](#PluginConfigurationMetadata)
* [type PluginView](#PluginView)
* [type PluginsResponse](#PluginsResponse)
* [type PluginsService](#PluginsService)
  * [func (ps *PluginsService) Get(ctx context.Context, name string) (p *Plugin, resp *APIResponse, err error)](#PluginsService.Get)
  * [func (ps *PluginsService) List(ctx context.Context) (*PluginsResponse, *APIResponse, error)](#PluginsService.List)
* [type Properties](#Properties)
  * [func NewPropertiesFrame(frame [][]string) *Properties](#NewPropertiesFrame)
  * [func (pr *Properties) AddRow(r []string)](#Properties.AddRow)
  * [func (pr Properties) Get(row int, column string) string](#Properties.Get)
  * [func (pr *Properties) MarshalJSON() ([]byte, error)](#Properties.MarshalJSON)
  * [func (pr Properties) MarshallCSV() (string, error)](#Properties.MarshallCSV)
  * [func (pr *Properties) SetRow(row int, r []string)](#Properties.SetRow)
  * [func (pr *Properties) UnmarshallCSV(raw string) error](#Properties.UnmarshallCSV)
  * [func (pr *Properties) Write(p []byte) (n int, err error)](#Properties.Write)
* [type PropertiesService](#PropertiesService)
  * [func (ps *PropertiesService) Create(ctx context.Context, name string, value string, pr *PropertyRequest) (responseIsValid bool, resp *APIResponse, err error)](#PropertiesService.Create)
  * [func (ps *PropertiesService) Get(ctx context.Context, name string, pr *PropertyRequest) (*Properties, *APIResponse, error)](#PropertiesService.Get)
  * [func (ps *PropertiesService) List(ctx context.Context, pr *PropertyRequest) (*Properties, *APIResponse, error)](#PropertiesService.List)
  * [func (ps *PropertiesService) ListHistorical(ctx context.Context, pr *PropertyRequest) (*Properties, *APIResponse, error)](#PropertiesService.ListHistorical)
* [type PropertyCreateResponse](#PropertyCreateResponse)
* [type PropertyRequest](#PropertyRequest)
* [type Stage](#Stage)
  * [func (s *Stage) Clean()](#Stage.Clean)
  * [func (s *Stage) JSONString() (string, error)](#Stage.JSONString)
  * [func (s *Stage) Validate() error](#Stage.Validate)
* [type StageContainer](#StageContainer)
* [type StagesService](#StagesService)
* [type StringResponse](#StringResponse)
* [type Tab](#Tab)
* [type Task](#Task)
  * [func (t *Task) Validate() error](#Task.Validate)
* [type TaskAttributes](#TaskAttributes)
  * [func (t *TaskAttributes) ValidateAnt() error](#TaskAttributes.ValidateAnt)
  * [func (t *TaskAttributes) ValidateExec() error](#TaskAttributes.ValidateExec)
* [type TaskPluginConfiguration](#TaskPluginConfiguration)
* [type TimeoutField](#TimeoutField)
  * [func (tf TimeoutField) MarshalJSON() (b []byte, err error)](#TimeoutField.MarshalJSON)
  * [func (tf *TimeoutField) UnmarshalJSON(b []byte) (err error)](#TimeoutField.UnmarshalJSON)
* [type Version](#Version)
* [type Versioned](#Versioned)

#### <a name="pkg-examples">Examples</a>
* [AgentsService.List](#example_AgentsService_List)
* [ConfigRepoService.Get](#example_ConfigRepoService_Get)
* [ConfigRepoService.List](#example_ConfigRepoService_List)
* [PipelineConfigsService.Get](#example_PipelineConfigsService_Get)

#### <a name="pkg-files">Package files</a>
[agent.go](https://github.com/beamly/go-gocd/tree/master/gocd/agent.go) [approval.go](https://github.com/beamly/go-gocd/tree/master/gocd/approval.go) [authentication.go](https://github.com/beamly/go-gocd/tree/master/gocd/authentication.go) [config.go](https://github.com/beamly/go-gocd/tree/master/gocd/config.go) [config_repo.go](https://github.com/beamly/go-gocd/tree/master/gocd/config_repo.go) [configuration.go](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go) [configuration_task.go](https://github.com/beamly/go-gocd/tree/master/gocd/configuration_task.go) [doc.go](https://github.com/beamly/go-gocd/tree/master/gocd/doc.go) [encryption.go](https://github.com/beamly/go-gocd/tree/master/gocd/encryption.go) [environment.go](https://github.com/beamly/go-gocd/tree/master/gocd/environment.go) [genericactions.go](https://github.com/beamly/go-gocd/tree/master/gocd/genericactions.go) [gocd.go](https://github.com/beamly/go-gocd/tree/master/gocd/gocd.go) [jobs.go](https://github.com/beamly/go-gocd/tree/master/gocd/jobs.go) [jobs_validation.go](https://github.com/beamly/go-gocd/tree/master/gocd/jobs_validation.go) [links.go](https://github.com/beamly/go-gocd/tree/master/gocd/links.go) [logging.go](https://github.com/beamly/go-gocd/tree/master/gocd/logging.go) [pipeline.go](https://github.com/beamly/go-gocd/tree/master/gocd/pipeline.go) [pipeline_material.go](https://github.com/beamly/go-gocd/tree/master/gocd/pipeline_material.go) [pipelineconfig.go](https://github.com/beamly/go-gocd/tree/master/gocd/pipelineconfig.go) [pipelinegroups.go](https://github.com/beamly/go-gocd/tree/master/gocd/pipelinegroups.go) [pipelinetemplate.go](https://github.com/beamly/go-gocd/tree/master/gocd/pipelinetemplate.go) [plugin.go](https://github.com/beamly/go-gocd/tree/master/gocd/plugin.go) [properties.go](https://github.com/beamly/go-gocd/tree/master/gocd/properties.go) [resource.go](https://github.com/beamly/go-gocd/tree/master/gocd/resource.go) [resource_agent.go](https://github.com/beamly/go-gocd/tree/master/gocd/resource_agent.go) [resource_approval.go](https://github.com/beamly/go-gocd/tree/master/gocd/resource_approval.go) [resource_config_repo.go](https://github.com/beamly/go-gocd/tree/master/gocd/resource_config_repo.go) [resource_environment.go](https://github.com/beamly/go-gocd/tree/master/gocd/resource_environment.go) [resource_jobs.go](https://github.com/beamly/go-gocd/tree/master/gocd/resource_jobs.go) [resource_pipeline.go](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline.go) [resource_pipeline_material.go](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material.go) [resource_pipeline_material_dependency.go](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_dependency.go) [resource_pipeline_material_git.go](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_git.go) [resource_pipeline_material_hg.go](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_hg.go) [resource_pipeline_material_p4.go](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_p4.go) [resource_pipeline_material_pkg.go](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_pkg.go) [resource_pipeline_material_plugin.go](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_plugin.go) [resource_pipeline_material_svn.go](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_svn.go) [resource_pipeline_material_tfs.go](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_tfs.go) [resource_pipelinegroups.go](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipelinegroups.go) [resource_pipelinetemplate.go](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipelinetemplate.go) [resource_properties.go](https://github.com/beamly/go-gocd/tree/master/gocd/resource_properties.go) [resource_stages.go](https://github.com/beamly/go-gocd/tree/master/gocd/resource_stages.go) [resource_task.go](https://github.com/beamly/go-gocd/tree/master/gocd/resource_task.go) [stages.go](https://github.com/beamly/go-gocd/tree/master/gocd/stages.go) 


## <a name="pkg-constants">Constants</a>
``` go
const (
    EnvVarDefaultProfile = "GOCD_DEFAULT_PROFILE"
    EnvVarServer         = "GOCD_SERVER"
    EnvVarUsername       = "GOCD_USERNAME"
    EnvVarPassword       = "GOCD_PASSWORD"
    EnvVarSkipSsl        = "GOCD_SKIP_SSL_CHECK"
)
```
Environment variables for configuration.

``` go
const (
    // JobStateTransitionPassed "Passed"
    JobStateTransitionPassed = "Passed"
    // JobStateTransitionScheduled "Scheduled"
    JobStateTransitionScheduled = "Scheduled"
)
```
``` go
const (
    LogLevelEnvVarName = "GOCD_LOG_LEVEL"
    LogLevelDefault    = "WARNING"
    LogTypeEnvVarName  = "GOCD_LOG_TYPE"
    LogTypeDefault     = "TEXT"
)
```
Set logging level and type constants

``` go
const ConfigDirectoryPath = "~/.gocd.conf"
```
ConfigDirectoryPath is the default location of the `.gocdconf` configuration file




## <a name="CheckResponse">func</a> [CheckResponse](https://github.com/beamly/go-gocd/tree/master/gocd/gocd.go?s=7716:7769#L311)
``` go
func CheckResponse(response *APIResponse) (err error)
```
CheckResponse asserts that the http response status code was 2xx.



## <a name="ConfigFilePath">func</a> [ConfigFilePath](https://github.com/beamly/go-gocd/tree/master/gocd/config.go?s=1948:2000#L88)
``` go
func ConfigFilePath() (configPath string, err error)
```
ConfigFilePath specifies the default path to a config file



## <a name="LoadConfigByName">func</a> [LoadConfigByName](https://github.com/beamly/go-gocd/tree/master/gocd/config.go?s=863:929#L34)
``` go
func LoadConfigByName(name string, cfg *Configuration) (err error)
```
LoadConfigByName loads configurations from yaml at the default file location



## <a name="LoadConfigFromFile">func</a> [LoadConfigFromFile](https://github.com/beamly/go-gocd/tree/master/gocd/config.go?s=1495:1564#L64)
``` go
func LoadConfigFromFile() (cfgs map[string]*Configuration, err error)
```
LoadConfigFromFile on disk and return it as a Configuration item



## <a name="SetupLogging">func</a> [SetupLogging](https://github.com/beamly/go-gocd/tree/master/gocd/logging.go?s=885:922#L46)
``` go
func SetupLogging(log *logrus.Logger)
```
SetupLogging based on Environment Variables


	Set Logging level with $GOCD_LOG_LEVEL
	Allowed Values:
	  - DEBUG
	  - INFO
	  - WARNING
	  - ERROR
	  - FATAL
	  - PANIC
	
	Set Logging type  with $GOCD_LOG_TYPE
	Allowed Values:
	  - JSON
	  - TEXT




## <a name="APIClientRequest">type</a> [APIClientRequest](https://github.com/beamly/go-gocd/tree/master/gocd/genericactions.go?s=175:375#L14)
``` go
type APIClientRequest struct {
    Method       string
    Path         string
    APIVersion   string
    RequestBody  interface{}
    ResponseType string
    ResponseBody interface{}
    Headers      map[string]string
}
```
APIClientRequest helper struct to reduce amount of code.










## <a name="APIRequest">type</a> [APIRequest](https://github.com/beamly/go-gocd/tree/master/gocd/gocd.go?s=1392:1451#L63)
``` go
type APIRequest struct {
    HTTP *http.Request
    Body string
}
```
APIRequest encapsulates the net/http.Request object, and a string representing the Body.










## <a name="APIResponse">type</a> [APIResponse](https://github.com/beamly/go-gocd/tree/master/gocd/gocd.go?s=1210:1298#L56)
``` go
type APIResponse struct {
    HTTP    *http.Response
    Body    string
    Request *APIRequest
}
```
APIResponse encapsulates the net/http.Response object, a string representing the Body, and a gocd.Request object
encapsulating the response from the API.










## <a name="Agent">type</a> [Agent](https://github.com/beamly/go-gocd/tree/master/gocd/agent.go?s=442:1447#L20)
``` go
type Agent struct {
    UUID             string        `json:"uuid,omitempty"`
    Hostname         string        `json:"hostname,omitempty"`
    ElasticAgentID   string        `json:"elastic_agent_id,omitempty"`
    ElasticPluginID  string        `json:"elastic_plugin_id,omitempty"`
    IPAddress        string        `json:"ip_address,omitempty"`
    Sandbox          string        `json:"sandbox,omitempty"`
    OperatingSystem  string        `json:"operating_system,omitempty"`
    FreeSpace        int           `json:"free_space,omitempty"`
    AgentConfigState string        `json:"agent_config_state,omitempty"`
    AgentState       string        `json:"agent_state,omitempty"`
    Resources        []string      `json:"resources,omitempty"`
    Environments     []string      `json:"environments,omitempty"`
    BuildState       string        `json:"build_state,omitempty"`
    BuildDetails     *BuildDetails `json:"build_details,omitempty"`
    Links            *HALLinks     `json:"_links,omitempty,omitempty"`
    // contains filtered or unexported fields
}
```
Agent represents agent in GoCD










### <a name="Agent.GetLinks">func</a> (\*Agent) [GetLinks](https://github.com/beamly/go-gocd/tree/master/gocd/resource_agent.go?s=54:90#L4)
``` go
func (a *Agent) GetLinks() *HALLinks
```
GetLinks returns HAL links for agent




### <a name="Agent.RemoveLinks">func</a> (\*Agent) [RemoveLinks](https://github.com/beamly/go-gocd/tree/master/gocd/resource_agent.go?s=210:239#L9)
``` go
func (a *Agent) RemoveLinks()
```
RemoveLinks sets the `Link` attribute as `nil`. Used when rendering an `Agent` struct to JSON.




## <a name="AgentBulkOperationUpdate">type</a> [AgentBulkOperationUpdate](https://github.com/beamly/go-gocd/tree/master/gocd/agent.go?s=2239:2362#L54)
``` go
type AgentBulkOperationUpdate struct {
    Add    []string `json:"add,omitempty"`
    Remove []string `json:"remove,omitempty"`
}
```
AgentBulkOperationUpdate describes an action to be performed on an Environment or Resource during an agent update.










## <a name="AgentBulkOperationsUpdate">type</a> [AgentBulkOperationsUpdate](https://github.com/beamly/go-gocd/tree/master/gocd/agent.go?s=1937:2119#L48)
``` go
type AgentBulkOperationsUpdate struct {
    Environments *AgentBulkOperationUpdate `json:"environments,omitempty"`
    Resources    *AgentBulkOperationUpdate `json:"resources,omitempty"`
}
```
AgentBulkOperationsUpdate describes the structure for a single Operation in AgentBulkUpdate the PUT payload when
updating multiple agents










## <a name="AgentBulkUpdate">type</a> [AgentBulkUpdate](https://github.com/beamly/go-gocd/tree/master/gocd/agent.go?s=1542:1791#L40)
``` go
type AgentBulkUpdate struct {
    Uuids            []string                   `json:"uuids"`
    Operations       *AgentBulkOperationsUpdate `json:"operations,omitempty"`
    AgentConfigState string                     `json:"agent_config_state,omitempty"`
}
```
AgentBulkUpdate describes the structure for the PUT payload when updating multiple agents










## <a name="AgentsResponse">type</a> [AgentsResponse](https://github.com/beamly/go-gocd/tree/master/gocd/agent.go?s=244:406#L12)
``` go
type AgentsResponse struct {
    Links    *HALLinks `json:"_links,omitempty"`
    Embedded *struct {
        Agents []*Agent `json:"agents"`
    } `json:"_embedded,omitempty"`
}
```
AgentsResponse describes the structure of the API response when listing collections of agent objects










## <a name="AgentsService">type</a> [AgentsService](https://github.com/beamly/go-gocd/tree/master/gocd/agent.go?s=112:138#L9)
``` go
type AgentsService service
```
AgentsService describes actions which can be performed on agents










### <a name="AgentsService.BulkUpdate">func</a> (\*AgentsService) [BulkUpdate](https://github.com/beamly/go-gocd/tree/master/gocd/agent.go?s=3862:3988#L100)
``` go
func (s *AgentsService) BulkUpdate(ctx context.Context, agents AgentBulkUpdate) (message string, resp *APIResponse, err error)
```
BulkUpdate will change the configuration for multiple agents in a single request.




### <a name="AgentsService.Delete">func</a> (\*AgentsService) [Delete](https://github.com/beamly/go-gocd/tree/master/gocd/agent.go?s=3619:3713#L95)
``` go
func (s *AgentsService) Delete(ctx context.Context, uuid string) (string, *APIResponse, error)
```
Delete will remove an existing agent. Note: The agent must be disabled, and not currently building to be deleted.




### <a name="AgentsService.Get">func</a> (\*AgentsService) [Get](https://github.com/beamly/go-gocd/tree/master/gocd/agent.go?s=3119:3210#L85)
``` go
func (s *AgentsService) Get(ctx context.Context, uuid string) (*Agent, *APIResponse, error)
```
Get will retrieve a single agent based on the provided UUID.




### <a name="AgentsService.JobRunHistory">func</a> (\*AgentsService) [JobRunHistory](https://github.com/beamly/go-gocd/tree/master/gocd/agent.go?s=4287:4402#L113)
``` go
func (s *AgentsService) JobRunHistory(ctx context.Context, uuid string) (jobs []*Job, resp *APIResponse, err error)
```
JobRunHistory will return a list of Jobs run on the agent identified by `uuid`.




### <a name="AgentsService.List">func</a> (\*AgentsService) [List](https://github.com/beamly/go-gocd/tree/master/gocd/agent.go?s=2687:2784#L68)
``` go
func (s *AgentsService) List(ctx context.Context) (agents []*Agent, resp *APIResponse, err error)
```
List will retrieve all agents, their status, and metadata from the GoCD Server.




### <a name="AgentsService.Update">func</a> (\*AgentsService) [Update](https://github.com/beamly/go-gocd/tree/master/gocd/agent.go?s=3332:3440#L90)
``` go
func (s *AgentsService) Update(ctx context.Context, uuid string, agent *Agent) (*Agent, *APIResponse, error)
```
Update will modify the configuration for an existing agents.




## <a name="Approval">type</a> [Approval](https://github.com/beamly/go-gocd/tree/master/gocd/approval.go?s=116:257#L4)
``` go
type Approval struct {
    Type          string         `json:"type,omitempty"`
    Authorization *Authorization `json:"authorization,omitempty"`
}
```
Approval represents a request/response object describing the approval configuration for a GoCD Job










### <a name="Approval.Clean">func</a> (\*Approval) [Clean](https://github.com/beamly/go-gocd/tree/master/gocd/resource_approval.go?s=113:139#L5)
``` go
func (a *Approval) Clean()
```
Clean ensures integrity of the schema by making sure
empty elements are not printed to json.




## <a name="Artifact">type</a> [Artifact](https://github.com/beamly/go-gocd/tree/master/gocd/jobs.go?s=2072:2207#L44)
``` go
type Artifact struct {
    Type        string `json:"type"`
    Source      string `json:"source"`
    Destination string `json:"destination"`
}
```
Artifact describes the result of a job










## <a name="Auth">type</a> [Auth](https://github.com/beamly/go-gocd/tree/master/gocd/gocd.go?s=2801:2855#L114)
``` go
type Auth struct {
    Username string
    Password string
}
```
Auth structure wrapping the Username and Password variables, which are used to get an Auth cookie header used for
subsequent requests.










## <a name="Authorization">type</a> [Authorization](https://github.com/beamly/go-gocd/tree/master/gocd/approval.go?s=431:542#L11)
``` go
type Authorization struct {
    Users []string `json:"users,omitempty"`
    Roles []string `json:"roles,omitempty"`
}
```
Authorization describes the access control for a "manual" approval type. Specifies who (role or users) can approve
the job to move to the next stage in the pipeline.










## <a name="BuildCause">type</a> [BuildCause](https://github.com/beamly/go-gocd/tree/master/gocd/pipeline.go?s=2791:3074#L77)
``` go
type BuildCause struct {
    Approver          string             `json:"approver,omitempty"`
    MaterialRevisions []MaterialRevision `json:"material_revisions"`
    TriggerForced     bool               `json:"trigger_forced"`
    TriggerMessage    string             `json:"trigger_message"`
}
```
BuildCause describes the triggers which caused the build to start.










## <a name="BuildDetails">type</a> [BuildDetails](https://github.com/beamly/go-gocd/tree/master/gocd/agent.go?s=2432:2602#L60)
``` go
type BuildDetails struct {
    Links    *HALLinks `json:"_links"`
    Pipeline string    `json:"pipeline"`
    Stage    string    `json:"stage"`
    Job      string    `json:"job"`
}
```
BuildDetails describes the builds being performed on this agent.










## <a name="CipherText">type</a> [CipherText](https://github.com/beamly/go-gocd/tree/master/gocd/encryption.go?s=247:366#L11)
``` go
type CipherText struct {
    EncryptedValue string    `json:"encrypted_value"`
    Links          *HALLinks `json:"_links"`
}
```
CipherText sescribes the response from the api with an encrypted value.










## <a name="Client">type</a> [Client](https://github.com/beamly/go-gocd/tree/master/gocd/gocd.go?s=1552:2302#L69)
``` go
type Client struct {
    BaseURL  *url.URL
    Username string
    Password string

    UserAgent string

    Log *logrus.Logger

    Agents            *AgentsService
    PipelineGroups    *PipelineGroupsService
    Stages            *StagesService
    Jobs              *JobsService
    PipelineTemplates *PipelineTemplatesService
    Pipelines         *PipelinesService
    PipelineConfigs   *PipelineConfigsService
    Configuration     *ConfigurationService
    ConfigRepos       *ConfigRepoService
    Encryption        *EncryptionService
    Plugins           *PluginsService
    Environments      *EnvironmentsService
    Properties        *PropertiesService
    // contains filtered or unexported fields
}
```
Client struct which acts as an interface to the GoCD Server. Exposes resource service handlers.







### <a name="NewClient">func</a> [NewClient](https://github.com/beamly/go-gocd/tree/master/gocd/gocd.go?s=3355:3422#L131)
``` go
func NewClient(cfg *Configuration, httpClient *http.Client) *Client
```
NewClient creates a new client based on the provided configuration payload, and optionally a custom httpClient to
allow overriding of http client structures.





### <a name="Client.Do">func</a> (\*Client) [Do](https://github.com/beamly/go-gocd/tree/master/gocd/gocd.go?s=6506:6621#L255)
``` go
func (c *Client) Do(ctx context.Context, req *APIRequest, v interface{}, responseType string) (*APIResponse, error)
```
Do takes an HTTP request and resposne the response from the GoCD API endpoint.




### <a name="Client.Lock">func</a> (\*Client) [Lock](https://github.com/beamly/go-gocd/tree/master/gocd/gocd.go?s=4631:4654#L176)
``` go
func (c *Client) Lock()
```
Lock the client until release




### <a name="Client.Login">func</a> (\*Client) [Login](https://github.com/beamly/go-gocd/tree/master/gocd/authentication.go?s=167:222#L7)
``` go
func (c *Client) Login(ctx context.Context) (err error)
```
Login sends basic auth to the GoCD Server and sets an auth cookie in the client to enable cookie based auth
for future requests.




### <a name="Client.NewRequest">func</a> (\*Client) [NewRequest](https://github.com/beamly/go-gocd/tree/master/gocd/gocd.go?s=4838:4954#L186)
``` go
func (c *Client) NewRequest(method, urlStr string, body interface{}, apiVersion string) (req *APIRequest, err error)
```
NewRequest creates an HTTP requests to the GoCD API endpoints.




### <a name="Client.Unlock">func</a> (\*Client) [Unlock](https://github.com/beamly/go-gocd/tree/master/gocd/gocd.go?s=4720:4745#L181)
``` go
func (c *Client) Unlock()
```
Unlock the client after a lock action




## <a name="ConfigApproval">type</a> [ConfigApproval](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go?s=2573:2662#L60)
``` go
type ConfigApproval struct {
    Type string `xml:"type,attr,omitempty" json:",omitempty"`
}
```
ConfigApproval part of cruise-control.xml. @TODO better documentation










## <a name="ConfigArtifact">type</a> [ConfigArtifact](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go?s=2365:2498#L54)
``` go
type ConfigArtifact struct {
    Src         string `xml:"src,attr"`
    Destination string `xml:"dest,attr,omitempty" json:",omitempty"`
}
```
ConfigArtifact part of cruise-control.xml. @TODO better documentation










## <a name="ConfigAuthConfig">type</a> [ConfigAuthConfig](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go?s=7269:7443#L182)
``` go
type ConfigAuthConfig struct {
    ID         string           `xml:"id,attr"`
    PluginID   string           `xml:"pluginId,attr"`
    Properties []ConfigProperty `xml:"property"`
}
```
ConfigAuthConfig part of cruise-control.xml. @TODO better documentation










## <a name="ConfigElastic">type</a> [ConfigElastic](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go?s=7517:7604#L189)
``` go
type ConfigElastic struct {
    Profiles []ConfigElasticProfile `xml:"profiles>profile"`
}
```
ConfigElastic part of cruise-control.xml. @TODO better documentation










## <a name="ConfigElasticProfile">type</a> [ConfigElasticProfile](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go?s=7685:7863#L194)
``` go
type ConfigElasticProfile struct {
    ID         string           `xml:"id,attr"`
    PluginID   string           `xml:"pluginId,attr"`
    Properties []ConfigProperty `xml:"property"`
}
```
ConfigElasticProfile part of cruise-control.xml. @TODO better documentation










## <a name="ConfigEnvironmentVariable">type</a> [ConfigEnvironmentVariable](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go?s=2748:2849#L65)
``` go
type ConfigEnvironmentVariable struct {
    Name  string `xml:"name,attr"`
    Value string `xml:"value"`
}
```
ConfigEnvironmentVariable part of cruise-control.xml. @TODO better documentation










## <a name="ConfigFilter">type</a> [ConfigFilter](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go?s=3385:3459#L84)
``` go
type ConfigFilter struct {
    Ignore string `xml:"pattern,attr,omitempty"`
}
```
ConfigFilter part of cruise-control.xml. @TODO better documentation










## <a name="ConfigJob">type</a> [ConfigJob](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go?s=1837:2290#L45)
``` go
type ConfigJob struct {
    Name                 string                      `xml:"name,attr"`
    EnvironmentVariables []ConfigEnvironmentVariable `xml:"environmentvariables>variable" json:",omitempty"`
    Tasks                ConfigTasks                 `xml:"tasks"`
    Resources            []string                    `xml:"resources>resource" json:",omitempty"`
    Artifacts            []ConfigArtifact            `xml:"artifacts>artifact" json:",omitempty"`
}
```
ConfigJob part of cruise-control.xml. @TODO better documentation










## <a name="ConfigMaterialRepository">type</a> [ConfigMaterialRepository](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go?s=4468:4861#L115)
``` go
type ConfigMaterialRepository struct {
    ID                  string                    `xml:"id,attr"`
    Name                string                    `xml:"name,attr"`
    PluginConfiguration ConfigPluginConfiguration `xml:"pluginConfiguration"`
    Configuration       []ConfigProperty          `xml:"configuration>property"`
    Packages            []ConfigPackage           `xml:"packages>package"`
}
```
ConfigMaterialRepository part of cruise-control.xml. @TODO better documentation










## <a name="ConfigPackage">type</a> [ConfigPackage](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go?s=4935:5125#L124)
``` go
type ConfigPackage struct {
    ID            string           `xml:"id,attr"`
    Name          string           `xml:"name,attr"`
    Configuration []ConfigProperty `xml:"configuration>property"`
}
```
ConfigPackage part of cruise-control.xml. @TODO better documentation










## <a name="ConfigParam">type</a> [ConfigParam](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go?s=3531:3622#L89)
``` go
type ConfigParam struct {
    Name  string `xml:"name,attr"`
    Value string `xml:",chardata"`
}
```
ConfigParam part of cruise-control.xml. @TODO better documentation










## <a name="ConfigPipeline">type</a> [ConfigPipeline](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go?s=882:1513#L26)
``` go
type ConfigPipeline struct {
    Name                 string                      `xml:"name,attr"`
    LabelTemplate        string                      `xml:"labeltemplate,attr"`
    Params               []ConfigParam               `xml:"params>param"`
    GitMaterials         []GitRepositoryMaterial     `xml:"materials>git,omitempty"`
    PipelineMaterials    []PipelineMaterial          `xml:"materials>pipeline,omitempty"`
    Timer                string                      `xml:"timer"`
    EnvironmentVariables []ConfigEnvironmentVariable `xml:"environmentvariables>variable"`
    Stages               []ConfigStage               `xml:"stage"`
}
```
ConfigPipeline part of cruise-control.xml. @TODO better documentation










## <a name="ConfigPipelineGroup">type</a> [ConfigPipelineGroup](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go?s=680:807#L20)
``` go
type ConfigPipelineGroup struct {
    Name      string           `xml:"group,attr"`
    Pipelines []ConfigPipeline `xml:"pipeline"`
}
```
ConfigPipelineGroup contains a single pipeline groups










## <a name="ConfigPluginConfiguration">type</a> [ConfigPluginConfiguration](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go?s=5211:5321#L131)
``` go
type ConfigPluginConfiguration struct {
    ID      string `xml:"id,attr"`
    Version string `xml:"version,attr"`
}
```
ConfigPluginConfiguration part of cruise-control.xml. @TODO better documentation










## <a name="ConfigProperty">type</a> [ConfigProperty](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go?s=7938:8022#L201)
``` go
type ConfigProperty struct {
    Key   string `xml:"key"`
    Value string `xml:"value"`
}
```
ConfigProperty part of cruise-control.xml. @TODO better documentation










## <a name="ConfigRepo">type</a> [ConfigRepo](https://github.com/beamly/go-gocd/tree/master/gocd/config_repo.go?s=646:1062#L22)
``` go
type ConfigRepo struct {
    ID            string                `json:"id"`
    PluginID      string                `json:"plugin_id"`
    Material      Material              `json:"material"`
    Configuration []*ConfigRepoProperty `json:"configuration,omitempty"`
    Links         *HALLinks             `json:"_links,omitempty,omitempty"`
    Version       string                `json:"version,omitempty"`
    // contains filtered or unexported fields
}
```
ConfigRepo represents a config repo object in GoCD










### <a name="ConfigRepo.GetVersion">func</a> (\*ConfigRepo) [GetVersion](https://github.com/beamly/go-gocd/tree/master/gocd/resource_config_repo.go?s=207:257#L9)
``` go
func (c *ConfigRepo) GetVersion() (version string)
```
GetVersion retrieves a version string for this config repo




### <a name="ConfigRepo.SetVersion">func</a> (\*ConfigRepo) [SetVersion](https://github.com/beamly/go-gocd/tree/master/gocd/resource_config_repo.go?s=71:118#L4)
``` go
func (c *ConfigRepo) SetVersion(version string)
```
SetVersion sets a version string for this config repo




## <a name="ConfigRepoProperty">type</a> [ConfigRepoProperty](https://github.com/beamly/go-gocd/tree/master/gocd/config_repo.go?s=1137:1313#L33)
``` go
type ConfigRepoProperty struct {
    Key            string `json:"key"`
    Value          string `json:"value,omitempty"`
    EncryptedValue string `json:"encrypted_value,omitempty"`
}
```
ConfigRepoProperty represents a configuration related to a ConfigRepo










## <a name="ConfigRepoService">type</a> [ConfigRepoService](https://github.com/beamly/go-gocd/tree/master/gocd/config_repo.go?s=259:289#L11)
``` go
type ConfigRepoService service
```
ConfigRepoService allows admin users to define and manage config repos using
which pipelines defined in external repositories can be included in GoCD,
thereby allowing users to have their Pipeline as code.










### <a name="ConfigRepoService.Create">func</a> (\*ConfigRepoService) [Create](https://github.com/beamly/go-gocd/tree/master/gocd/config_repo.go?s=2261:2382#L71)
``` go
func (crs *ConfigRepoService) Create(ctx context.Context, cr *ConfigRepo) (out *ConfigRepo, resp *APIResponse, err error)
```
Create a config repo




### <a name="ConfigRepoService.Delete">func</a> (\*ConfigRepoService) [Delete](https://github.com/beamly/go-gocd/tree/master/gocd/config_repo.go?s=3079:3177#L99)
``` go
func (crs *ConfigRepoService) Delete(ctx context.Context, id string) (string, *APIResponse, error)
```
Delete the specified config repo




### <a name="ConfigRepoService.Get">func</a> (\*ConfigRepoService) [Get](https://github.com/beamly/go-gocd/tree/master/gocd/config_repo.go?s=1896:2009#L58)
``` go
func (crs *ConfigRepoService) Get(ctx context.Context, id string) (out *ConfigRepo, resp *APIResponse, err error)
```
Get fetches the config repo object for a specified id




### <a name="ConfigRepoService.List">func</a> (\*ConfigRepoService) [List](https://github.com/beamly/go-gocd/tree/master/gocd/config_repo.go?s=1439:1546#L41)
``` go
func (crs *ConfigRepoService) List(ctx context.Context) (repos []*ConfigRepo, resp *APIResponse, err error)
```
List returns all available config repos, these are config repositories that
are present in the in `cruise-config.xml`




### <a name="ConfigRepoService.Update">func</a> (\*ConfigRepoService) [Update](https://github.com/beamly/go-gocd/tree/master/gocd/config_repo.go?s=2663:2795#L85)
``` go
func (crs *ConfigRepoService) Update(ctx context.Context, id string, cr *ConfigRepo) (out *ConfigRepo, resp *APIResponse, err error)
```
Update config repos for specified config repo id




## <a name="ConfigReposListResponse">type</a> [ConfigReposListResponse](https://github.com/beamly/go-gocd/tree/master/gocd/config_repo.go?s=409:590#L14)
``` go
type ConfigReposListResponse struct {
    Links    *HALLinks `json:"_links,omitempty"`
    Embedded *struct {
        Repos []*ConfigRepo `json:"config_repos"`
    } `json:"_embedded,omitempty"`
}
```
ConfigReposListResponse describes the structure of the API response when listing collections of ConfigRepo objects










## <a name="ConfigRepository">type</a> [ConfigRepository](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go?s=3699:3863#L95)
``` go
type ConfigRepository struct {
    Plugin string              `xml:"plugin,attr"`
    ID     string              `xml:"id,attr"`
    Git    ConfigRepositoryGit `xml:"git"`
}
```
ConfigRepository part of cruise-control.xml. @TODO better documentation










## <a name="ConfigRepositoryGit">type</a> [ConfigRepositoryGit](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go?s=3943:4007#L102)
``` go
type ConfigRepositoryGit struct {
    URL string `xml:"url,attr"`
}
```
ConfigRepositoryGit part of cruise-control.xml. @TODO better documentation










## <a name="ConfigRole">type</a> [ConfigRole](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go?s=7097:7192#L176)
``` go
type ConfigRole struct {
    Name  string   `xml:"name,attr"`
    Users []string `xml:"users>user"`
}
```
ConfigRole part of cruise-control.xml. @TODO better documentation










## <a name="ConfigSCM">type</a> [ConfigSCM](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go?s=4077:4383#L107)
``` go
type ConfigSCM struct {
    ID                  string                    `xml:"id,attr"`
    Name                string                    `xml:"name,attr"`
    PluginConfiguration ConfigPluginConfiguration `xml:"pluginConfiguration"`
    Configuration       []ConfigProperty          `xml:"configuration>property"`
}
```
ConfigSCM part of cruise-control.xml. @TODO better documentation










## <a name="ConfigSecurity">type</a> [ConfigSecurity](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go?s=6632:6885#L163)
``` go
type ConfigSecurity struct {
    AuthConfigs  []ConfigAuthConfig `xml:"authConfigs>authConfig"`
    Roles        []ConfigRole       `xml:"roles>role"`
    Admins       []string           `xml:"admins>user"`
    PasswordFile PasswordFilePath   `xml:"passwordFile"`
}
```
ConfigSecurity part of cruise-control.xml. @TODO better documentation










## <a name="ConfigServer">type</a> [ConfigServer](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go?s=5394:6285#L137)
``` go
type ConfigServer struct {
    MailHost                  MailHost       `xml:"mailhost"`
    Security                  ConfigSecurity `xml:"security"`
    Elastic                   ConfigElastic  `xml:"elastic"`
    ArtifactsDir              string         `xml:"artifactsdir,attr"`
    SiteURL                   string         `xml:"siteUrl,attr"`
    SecureSiteURL             string         `xml:"secureSiteUrl,attr"`
    PurgeStart                string         `xml:"purgeStart,attr"`
    PurgeUpTo                 string         `xml:"purgeUpto,attr"`
    JobTimeout                int            `xml:"jobTimeout,attr"`
    AgentAutoRegisterKey      string         `xml:"agentAutoRegisterKey,attr"`
    WebhookSecret             string         `xml:"webhookSecret,attr"`
    CommandRepositoryLocation string         `xml:"commandRepositoryLocation,attr"`
    ServerID                  string         `xml:"serverId,attr"`
}
```
ConfigServer part of cruise-control.xml. @TODO better documentation










## <a name="ConfigStage">type</a> [ConfigStage](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go?s=1585:1767#L38)
``` go
type ConfigStage struct {
    Name     string         `xml:"name,attr"`
    Approval ConfigApproval `xml:"approval,omitempty" json:",omitempty"`
    Jobs     []ConfigJob    `xml:"jobs>job"`
}
```
ConfigStage part of cruise-control.xml. @TODO better documentation










## <a name="ConfigTask">type</a> [ConfigTask](https://github.com/beamly/go-gocd/tree/master/gocd/configuration_task.go?s=243:1076#L13)
``` go
type ConfigTask struct {
    // Because we need to preserve the order of tasks, and we have an array of elements with mixed types,
    // we need to use this generic xml type for tasks.
    XMLName  xml.Name        `json:",omitempty"`
    Type     string          `xml:"type,omitempty"`
    RunIf    ConfigTaskRunIf `xml:"runif"`
    Command  string          `xml:"command,attr,omitempty"  json:",omitempty"`
    Args     []string        `xml:"arg,omitempty"  json:",omitempty"`
    Pipeline string          `xml:"pipeline,attr,omitempty"  json:",omitempty"`
    Stage    string          `xml:"stage,attr,omitempty"  json:",omitempty"`
    Job      string          `xml:"job,attr,omitempty"  json:",omitempty"`
    SrcFile  string          `xml:"srcfile,attr,omitempty"  json:",omitempty"`
    SrcDir   string          `xml:"srcdir,attr,omitempty"  json:",omitempty"`
}
```
ConfigTask part of cruise-control.xml. @TODO better documentation










## <a name="ConfigTaskRunIf">type</a> [ConfigTaskRunIf](https://github.com/beamly/go-gocd/tree/master/gocd/configuration_task.go?s=1152:1218#L29)
``` go
type ConfigTaskRunIf struct {
    Status string `xml:"status,attr"`
}
```
ConfigTaskRunIf part of cruise-control.xml. @TODO better documentation










## <a name="ConfigTasks">type</a> [ConfigTasks](https://github.com/beamly/go-gocd/tree/master/gocd/configuration_task.go?s=112:172#L8)
``` go
type ConfigTasks struct {
    Tasks []ConfigTask `xml:",any"`
}
```
ConfigTasks part of cruise-control.xml. @TODO better documentation










## <a name="ConfigXML">type</a> [ConfigXML](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go?s=246:621#L11)
``` go
type ConfigXML struct {
    Repositories       []ConfigMaterialRepository `xml:"repositories>repository"`
    Server             ConfigServer               `xml:"server"`
    SCMS               []ConfigSCM                `xml:"scms>scm"`
    ConfigRepositories []ConfigRepository         `xml:"config-repos>config-repo"`
    PipelineGroups     []ConfigPipelineGroup      `xml:"pipelines"`
}
```
ConfigXML part of cruise-control.xml. @TODO better documentation










## <a name="Configuration">type</a> [Configuration](https://github.com/beamly/go-gocd/tree/master/gocd/config.go?s=554:781#L26)
``` go
type Configuration struct {
    Server       string
    Username     string `yaml:"username,omitempty"`
    Password     string `yaml:"password,omitempty"`
    SkipSslCheck bool   `yaml:"skip_ssl_check,omitempty" survey:"skip_ssl_check"`
}
```
Configuration describes a single connection to a GoCD server










### <a name="Configuration.Client">func</a> (\*Configuration) [Client](https://github.com/beamly/go-gocd/tree/master/gocd/gocd.go?s=3119:3159#L125)
``` go
func (c *Configuration) Client() *Client
```
Client returns a client which allows us to interact with the GoCD Server.




### <a name="Configuration.HasAuth">func</a> (\*Configuration) [HasAuth](https://github.com/beamly/go-gocd/tree/master/gocd/gocd.go?s=2949:2987#L120)
``` go
func (c *Configuration) HasAuth() bool
```
HasAuth checks whether or not we have the required Username/Password variables provided.




## <a name="ConfigurationService">type</a> [ConfigurationService](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go?s=143:176#L8)
``` go
type ConfigurationService service
```
ConfigurationService describes the HAL _link resource for the api response object for a pipelineconfig










### <a name="ConfigurationService.Get">func</a> (\*ConfigurationService) [Get](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go?s=8459:8561#L217)
``` go
func (cs *ConfigurationService) Get(ctx context.Context) (cx *ConfigXML, resp *APIResponse, err error)
```
Get the config.xml document from the server and... render it as JSON... 'cause... eyugh.




### <a name="ConfigurationService.GetVersion">func</a> (\*ConfigurationService) [GetVersion](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go?s=8827:8933#L228)
``` go
func (cs *ConfigurationService) GetVersion(ctx context.Context) (v *Version, resp *APIResponse, err error)
```
GetVersion of the GoCD server and other metadata about the software version.




## <a name="EmbeddedEnvironments">type</a> [EmbeddedEnvironments](https://github.com/beamly/go-gocd/tree/master/gocd/environment.go?s=440:527#L17)
``` go
type EmbeddedEnvironments struct {
    Environments []*Environment `json:"environments"`
}
```
EmbeddedEnvironments encapsulates the environment struct










## <a name="EncryptionService">type</a> [EncryptionService](https://github.com/beamly/go-gocd/tree/master/gocd/encryption.go?s=140:170#L8)
``` go
type EncryptionService service
```
EncryptionService describes the HAL _link resource for the api response object for a pipelineconfig










### <a name="EncryptionService.Encrypt">func</a> (\*EncryptionService) [Encrypt](https://github.com/beamly/go-gocd/tree/master/gocd/encryption.go?s=430:551#L17)
``` go
func (es *EncryptionService) Encrypt(ctx context.Context, plaintext string) (c *CipherText, resp *APIResponse, err error)
```
Encrypt takes a plaintext value and returns a cipher text.




## <a name="Environment">type</a> [Environment](https://github.com/beamly/go-gocd/tree/master/gocd/environment.go?s=586:1036#L22)
``` go
type Environment struct {
    Links                *HALLinks              `json:"_links,omitempty"`
    Name                 string                 `json:"name"`
    Pipelines            []*Pipeline            `json:"pipelines,omitempty"`
    Agents               []*Agent               `json:"agents,omitempty"`
    EnvironmentVariables []*EnvironmentVariable `json:"environment_variables,omitempty"`
    Version              string                 `json:"version"`
}
```
Environment describes a group of pipelines and agents










### <a name="Environment.GetLinks">func</a> (\*Environment) [GetLinks](https://github.com/beamly/go-gocd/tree/master/gocd/resource_environment.go?s=629:673#L28)
``` go
func (env *Environment) GetLinks() *HALLinks
```
GetLinks from the Environment




### <a name="Environment.GetVersion">func</a> (\*Environment) [GetVersion](https://github.com/beamly/go-gocd/tree/master/gocd/resource_environment.go?s=889:942#L38)
``` go
func (env *Environment) GetVersion() (version string)
```
GetVersion retrieves a version string for this pipeline




### <a name="Environment.RemoveLinks">func</a> (\*Environment) [RemoveLinks](https://github.com/beamly/go-gocd/tree/master/gocd/resource_environment.go?s=427:464#L17)
``` go
func (env *Environment) RemoveLinks()
```
RemoveLinks gets the Environment ready to be submitted to the GoCD API.




### <a name="Environment.SetVersion">func</a> (\*Environment) [SetVersion](https://github.com/beamly/go-gocd/tree/master/gocd/resource_environment.go?s=751:801#L33)
``` go
func (env *Environment) SetVersion(version string)
```
SetVersion sets a version string for this pipeline




## <a name="EnvironmentPatchRequest">type</a> [EnvironmentPatchRequest](https://github.com/beamly/go-gocd/tree/master/gocd/environment.go?s=1116:1371#L32)
``` go
type EnvironmentPatchRequest struct {
    Pipelines            *PatchStringAction          `json:"pipelines"`
    Agents               *PatchStringAction          `json:"agents"`
    EnvironmentVariables *EnvironmentVariablesAction `json:"environment_variables"`
}
```
EnvironmentPatchRequest describes the actions to perform on an environment










## <a name="EnvironmentVariable">type</a> [EnvironmentVariable](https://github.com/beamly/go-gocd/tree/master/gocd/jobs.go?s=2551:2768#L64)
``` go
type EnvironmentVariable struct {
    Name           string `json:"name"`
    Value          string `json:"value,omitempty"`
    EncryptedValue string `json:"encrypted_value,omitempty"`
    Secure         bool   `json:"secure"`
}
```
EnvironmentVariable describes an environment variable key/pair.










## <a name="EnvironmentVariablesAction">type</a> [EnvironmentVariablesAction](https://github.com/beamly/go-gocd/tree/master/gocd/environment.go?s=1469:1602#L39)
``` go
type EnvironmentVariablesAction struct {
    Add    []*EnvironmentVariable `json:"add"`
    Remove []*EnvironmentVariable `json:"remove"`
}
```
EnvironmentVariablesAction describes a collection of Environment Variables to add or remove.










## <a name="EnvironmentsResponse">type</a> [EnvironmentsResponse](https://github.com/beamly/go-gocd/tree/master/gocd/environment.go?s=243:378#L11)
``` go
type EnvironmentsResponse struct {
    Links    *HALLinks             `json:"_links"`
    Embedded *EmbeddedEnvironments `json:"_embedded"`
}
```
EnvironmentsResponse describes the response obejct for a plugin API call.










### <a name="EnvironmentsResponse.GetLinks">func</a> (\*EnvironmentsResponse) [GetLinks](https://github.com/beamly/go-gocd/tree/master/gocd/resource_environment.go?s=277:329#L12)
``` go
func (er *EnvironmentsResponse) GetLinks() *HALLinks
```
GetLinks from the EnvironmentResponse




### <a name="EnvironmentsResponse.RemoveLinks">func</a> (\*EnvironmentsResponse) [RemoveLinks](https://github.com/beamly/go-gocd/tree/master/gocd/resource_environment.go?s=98:143#L4)
``` go
func (er *EnvironmentsResponse) RemoveLinks()
```
RemoveLinks gets the EnvironmentsResponse ready to be submitted to the GoCD API.




## <a name="EnvironmentsService">type</a> [EnvironmentsService](https://github.com/beamly/go-gocd/tree/master/gocd/environment.go?s=132:164#L8)
``` go
type EnvironmentsService service
```
EnvironmentsService exposes calls for interacting with Environment objects in the GoCD API.










### <a name="EnvironmentsService.Create">func</a> (\*EnvironmentsService) [Create](https://github.com/beamly/go-gocd/tree/master/gocd/environment.go?s=2330:2448#L68)
``` go
func (es *EnvironmentsService) Create(ctx context.Context, name string) (e *Environment, resp *APIResponse, err error)
```
Create an environment




### <a name="EnvironmentsService.Delete">func</a> (\*EnvironmentsService) [Delete](https://github.com/beamly/go-gocd/tree/master/gocd/environment.go?s=2127:2228#L63)
``` go
func (es *EnvironmentsService) Delete(ctx context.Context, name string) (string, *APIResponse, error)
```
Delete an environment




### <a name="EnvironmentsService.Get">func</a> (\*EnvironmentsService) [Get](https://github.com/beamly/go-gocd/tree/master/gocd/environment.go?s=2686:2801#L82)
``` go
func (es *EnvironmentsService) Get(ctx context.Context, name string) (e *Environment, resp *APIResponse, err error)
```
Get a single environment by name




### <a name="EnvironmentsService.List">func</a> (\*EnvironmentsService) [List](https://github.com/beamly/go-gocd/tree/master/gocd/environment.go?s=1802:1914#L51)
``` go
func (es *EnvironmentsService) List(ctx context.Context) (e *EnvironmentsResponse, resp *APIResponse, err error)
```
List all environments




### <a name="EnvironmentsService.Patch">func</a> (\*EnvironmentsService) [Patch](https://github.com/beamly/go-gocd/tree/master/gocd/environment.go?s=3090:3239#L94)
``` go
func (es *EnvironmentsService) Patch(ctx context.Context, name string, patch *EnvironmentPatchRequest) (e *Environment, resp *APIResponse, err error)
```
Patch an environments configuration by adding or removing pipelines, agents, environment variables




## <a name="GitRepositoryMaterial">type</a> [GitRepositoryMaterial](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go?s=3178:3312#L78)
``` go
type GitRepositoryMaterial struct {
    URL     string         `xml:"url,attr"`
    Filters []ConfigFilter `xml:"filter>ignore,omitempty"`
}
```
GitRepositoryMaterial part of cruise-control.xml. @TODO better documentation










## <a name="HALContainer">type</a> [HALContainer](https://github.com/beamly/go-gocd/tree/master/gocd/resource.go?s=387:455#L17)
``` go
type HALContainer interface {
    RemoveLinks()
    GetLinks() *HALLinks
}
```
HALContainer represents objects with HAL _link and _embedded resources.










## <a name="HALLink">type</a> [HALLink](https://github.com/beamly/go-gocd/tree/master/gocd/resource.go?s=632:683#L29)
``` go
type HALLink struct {
    Name string
    URL  *url.URL
}
```
HALLink describes a HAL link










## <a name="HALLinks">type</a> [HALLinks](https://github.com/beamly/go-gocd/tree/master/gocd/links.go?s=206:248#L15)
``` go
type HALLinks struct {
    // contains filtered or unexported fields
}
```
HALLinks describes a collection of HALLinks










### <a name="HALLinks.Add">func</a> (\*HALLinks) [Add](https://github.com/beamly/go-gocd/tree/master/gocd/links.go?s=264:302#L20)
``` go
func (al *HALLinks) Add(link *HALLink)
```
Add a link




### <a name="HALLinks.Get">func</a> (HALLinks) [Get](https://github.com/beamly/go-gocd/tree/master/gocd/links.go?s=368:419#L25)
``` go
func (al HALLinks) Get(name string) (link *HALLink)
```
Get a HALLink by name




### <a name="HALLinks.GetOk">func</a> (HALLinks) [GetOk](https://github.com/beamly/go-gocd/tree/master/gocd/links.go?s=525:587#L31)
``` go
func (al HALLinks) GetOk(name string) (link *HALLink, ok bool)
```
GetOk a HALLink by name, and if it doesn't exist, return false




### <a name="HALLinks.Keys">func</a> (HALLinks) [Keys](https://github.com/beamly/go-gocd/tree/master/gocd/links.go?s=778:819#L43)
``` go
func (al HALLinks) Keys() (keys []string)
```
Keys returns a string list of link names




### <a name="HALLinks.MarshallJSON">func</a> (HALLinks) [MarshallJSON](https://github.com/beamly/go-gocd/tree/master/gocd/links.go?s=978:1027#L52)
``` go
func (al HALLinks) MarshallJSON() ([]byte, error)
```
MarshallJSON allows the encoding of links into JSON




### <a name="HALLinks.UnmarshalJSON">func</a> (\*HALLinks) [UnmarshalJSON](https://github.com/beamly/go-gocd/tree/master/gocd/links.go?s=1230:1285#L61)
``` go
func (al *HALLinks) UnmarshalJSON(j []byte) (err error)
```
UnmarshalJSON allows the decoding of links from JSON




## <a name="Job">type</a> [Job](https://github.com/beamly/go-gocd/tree/master/gocd/jobs.go?s=354:2028#L18)
``` go
type Job struct {
    AgentUUID            string                 `json:"agent_uuid,omitempty"`
    Name                 string                 `json:"name"`
    JobStateTransitions  []*JobStateTransition  `json:"job_state_transitions,omitempty"`
    ScheduledDate        int                    `json:"scheduled_date,omitempty"`
    OriginalJobID        string                 `json:"original_job_id,omitempty"`
    PipelineCounter      int                    `json:"pipeline_counter,omitempty"`
    Rerun                bool                   `json:"rerun,omitempty"`
    PipelineName         string                 `json:"pipeline_name,omitempty"`
    Result               string                 `json:"result,omitempty"`
    State                string                 `json:"state,omitempty"`
    ID                   int                    `json:"id,omitempty"`
    StageCounter         string                 `json:"stage_counter,omitempty"`
    StageName            string                 `json:"stage_name,omitempty"`
    RunInstanceCount     int                    `json:"run_instance_count,omitempty"`
    Timeout              TimeoutField           `json:"timeout,omitempty"`
    EnvironmentVariables []*EnvironmentVariable `json:"environment_variables,omitempty"`
    Properties           []*JobProperty         `json:"properties,omitempty"`
    Resources            []string               `json:"resources,omitempty"`
    Tasks                []*Task                `json:"tasks,omitempty"`
    Tabs                 []*Tab                 `json:"tabs,omitempty"`
    Artifacts            []*Artifact            `json:"artifacts,omitempty"`
    ElasticProfileID     string                 `json:"elastic_profile_id,omitempty"`
}
```
Job describes a job which can be performed in GoCD










### <a name="Job.JSONString">func</a> (\*Job) [JSONString](https://github.com/beamly/go-gocd/tree/master/gocd/resource_jobs.go?s=127:178#L10)
``` go
func (j *Job) JSONString() (body string, err error)
```
JSONString returns a string of this stage as a JSON object.




### <a name="Job.Validate">func</a> (\*Job) [Validate](https://github.com/beamly/go-gocd/tree/master/gocd/resource_jobs.go?s=377:413#L23)
``` go
func (j *Job) Validate() (err error)
```
Validate a job structure has non-nil values on correct attributes




## <a name="JobProperty">type</a> [JobProperty](https://github.com/beamly/go-gocd/tree/master/gocd/jobs.go?s=2365:2482#L57)
``` go
type JobProperty struct {
    Name   string `json:"name"`
    Source string `json:"source"`
    XPath  string `json:"xpath"`
}
```
JobProperty describes the property for a job










## <a name="JobRunHistoryResponse">type</a> [JobRunHistoryResponse](https://github.com/beamly/go-gocd/tree/master/gocd/jobs.go?s=5357:5512#L129)
``` go
type JobRunHistoryResponse struct {
    Jobs       []*Job              `json:"jobs,omitempty"`
    Pagination *PaginationResponse `json:"pagination,omitempty"`
}
```
JobRunHistoryResponse describes the api response from










## <a name="JobSchedule">type</a> [JobSchedule](https://github.com/beamly/go-gocd/tree/master/gocd/jobs.go?s=5689:6117#L140)
``` go
type JobSchedule struct {
    Name                 string               `xml:"name,attr"`
    ID                   string               `xml:"id,attr"`
    Link                 JobScheduleLink      `xml:"link"`
    BuildLocator         string               `xml:"buildLocator"`
    Resources            []string             `xml:"resources>resource"`
    EnvironmentVariables *[]JobScheduleEnvVar `xml:"environmentVariables,omitempty>variable"`
}
```
JobSchedule describes the event causes for a job










## <a name="JobScheduleEnvVar">type</a> [JobScheduleEnvVar](https://github.com/beamly/go-gocd/tree/master/gocd/jobs.go?s=6195:6292#L150)
``` go
type JobScheduleEnvVar struct {
    Name  string `xml:"name,attr"`
    Value string `xml:",innerxml"`
}
```
JobScheduleEnvVar describes the environmnet variables for a job schedule










## <a name="JobScheduleLink">type</a> [JobScheduleLink](https://github.com/beamly/go-gocd/tree/master/gocd/jobs.go?s=6356:6448#L156)
``` go
type JobScheduleLink struct {
    Rel  string `xml:"rel,attr"`
    HRef string `xml:"href,attr"`
}
```
JobScheduleLink describes the HAL links for a job schedule










## <a name="JobScheduleResponse">type</a> [JobScheduleResponse](https://github.com/beamly/go-gocd/tree/master/gocd/jobs.go?s=5567:5635#L135)
``` go
type JobScheduleResponse struct {
    Jobs []*JobSchedule `xml:"job"`
}
```
JobScheduleResponse contains a collection of jobs










## <a name="JobStateTransition">type</a> [JobStateTransition](https://github.com/beamly/go-gocd/tree/master/gocd/jobs.go?s=5108:5298#L122)
``` go
type JobStateTransition struct {
    StateChangeTime int    `json:"state_change_time,omitempty"`
    ID              int    `json:"id,omitempty"`
    State           string `json:"state,omitempty"`
}
```
JobStateTransition describes a State Transition object in a GoCD api response










## <a name="JobsService">type</a> [JobsService](https://github.com/beamly/go-gocd/tree/master/gocd/jobs.go?s=274:298#L15)
``` go
type JobsService service
```
JobsService describes actions which can be performed on jobs










### <a name="JobsService.ListScheduled">func</a> (\*JobsService) [ListScheduled](https://github.com/beamly/go-gocd/tree/master/gocd/jobs.go?s=6619:6728#L165)
``` go
func (js *JobsService) ListScheduled(ctx context.Context) (jobs []*JobSchedule, resp *APIResponse, err error)
```
ListScheduled lists Pipeline groups




## <a name="MailHost">type</a> [MailHost](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go?s=6354:6557#L154)
``` go
type MailHost struct {
    Hostname string `xml:"hostname,attr"`
    Port     int    `xml:"port,attr"`
    TLS      bool   `xml:"tls,attr"`
    From     string `xml:"from,attr"`
    Admin    string `xml:"admin,attr"`
}
```
MailHost part of cruise-control.xml. @TODO better documentation










## <a name="Material">type</a> [Material](https://github.com/beamly/go-gocd/tree/master/gocd/pipeline.go?s=1869:2113#L49)
``` go
type Material struct {
    Type        string            `json:"type"`
    Fingerprint string            `json:"fingerprint,omitempty"`
    Description string            `json:"description,omitempty"`
    Attributes  MaterialAttribute `json:"attributes"`
}
```
Material describes an artifact dependency for a pipeline object.










### <a name="Material.Equal">func</a> (Material) [Equal](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material.go?s=158:220#L10)
``` go
func (m Material) Equal(a *Material) (isEqual bool, err error)
```
Equal is true if the two materials are logically equivalent. Not neccesarily literally equal.




### <a name="Material.Ingest">func</a> (\*Material) [Ingest](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material.go?s=557:626#L29)
``` go
func (m *Material) Ingest(payload map[string]interface{}) (err error)
```
Ingest an abstract structure




### <a name="Material.IngestAttributeGenerics">func</a> (\*Material) [IngestAttributeGenerics](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material.go?s=1360:1429#L65)
``` go
func (m *Material) IngestAttributeGenerics(i interface{}) (err error)
```
IngestAttributeGenerics to Material and perform some error checking




### <a name="Material.IngestAttributes">func</a> (\*Material) [IngestAttributes](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material.go?s=1585:1670#L73)
``` go
func (m *Material) IngestAttributes(rawAttributes map[string]interface{}) (err error)
```
IngestAttributes to Material from an abstract structure




### <a name="Material.IngestType">func</a> (\*Material) [IngestType](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material.go?s=1095:1156#L57)
``` go
func (m *Material) IngestType(payload map[string]interface{})
```
IngestType of Material if it is provided




### <a name="Material.UnmarshalJSON">func</a> (\*Material) [UnmarshalJSON](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material.go?s=350:404#L19)
``` go
func (m *Material) UnmarshalJSON(b []byte) (err error)
```
UnmarshalJSON string into a Material struct




## <a name="MaterialAttribute">type</a> [MaterialAttribute](https://github.com/beamly/go-gocd/tree/master/gocd/pipeline_material.go?s=106:282#L4)
``` go
type MaterialAttribute interface {
    GenerateGeneric() map[string]interface{}
    HasFilter() bool
    GetFilter() *MaterialFilter
    // contains filtered or unexported methods
}
```
MaterialAttribute describes the behaviour of the GoCD material structures for a pipeline










## <a name="MaterialAttributesDependency">type</a> [MaterialAttributesDependency](https://github.com/beamly/go-gocd/tree/master/gocd/pipeline_material.go?s=3140:3348#L89)
``` go
type MaterialAttributesDependency struct {
    Name       string `json:"name,omitempty"`
    Pipeline   string `json:"pipeline"`
    Stage      string `json:"stage"`
    AutoUpdate bool   `json:"auto_update,omitempty"`
}
```
MaterialAttributesDependency describes a Pipeline dependency material










### <a name="MaterialAttributesDependency.GenerateGeneric">func</a> (MaterialAttributesDependency) [GenerateGeneric](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_dependency.go?s=418:503#L19)
``` go
func (mad MaterialAttributesDependency) GenerateGeneric() (ma map[string]interface{})
```
GenerateGeneric form (map[string]interface) of the material filter




### <a name="MaterialAttributesDependency.GetFilter">func</a> (MaterialAttributesDependency) [GetFilter](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_dependency.go?s=928:995#L45)
``` go
func (mad MaterialAttributesDependency) GetFilter() *MaterialFilter
```
GetFilter from material attribute




### <a name="MaterialAttributesDependency.HasFilter">func</a> (MaterialAttributesDependency) [HasFilter](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_dependency.go?s=815:871#L40)
``` go
func (mad MaterialAttributesDependency) HasFilter() bool
```
HasFilter in this material attribute




## <a name="MaterialAttributesGit">type</a> [MaterialAttributesGit](https://github.com/beamly/go-gocd/tree/master/gocd/pipeline_material.go?s=334:839#L12)
``` go
type MaterialAttributesGit struct {
    Name   string `json:"name,omitempty"`
    URL    string `json:"url,omitempty"`
    Branch string `json:"branch,omitempty"`

    SubmoduleFolder string `json:"submodule_folder,omitempty"`
    ShallowClone    bool   `json:"shallow_clone,omitempty"`

    Destination  string          `json:"destination,omitempty"`
    Filter       *MaterialFilter `json:"filter,omitempty"`
    InvertFilter bool            `json:"invert_filter"`
    AutoUpdate   bool            `json:"auto_update,omitempty"`
}
```
MaterialAttributesGit describes a git material










### <a name="MaterialAttributesGit.GenerateGeneric">func</a> (MaterialAttributesGit) [GenerateGeneric](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_git.go?s=559:637#L25)
``` go
func (mag MaterialAttributesGit) GenerateGeneric() (ma map[string]interface{})
```
GenerateGeneric form (map[string]interface) of the material filter




### <a name="MaterialAttributesGit.GetFilter">func</a> (MaterialAttributesGit) [GetFilter](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_git.go?s=1198:1258#L48)
``` go
func (mag MaterialAttributesGit) GetFilter() *MaterialFilter
```
GetFilter from material attribute




### <a name="MaterialAttributesGit.HasFilter">func</a> (MaterialAttributesGit) [HasFilter](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_git.go?s=1093:1142#L43)
``` go
func (mag MaterialAttributesGit) HasFilter() bool
```
HasFilter in this material attribute




## <a name="MaterialAttributesHg">type</a> [MaterialAttributesHg](https://github.com/beamly/go-gocd/tree/master/gocd/pipeline_material.go?s=1511:1832#L43)
``` go
type MaterialAttributesHg struct {
    Name string `json:"name,omitempty"`
    URL  string `json:"url"`

    Destination  string          `json:"destination"`
    Filter       *MaterialFilter `json:"filter,omitempty"`
    InvertFilter bool            `json:"invert_filter"`
    AutoUpdate   bool            `json:"auto_update,omitempty"`
}
```
MaterialAttributesHg describes a Mercurial material type










### <a name="MaterialAttributesHg.GenerateGeneric">func</a> (MaterialAttributesHg) [GenerateGeneric](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_hg.go?s=443:520#L18)
``` go
func (mhg MaterialAttributesHg) GenerateGeneric() (ma map[string]interface{})
```
GenerateGeneric form (map[string]interface) of the material filter




### <a name="MaterialAttributesHg.GetFilter">func</a> (MaterialAttributesHg) [GetFilter](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_hg.go?s=713:772#L29)
``` go
func (mhg MaterialAttributesHg) GetFilter() *MaterialFilter
```
GetFilter from material attribute




### <a name="MaterialAttributesHg.HasFilter">func</a> (MaterialAttributesHg) [HasFilter](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_hg.go?s=609:657#L24)
``` go
func (mhg MaterialAttributesHg) HasFilter() bool
```
HasFilter in this material attribute




## <a name="MaterialAttributesP4">type</a> [MaterialAttributesP4](https://github.com/beamly/go-gocd/tree/master/gocd/pipeline_material.go?s=1893:2443#L54)
``` go
type MaterialAttributesP4 struct {
    Name       string `json:"name,omitempty"`
    Port       string `json:"port"`
    UseTickets bool   `json:"use_tickets"`
    View       string `json:"view"`

    Username          string `json:"username"`
    Password          string `json:"password"`
    EncryptedPassword string `json:"encrypted_password"`

    Destination  string          `json:"destination"`
    Filter       *MaterialFilter `json:"filter,omitempty"`
    InvertFilter bool            `json:"invert_filter"`
    AutoUpdate   bool            `json:"auto_update,omitempty"`
}
```
MaterialAttributesP4 describes a Perforce material type










### <a name="MaterialAttributesP4.GenerateGeneric">func</a> (MaterialAttributesP4) [GenerateGeneric](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_p4.go?s=491:568#L20)
``` go
func (mp4 MaterialAttributesP4) GenerateGeneric() (ma map[string]interface{})
```
GenerateGeneric form (map[string]interface) of the material filter




### <a name="MaterialAttributesP4.GetFilter">func</a> (MaterialAttributesP4) [GetFilter](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_p4.go?s=761:820#L31)
``` go
func (mp4 MaterialAttributesP4) GetFilter() *MaterialFilter
```
GetFilter from material attribute




### <a name="MaterialAttributesP4.HasFilter">func</a> (MaterialAttributesP4) [HasFilter](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_p4.go?s=657:705#L26)
``` go
func (mp4 MaterialAttributesP4) HasFilter() bool
```
HasFilter in this material attribute




## <a name="MaterialAttributesPackage">type</a> [MaterialAttributesPackage](https://github.com/beamly/go-gocd/tree/master/gocd/pipeline_material.go?s=3409:3475#L97)
``` go
type MaterialAttributesPackage struct {
    Ref string `json:"ref"`
}
```
MaterialAttributesPackage describes a package reference










### <a name="MaterialAttributesPackage.GenerateGeneric">func</a> (MaterialAttributesPackage) [GenerateGeneric](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_pkg.go?s=372:455#L16)
``` go
func (mapk MaterialAttributesPackage) GenerateGeneric() (ma map[string]interface{})
```
GenerateGeneric form (map[string]interface) of the material filter




### <a name="MaterialAttributesPackage.GetFilter">func</a> (MaterialAttributesPackage) [GetFilter](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_pkg.go?s=655:720#L27)
``` go
func (mapk MaterialAttributesPackage) GetFilter() *MaterialFilter
```
GetFilter from material attribute




### <a name="MaterialAttributesPackage.HasFilter">func</a> (MaterialAttributesPackage) [HasFilter](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_pkg.go?s=544:598#L22)
``` go
func (mapk MaterialAttributesPackage) HasFilter() bool
```
HasFilter in this material attribute




## <a name="MaterialAttributesPlugin">type</a> [MaterialAttributesPlugin](https://github.com/beamly/go-gocd/tree/master/gocd/pipeline_material.go?s=3533:3759#L102)
``` go
type MaterialAttributesPlugin struct {
    Ref string `json:"ref"`

    Destination  string          `json:"destination"`
    Filter       *MaterialFilter `json:"filter,omitempty"`
    InvertFilter bool            `json:"invert_filter"`
}
```
MaterialAttributesPlugin describes a plugin material










### <a name="MaterialAttributesPlugin.GenerateGeneric">func</a> (MaterialAttributesPlugin) [GenerateGeneric](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_plugin.go?s=416:498#L18)
``` go
func (mapp MaterialAttributesPlugin) GenerateGeneric() (ma map[string]interface{})
```
GenerateGeneric form (map[string]interface) of the material filter




### <a name="MaterialAttributesPlugin.GetFilter">func</a> (MaterialAttributesPlugin) [GetFilter](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_plugin.go?s=696:760#L29)
``` go
func (mapp MaterialAttributesPlugin) GetFilter() *MaterialFilter
```
GetFilter from material attribute




### <a name="MaterialAttributesPlugin.HasFilter">func</a> (MaterialAttributesPlugin) [HasFilter](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_plugin.go?s=587:640#L24)
``` go
func (mapp MaterialAttributesPlugin) HasFilter() bool
```
HasFilter in this material attribute




## <a name="MaterialAttributesSvn">type</a> [MaterialAttributesSvn](https://github.com/beamly/go-gocd/tree/master/gocd/pipeline_material.go?s=892:1449#L27)
``` go
type MaterialAttributesSvn struct {
    Name              string `json:"name,omitempty"`
    URL               string `json:"url,omitempty"`
    Username          string `json:"username"`
    Password          string `json:"password"`
    EncryptedPassword string `json:"encrypted_password"`

    CheckExternals bool `json:"check_externals"`

    Destination  string          `json:"destination,omitempty"`
    Filter       *MaterialFilter `json:"filter,omitempty"`
    InvertFilter bool            `json:"invert_filter"`
    AutoUpdate   bool            `json:"auto_update,omitempty"`
}
```
MaterialAttributesSvn describes a material type










### <a name="MaterialAttributesSvn.GenerateGeneric">func</a> (MaterialAttributesSvn) [GenerateGeneric](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_svn.go?s=471:549#L18)
``` go
func (mas MaterialAttributesSvn) GenerateGeneric() (ma map[string]interface{})
```
GenerateGeneric form (map[string]interface) of the material filter




### <a name="MaterialAttributesSvn.GetFilter">func</a> (MaterialAttributesSvn) [GetFilter](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_svn.go?s=743:803#L29)
``` go
func (mas MaterialAttributesSvn) GetFilter() *MaterialFilter
```
GetFilter from material attribute




### <a name="MaterialAttributesSvn.HasFilter">func</a> (MaterialAttributesSvn) [HasFilter](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_svn.go?s=638:687#L24)
``` go
func (mas MaterialAttributesSvn) HasFilter() bool
```
HasFilter in this material attribute




## <a name="MaterialAttributesTfs">type</a> [MaterialAttributesTfs](https://github.com/beamly/go-gocd/tree/master/gocd/pipeline_material.go?s=2514:3065#L71)
``` go
type MaterialAttributesTfs struct {
    Name string `json:"name,omitempty"`

    URL         string `json:"url"`
    ProjectPath string `json:"project_path"`
    Domain      string `json:"domain"`

    Username          string `json:"username"`
    Password          string `json:"password"`
    EncryptedPassword string `json:"encrypted_password"`

    Destination  string          `json:"destination"`
    Filter       *MaterialFilter `json:"filter,omitempty"`
    InvertFilter bool            `json:"invert_filter"`
    AutoUpdate   bool            `json:"auto_update,omitempty"`
}
```
MaterialAttributesTfs describes a Team Foundation Server material










### <a name="MaterialAttributesTfs.GenerateGeneric">func</a> (MaterialAttributesTfs) [GenerateGeneric](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_tfs.go?s=659:738#L28)
``` go
func (mtfs MaterialAttributesTfs) GenerateGeneric() (ma map[string]interface{})
```
GenerateGeneric form (map[string]interface) of the material filter




### <a name="MaterialAttributesTfs.GetFilter">func</a> (MaterialAttributesTfs) [GetFilter](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_tfs.go?s=933:994#L39)
``` go
func (mtfs MaterialAttributesTfs) GetFilter() *MaterialFilter
```
GetFilter from material attribute




### <a name="MaterialAttributesTfs.HasFilter">func</a> (MaterialAttributesTfs) [HasFilter](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material_tfs.go?s=827:877#L34)
``` go
func (mtfs MaterialAttributesTfs) HasFilter() bool
```
HasFilter in this material attribute




## <a name="MaterialFilter">type</a> [MaterialFilter](https://github.com/beamly/go-gocd/tree/master/gocd/pipeline.go?s=2165:2228#L57)
``` go
type MaterialFilter struct {
    Ignore []string `json:"ignore"`
}
```
MaterialFilter describes which globs to ignore










### <a name="MaterialFilter.GenerateGeneric">func</a> (\*MaterialFilter) [GenerateGeneric](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline_material.go?s=2884:2954#L115)
``` go
func (mf *MaterialFilter) GenerateGeneric() (g map[string]interface{})
```
GenerateGeneric form (map[string]interface) of the material filter




## <a name="MaterialRevision">type</a> [MaterialRevision](https://github.com/beamly/go-gocd/tree/master/gocd/pipeline.go?s=3189:3502#L85)
``` go
type MaterialRevision struct {
    Modifications []Modification `json:"modifications"`
    Material      struct {
        Description string `json:"description"`
        Fingerprint string `json:"fingerprint"`
        Type        string `json:"type"`
        ID          int    `json:"id"`
    } `json:"material"`
    Changed bool `json:"changed"`
}
```
MaterialRevision describes the uniquely identifiable version for the material which was pulled for this build










## <a name="Modification">type</a> [Modification](https://github.com/beamly/go-gocd/tree/master/gocd/pipeline.go?s=3595:3861#L97)
``` go
type Modification struct {
    EmailAddress string `json:"email_address"`
    ID           int    `json:"id"`
    ModifiedTime int    `json:"modified_time"`
    UserName     string `json:"user_name"`
    Comment      string `json:"comment"`
    Revision     string `json:"revision"`
}
```
Modification describes the commit/revision for the material which kicked off the build.










## <a name="PaginationResponse">type</a> [PaginationResponse](https://github.com/beamly/go-gocd/tree/master/gocd/gocd.go?s=2379:2505#L100)
``` go
type PaginationResponse struct {
    Offset   int `json:"offset"`
    Total    int `json:"total"`
    PageSize int `json:"page_size"`
}
```
PaginationResponse is a struct used to handle paging through resposnes.










## <a name="Parameter">type</a> [Parameter](https://github.com/beamly/go-gocd/tree/master/gocd/pipeline.go?s=1546:1628#L37)
``` go
type Parameter struct {
    Name  string `json:"name"`
    Value string `json:"value"`
}
```
Parameter represents a key/value










## <a name="PasswordFilePath">type</a> [PasswordFilePath](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go?s=6963:7026#L171)
``` go
type PasswordFilePath struct {
    Path string `xml:"path,attr"`
}
```
PasswordFilePath describes the location to set of user/passwords on disk










## <a name="PatchStringAction">type</a> [PatchStringAction](https://github.com/beamly/go-gocd/tree/master/gocd/environment.go?s=1679:1775#L45)
``` go
type PatchStringAction struct {
    Add    []string `json:"add"`
    Remove []string `json:"remove"`
}
```
PatchStringAction describes a collection of resources to add or remove.










## <a name="Pipeline">type</a> [Pipeline](https://github.com/beamly/go-gocd/tree/master/gocd/pipeline.go?s=378:1508#L18)
``` go
type Pipeline struct {
    Group                 string                 `json:"group,omitempty"`
    Links                 *HALLinks              `json:"_links,omitempty"`
    Name                  string                 `json:"name"`
    LabelTemplate         string                 `json:"label_template,omitempty"`
    EnablePipelineLocking bool                   `json:"enable_pipeline_locking,omitempty"`
    Template              string                 `json:"template,omitempty"`
    Origin                *PipelineConfigOrigin  `json:"origin,omitempty"`
    Parameters            []*Parameter           `json:"parameters,omitempty"`
    EnvironmentVariables  []*EnvironmentVariable `json:"environment_variables,omitempty"`
    Materials             []Material             `json:"materials,omitempty"`
    Label                 string                 `json:"label,omitempty"`
    Stages                []*Stage               `json:"stages,omitempty"`
    Version               string                 `json:"version,omitempty"`
}
```
Pipeline describes a pipeline object










### <a name="Pipeline.AddStage">func</a> (\*Pipeline) [AddStage](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline.go?s=750:791#L40)
``` go
func (p *Pipeline) AddStage(stage *Stage)
```
AddStage appends a stage to this pipeline




### <a name="Pipeline.GetLinks">func</a> (\*Pipeline) [GetLinks](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline.go?s=447:486#L25)
``` go
func (p *Pipeline) GetLinks() *HALLinks
```
GetLinks from pipeline




### <a name="Pipeline.GetName">func</a> (\*Pipeline) [GetName](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline.go?s=535:570#L30)
``` go
func (p *Pipeline) GetName() string
```
GetName of the pipeline




### <a name="Pipeline.GetStage">func</a> (\*Pipeline) [GetStage](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline.go?s=146:206#L9)
``` go
func (p *Pipeline) GetStage(stageName string) (stage *Stage)
```
GetStage from the pipeline template




### <a name="Pipeline.GetStages">func</a> (\*Pipeline) [GetStages](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline.go?s=45:84#L4)
``` go
func (p *Pipeline) GetStages() []*Stage
```
GetStages from the pipeline




### <a name="Pipeline.GetVersion">func</a> (\*Pipeline) [GetVersion](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline.go?s=1252:1300#L61)
``` go
func (p *Pipeline) GetVersion() (version string)
```
GetVersion retrieves a version string for this pipeline




### <a name="Pipeline.RemoveLinks">func</a> (\*Pipeline) [RemoveLinks](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline.go?s=368:400#L20)
``` go
func (p *Pipeline) RemoveLinks()
```
RemoveLinks from the pipeline object for json marshalling.




### <a name="Pipeline.SetStage">func</a> (\*Pipeline) [SetStage](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline.go?s=883:927#L45)
``` go
func (p *Pipeline) SetStage(newStage *Stage)
```
SetStage replaces a stage if it already exists




### <a name="Pipeline.SetStages">func</a> (\*Pipeline) [SetStages](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline.go?s=635:680#L35)
``` go
func (p *Pipeline) SetStages(stages []*Stage)
```
SetStages overwrites any existing stages




### <a name="Pipeline.SetVersion">func</a> (\*Pipeline) [SetVersion](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline.go?s=1121:1166#L56)
``` go
func (p *Pipeline) SetVersion(version string)
```
SetVersion sets a version string for this pipeline




## <a name="PipelineConfigOrigin">type</a> [PipelineConfigOrigin](https://github.com/beamly/go-gocd/tree/master/gocd/pipeline.go?s=1709:1799#L43)
``` go
type PipelineConfigOrigin struct {
    Type string `json:"type"`
    File string `json:"file"`
}
```
PipelineConfigOrigin describes where a pipeline config is being loaded from










## <a name="PipelineConfigRequest">type</a> [PipelineConfigRequest](https://github.com/beamly/go-gocd/tree/master/gocd/pipelineconfig.go?s=278:398#L12)
``` go
type PipelineConfigRequest struct {
    Group    string    `json:"group,omitempty"`
    Pipeline *Pipeline `json:"pipeline"`
}
```
PipelineConfigRequest describes a request object for creating or updating pipelines










### <a name="PipelineConfigRequest.GetVersion">func</a> (\*PipelineConfigRequest) [GetVersion](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline.go?s=1357:1419#L66)
``` go
func (pr *PipelineConfigRequest) GetVersion() (version string)
```
GetVersion of pipeline config




### <a name="PipelineConfigRequest.SetVersion">func</a> (\*PipelineConfigRequest) [SetVersion](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipeline.go?s=1491:1550#L71)
``` go
func (pr *PipelineConfigRequest) SetVersion(version string)
```
SetVersion of pipeline config




## <a name="PipelineConfigsService">type</a> [PipelineConfigsService](https://github.com/beamly/go-gocd/tree/master/gocd/pipelineconfig.go?s=154:189#L9)
``` go
type PipelineConfigsService service
```
PipelineConfigsService describes the HAL _link resource for the api response object for a pipelineconfig










### <a name="PipelineConfigsService.Create">func</a> (\*PipelineConfigsService) [Create](https://github.com/beamly/go-gocd/tree/master/gocd/pipelineconfig.go?s=1199:1333#L46)
``` go
func (pcs *PipelineConfigsService) Create(ctx context.Context, group string, p *Pipeline) (pr *Pipeline, resp *APIResponse, err error)
```
Create a pipeline configuration




### <a name="PipelineConfigsService.Delete">func</a> (\*PipelineConfigsService) [Delete](https://github.com/beamly/go-gocd/tree/master/gocd/pipelineconfig.go?s=1622:1727#L63)
``` go
func (pcs *PipelineConfigsService) Delete(ctx context.Context, name string) (string, *APIResponse, error)
```
Delete a pipeline configuration




### <a name="PipelineConfigsService.Get">func</a> (\*PipelineConfigsService) [Get](https://github.com/beamly/go-gocd/tree/master/gocd/pipelineconfig.go?s=457:573#L18)
``` go
func (pcs *PipelineConfigsService) Get(ctx context.Context, name string) (p *Pipeline, resp *APIResponse, err error)
```
Get a single PipelineTemplate object in the GoCD API.




### <a name="PipelineConfigsService.Update">func</a> (\*PipelineConfigsService) [Update](https://github.com/beamly/go-gocd/tree/master/gocd/pipelineconfig.go?s=790:923#L30)
``` go
func (pcs *PipelineConfigsService) Update(ctx context.Context, name string, p *Pipeline) (pr *Pipeline, resp *APIResponse, err error)
```
Update a pipeline configuration




## <a name="PipelineGroup">type</a> [PipelineGroup](https://github.com/beamly/go-gocd/tree/master/gocd/pipelinegroups.go?s=342:450#L12)
``` go
type PipelineGroup struct {
    Name      string      `json:"name"`
    Pipelines []*Pipeline `json:"pipelines"`
}
```
PipelineGroup describes a pipeline group API response.










## <a name="PipelineGroups">type</a> [PipelineGroups](https://github.com/beamly/go-gocd/tree/master/gocd/pipelinegroups.go?s=246:282#L9)
``` go
type PipelineGroups []*PipelineGroup
```
PipelineGroups represents a collection of pipeline groups










### <a name="PipelineGroups.GetGroupByPipeline">func</a> (\*PipelineGroups) [GetGroupByPipeline](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipelinegroups.go?s=443:522#L16)
``` go
func (pg *PipelineGroups) GetGroupByPipeline(pipeline *Pipeline) *PipelineGroup
```
GetGroupByPipeline finds the pipeline group for the pipeline supplied




### <a name="PipelineGroups.GetGroupByPipelineName">func</a> (\*PipelineGroups) [GetGroupByPipelineName](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipelinegroups.go?s=103:187#L4)
``` go
func (pg *PipelineGroups) GetGroupByPipelineName(pipelineName string) *PipelineGroup
```
GetGroupByPipelineName finds the pipeline group for the name of the pipeline supplied




## <a name="PipelineGroupsService">type</a> [PipelineGroupsService](https://github.com/beamly/go-gocd/tree/master/gocd/pipelinegroups.go?s=149:183#L6)
``` go
type PipelineGroupsService service
```
PipelineGroupsService describes the HAL _link resource for the api response object for a pipeline group response.










### <a name="PipelineGroupsService.List">func</a> (\*PipelineGroupsService) [List](https://github.com/beamly/go-gocd/tree/master/gocd/pipelinegroups.go?s=476:587#L18)
``` go
func (pgs *PipelineGroupsService) List(ctx context.Context, name string) (*PipelineGroups, *APIResponse, error)
```
List Pipeline groups




## <a name="PipelineHistory">type</a> [PipelineHistory](https://github.com/beamly/go-gocd/tree/master/gocd/pipeline.go?s=2294:2375#L62)
``` go
type PipelineHistory struct {
    Pipelines []*PipelineInstance `json:"pipelines"`
}
```
PipelineHistory describes the history of runs for a pipeline










## <a name="PipelineInstance">type</a> [PipelineInstance](https://github.com/beamly/go-gocd/tree/master/gocd/pipeline.go?s=2429:2719#L67)
``` go
type PipelineInstance struct {
    BuildCause   BuildCause `json:"build_cause"`
    CanRun       bool       `json:"can_run"`
    Name         string     `json:"name"`
    NaturalOrder int        `json:"natural_order"`
    Comment      string     `json:"comment"`
    Stages       []*Stage   `json:"stages"`
}
```
PipelineInstance describes a single pipeline run










## <a name="PipelineMaterial">type</a> [PipelineMaterial](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go?s=2926:3096#L71)
``` go
type PipelineMaterial struct {
    Name         string `xml:"pipelineName,attr"`
    StageName    string `xml:"stageName,attr"`
    MaterialName string `xml:"materialName,attr"`
}
```
PipelineMaterial part of cruise-control.xml. @TODO better documentation










## <a name="PipelineRequest">type</a> [PipelineRequest](https://github.com/beamly/go-gocd/tree/master/gocd/pipeline.go?s=232:336#L12)
``` go
type PipelineRequest struct {
    Group    string    `json:"group"`
    Pipeline *Pipeline `json:"pipeline"`
}
```
PipelineRequest describes a pipeline request object










## <a name="PipelineStatus">type</a> [PipelineStatus](https://github.com/beamly/go-gocd/tree/master/gocd/pipeline.go?s=3935:4072#L107)
``` go
type PipelineStatus struct {
    Locked      bool `json:"locked"`
    Paused      bool `json:"paused"`
    Schedulable bool `json:"schedulable"`
}
```
PipelineStatus describes whether a pipeline can be run or scheduled.










## <a name="PipelineTemplate">type</a> [PipelineTemplate](https://github.com/beamly/go-gocd/tree/master/gocd/pipelinetemplate.go?s=1135:1468#L41)
``` go
type PipelineTemplate struct {
    Links    *HALLinks                 `json:"_links,omitempty"`
    Name     string                    `json:"name"`
    Embedded *embeddedPipelineTemplate `json:"_embedded,omitempty"`
    Version  string                    `json:"template_version"`
    Stages   []*Stage                  `json:"stages,omitempty"`
}
```
PipelineTemplate describes a response from the API for a pipeline template object.










### <a name="PipelineTemplate.AddStage">func</a> (\*PipelineTemplate) [AddStage](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipelinetemplate.go?s=601:651#L29)
``` go
func (pt *PipelineTemplate) AddStage(stage *Stage)
```
AddStage appends a stage to this pipeline




### <a name="PipelineTemplate.GetName">func</a> (PipelineTemplate) [GetName](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipelinetemplate.go?s=367:410#L19)
``` go
func (pt PipelineTemplate) GetName() string
```
GetName of the pipeline template




### <a name="PipelineTemplate.GetStage">func</a> (PipelineTemplate) [GetStage](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipelinetemplate.go?s=164:224#L9)
``` go
func (pt PipelineTemplate) GetStage(stageName string) *Stage
```
GetStage from the pipeline template




### <a name="PipelineTemplate.GetStages">func</a> (PipelineTemplate) [GetStages](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipelinetemplate.go?s=54:101#L4)
``` go
func (pt PipelineTemplate) GetStages() []*Stage
```
GetStages from the pipeline template




### <a name="PipelineTemplate.GetVersion">func</a> (PipelineTemplate) [GetVersion](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipelinetemplate.go?s=1448:1504#L60)
``` go
func (pt PipelineTemplate) GetVersion() (version string)
```
GetVersion retrieves a version string for this pipeline




### <a name="PipelineTemplate.Pipelines">func</a> (PipelineTemplate) [Pipelines](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipelinetemplate.go?s=921:971#L39)
``` go
func (pt PipelineTemplate) Pipelines() []*Pipeline
```
Pipelines returns a list of Pipelines attached to this PipelineTemplate object.




### <a name="PipelineTemplate.RemoveLinks">func</a> (\*PipelineTemplate) [RemoveLinks](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipelinetemplate.go?s=775:816#L34)
``` go
func (pt *PipelineTemplate) RemoveLinks()
```
RemoveLinks gets the PipelineTemplate ready to be submitted to the GoCD API.




### <a name="PipelineTemplate.SetStage">func</a> (\*PipelineTemplate) [SetStage](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipelinetemplate.go?s=1057:1110#L44)
``` go
func (pt *PipelineTemplate) SetStage(newStage *Stage)
```
SetStage replaces a stage if it already exists




### <a name="PipelineTemplate.SetStages">func</a> (\*PipelineTemplate) [SetStages](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipelinetemplate.go?s=476:530#L24)
``` go
func (pt *PipelineTemplate) SetStages(stages []*Stage)
```
SetStages overwrites any existing stages




### <a name="PipelineTemplate.SetVersion">func</a> (\*PipelineTemplate) [SetVersion](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipelinetemplate.go?s=1307:1361#L55)
``` go
func (pt *PipelineTemplate) SetVersion(version string)
```
SetVersion sets a version string for this pipeline




## <a name="PipelineTemplateRequest">type</a> [PipelineTemplateRequest](https://github.com/beamly/go-gocd/tree/master/gocd/pipelinetemplate.go?s=266:406#L12)
``` go
type PipelineTemplateRequest struct {
    Name    string   `json:"name"`
    Stages  []*Stage `json:"stages"`
    Version string   `json:"version"`
}
```
PipelineTemplateRequest describes a PipelineTemplate










### <a name="PipelineTemplateRequest.GetVersion">func</a> (PipelineTemplateRequest) [GetVersion](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipelinetemplate.go?s=1731:1794#L70)
``` go
func (pt PipelineTemplateRequest) GetVersion() (version string)
```
GetVersion retrieves a version string for this pipeline




### <a name="PipelineTemplateRequest.SetVersion">func</a> (\*PipelineTemplateRequest) [SetVersion](https://github.com/beamly/go-gocd/tree/master/gocd/resource_pipelinetemplate.go?s=1583:1644#L65)
``` go
func (pt *PipelineTemplateRequest) SetVersion(version string)
```
SetVersion sets a version string for this pipeline




## <a name="PipelineTemplateResponse">type</a> [PipelineTemplateResponse](https://github.com/beamly/go-gocd/tree/master/gocd/pipelinetemplate.go?s=494:674#L19)
``` go
type PipelineTemplateResponse struct {
    Name     string `json:"name"`
    Embedded *struct {
        Pipelines []*struct {
            Name string `json:"name"`
        }
    } `json:"_embedded,omitempty"`
}
```
PipelineTemplateResponse describes an api response for a single pipeline templates










## <a name="PipelineTemplatesResponse">type</a> [PipelineTemplatesResponse](https://github.com/beamly/go-gocd/tree/master/gocd/pipelinetemplate.go?s=763:953#L29)
``` go
type PipelineTemplatesResponse struct {
    Links    *HALLinks `json:"_links,omitempty"`
    Embedded *struct {
        Templates []*PipelineTemplate `json:"templates"`
    } `json:"_embedded,omitempty"`
}
```
PipelineTemplatesResponse describes an api response for multiple pipeline templates










## <a name="PipelineTemplatesService">type</a> [PipelineTemplatesService](https://github.com/beamly/go-gocd/tree/master/gocd/pipelinetemplate.go?s=171:208#L9)
``` go
type PipelineTemplatesService service
```
PipelineTemplatesService describes the HAL _link resource for the api response object for a pipeline configuration objects.










### <a name="PipelineTemplatesService.Create">func</a> (\*PipelineTemplatesService) [Create](https://github.com/beamly/go-gocd/tree/master/gocd/pipelinetemplate.go?s=2299:2443#L75)
``` go
func (pts *PipelineTemplatesService) Create(ctx context.Context, name string, st []*Stage) (ptr *PipelineTemplate, resp *APIResponse, err error)
```
Create a new PipelineTemplate object in the GoCD API.




### <a name="PipelineTemplatesService.Delete">func</a> (\*PipelineTemplatesService) [Delete](https://github.com/beamly/go-gocd/tree/master/gocd/pipelinetemplate.go?s=3293:3400#L113)
``` go
func (pts *PipelineTemplatesService) Delete(ctx context.Context, name string) (string, *APIResponse, error)
```
Delete a PipelineTemplate from the GoCD API.




### <a name="PipelineTemplatesService.Get">func</a> (\*PipelineTemplatesService) [Get](https://github.com/beamly/go-gocd/tree/master/gocd/pipelinetemplate.go?s=1527:1654#L50)
``` go
func (pts *PipelineTemplatesService) Get(ctx context.Context, name string) (pt *PipelineTemplate, resp *APIResponse, err error)
```
Get a single PipelineTemplate object in the GoCD API.




### <a name="PipelineTemplatesService.List">func</a> (\*PipelineTemplatesService) [List](https://github.com/beamly/go-gocd/tree/master/gocd/pipelinetemplate.go?s=1900:2017#L62)
``` go
func (pts *PipelineTemplatesService) List(ctx context.Context) (pt []*PipelineTemplate, resp *APIResponse, err error)
```
List all PipelineTemplate objects in the GoCD API.




### <a name="PipelineTemplatesService.Update">func</a> (\*PipelineTemplatesService) [Update](https://github.com/beamly/go-gocd/tree/master/gocd/pipelinetemplate.go?s=2772:2931#L95)
``` go
func (pts *PipelineTemplatesService) Update(ctx context.Context, name string, template *PipelineTemplate) (ptr *PipelineTemplate, resp *APIResponse, err error)
```
Update an PipelineTemplate object in the GoCD API.




## <a name="PipelinesService">type</a> [PipelinesService](https://github.com/beamly/go-gocd/tree/master/gocd/pipeline.go?s=146:175#L9)
``` go
type PipelinesService service
```
PipelinesService describes the HAL _link resource for the api response object for a pipelineconfig










### <a name="PipelinesService.GetHistory">func</a> (\*PipelinesService) [GetHistory](https://github.com/beamly/go-gocd/tree/master/gocd/pipeline.go?s=5576:5713#L152)
``` go
func (pgs *PipelinesService) GetHistory(ctx context.Context, name string, offset int) (pt *PipelineHistory, resp *APIResponse, err error)
```
GetHistory returns a list of pipeline instances describing the pipeline history.




### <a name="PipelinesService.GetInstance">func</a> (\*PipelinesService) [GetInstance](https://github.com/beamly/go-gocd/tree/master/gocd/pipeline.go?s=5140:5279#L140)
``` go
func (pgs *PipelinesService) GetInstance(ctx context.Context, name string, offset int) (pt *PipelineInstance, resp *APIResponse, err error)
```
GetInstance of a pipeline run.




### <a name="PipelinesService.GetStatus">func</a> (\*PipelinesService) [GetStatus](https://github.com/beamly/go-gocd/tree/master/gocd/pipeline.go?s=4157:4292#L114)
``` go
func (pgs *PipelinesService) GetStatus(ctx context.Context, name string, offset int) (ps *PipelineStatus, resp *APIResponse, err error)
```
GetStatus returns a list of pipeline instanves describing the pipeline history.




### <a name="PipelinesService.Pause">func</a> (\*PipelinesService) [Pause](https://github.com/beamly/go-gocd/tree/master/gocd/pipeline.go?s=4528:4624#L125)
``` go
func (pgs *PipelinesService) Pause(ctx context.Context, name string) (bool, *APIResponse, error)
```
Pause allows a pipeline to handle new build events




### <a name="PipelinesService.ReleaseLock">func</a> (\*PipelinesService) [ReleaseLock](https://github.com/beamly/go-gocd/tree/master/gocd/pipeline.go?s=4945:5047#L135)
``` go
func (pgs *PipelinesService) ReleaseLock(ctx context.Context, name string) (bool, *APIResponse, error)
```
ReleaseLock frees a pipeline to handle new build events




### <a name="PipelinesService.Unpause">func</a> (\*PipelinesService) [Unpause](https://github.com/beamly/go-gocd/tree/master/gocd/pipeline.go?s=4733:4831#L130)
``` go
func (pgs *PipelinesService) Unpause(ctx context.Context, name string) (bool, *APIResponse, error)
```
Unpause allows a pipeline to handle new build events




## <a name="PluggableInstanceSettings">type</a> [PluggableInstanceSettings](https://github.com/beamly/go-gocd/tree/master/gocd/plugin.go?s=1017:1172#L31)
``` go
type PluggableInstanceSettings struct {
    Configurations []PluginConfiguration `json:"configurations"`
    View           PluginView            `json:"view"`
}
```
PluggableInstanceSettings describes plugin configuration










## <a name="Plugin">type</a> [Plugin](https://github.com/beamly/go-gocd/tree/master/gocd/plugin.go?s=430:955#L20)
``` go
type Plugin struct {
    Links                     *HALLinks                 `json:"_links"`
    ID                        string                    `json:"id"`
    Name                      string                    `json:"name"`
    DisplayName               string                    `json:"display_name"`
    Version                   string                    `json:"version"`
    Type                      string                    `json:"type"`
    PluggableInstanceSettings PluggableInstanceSettings `json:"pluggable_instance_settings"`
}
```
Plugin describes a single plugin resource.










## <a name="PluginConfiguration">type</a> [PluginConfiguration](https://github.com/beamly/go-gocd/tree/master/gocd/jobs.go?s=2830:2972#L72)
``` go
type PluginConfiguration struct {
    Key      string                      `json:"key"`
    Metadata PluginConfigurationMetadata `json:"metadata"`
}
```
PluginConfiguration describes how to reference a plugin.










## <a name="PluginConfigurationKVPair">type</a> [PluginConfigurationKVPair](https://github.com/beamly/go-gocd/tree/master/gocd/jobs.go?s=3323:3420#L85)
``` go
type PluginConfigurationKVPair struct {
    Key   string `json:"key"`
    Value string `json:"value"`
}
```
PluginConfigurationKVPair describes a key/value pair of plugin configurations.










## <a name="PluginConfigurationMetadata">type</a> [PluginConfigurationMetadata](https://github.com/beamly/go-gocd/tree/master/gocd/jobs.go?s=3073:3239#L78)
``` go
type PluginConfigurationMetadata struct {
    Secure         bool `json:"secure"`
    Required       bool `json:"required"`
    PartOfIdentity bool `json:"part_of_identity"`
}
```
PluginConfigurationMetadata describes the schema for a single configuration option for a plugin










## <a name="PluginView">type</a> [PluginView](https://github.com/beamly/go-gocd/tree/master/gocd/plugin.go?s=1229:1290#L37)
``` go
type PluginView struct {
    Template string `json:"template"`
}
```
PluginView describes any view attached to a plugin.










## <a name="PluginsResponse">type</a> [PluginsResponse](https://github.com/beamly/go-gocd/tree/master/gocd/plugin.go?s=230:382#L12)
``` go
type PluginsResponse struct {
    Links    *HALLinks `json:"_links"`
    Embedded struct {
        PluginInfo []*Plugin `json:"plugin_info"`
    } `json:"_embedded"`
}
```
PluginsResponse describes the response obejct for a plugin API call.










## <a name="PluginsService">type</a> [PluginsService](https://github.com/beamly/go-gocd/tree/master/gocd/plugin.go?s=129:156#L9)
``` go
type PluginsService service
```
PluginsService exposes calls for interacting with Plugin objects in the GoCD API.










### <a name="PluginsService.Get">func</a> (\*PluginsService) [Get](https://github.com/beamly/go-gocd/tree/master/gocd/plugin.go?s=1668:1773#L54)
``` go
func (ps *PluginsService) Get(ctx context.Context, name string) (p *Plugin, resp *APIResponse, err error)
```
Get retrieves information about a specific plugin.




### <a name="PluginsService.List">func</a> (\*PluginsService) [List](https://github.com/beamly/go-gocd/tree/master/gocd/plugin.go?s=1322:1413#L42)
``` go
func (ps *PluginsService) List(ctx context.Context) (*PluginsResponse, *APIResponse, error)
```
List retrieves all plugins




## <a name="Properties">type</a> [Properties](https://github.com/beamly/go-gocd/tree/master/gocd/resource_properties.go?s=289:433#L21)
``` go
type Properties struct {
    UnmarshallWithHeader bool
    IsDatum              bool
    Header               []string
    DataFrame            [][]string
}
```
Properties describes a properties resource in the GoCD API.







### <a name="NewPropertiesFrame">func</a> [NewPropertiesFrame](https://github.com/beamly/go-gocd/tree/master/gocd/resource_properties.go?s=513:566#L29)
``` go
func NewPropertiesFrame(frame [][]string) *Properties
```
NewPropertiesFrame generate a new data frame for properties on a gocd job.





### <a name="Properties.AddRow">func</a> (\*Properties) [AddRow](https://github.com/beamly/go-gocd/tree/master/gocd/resource_properties.go?s=1003:1043#L53)
``` go
func (pr *Properties) AddRow(r []string)
```
AddRow to an existing properties data frame




### <a name="Properties.Get">func</a> (Properties) [Get](https://github.com/beamly/go-gocd/tree/master/gocd/resource_properties.go?s=761:816#L42)
``` go
func (pr Properties) Get(row int, column string) string
```
Get a single parameter value for a given run of the job.




### <a name="Properties.MarshalJSON">func</a> (\*Properties) [MarshalJSON](https://github.com/beamly/go-gocd/tree/master/gocd/resource_properties.go?s=2440:2491#L116)
``` go
func (pr *Properties) MarshalJSON() ([]byte, error)
```
MarshalJSON converts the properties structure to a list of maps




### <a name="Properties.MarshallCSV">func</a> (Properties) [MarshallCSV](https://github.com/beamly/go-gocd/tree/master/gocd/resource_properties.go?s=1331:1381#L66)
``` go
func (pr Properties) MarshallCSV() (string, error)
```
MarshallCSV returns the data frame as a string




### <a name="Properties.SetRow">func</a> (\*Properties) [SetRow](https://github.com/beamly/go-gocd/tree/master/gocd/resource_properties.go?s=1118:1167#L58)
``` go
func (pr *Properties) SetRow(row int, r []string)
```
SetRow in an existing data frame




### <a name="Properties.UnmarshallCSV">func</a> (\*Properties) [UnmarshallCSV](https://github.com/beamly/go-gocd/tree/master/gocd/resource_properties.go?s=1716:1769#L83)
``` go
func (pr *Properties) UnmarshallCSV(raw string) error
```
UnmarshallCSV returns the data frame from a string




### <a name="Properties.Write">func</a> (\*Properties) [Write](https://github.com/beamly/go-gocd/tree/master/gocd/resource_properties.go?s=2153:2209#L104)
``` go
func (pr *Properties) Write(p []byte) (n int, err error)
```
Write the data frame to a byte stream as a csv.




## <a name="PropertiesService">type</a> [PropertiesService](https://github.com/beamly/go-gocd/tree/master/gocd/properties.go?s=125:155#L10)
``` go
type PropertiesService service
```
PropertiesService describes Actions which can be performed on agents










### <a name="PropertiesService.Create">func</a> (\*PropertiesService) [Create](https://github.com/beamly/go-gocd/tree/master/gocd/properties.go?s=1566:1723#L52)
``` go
func (ps *PropertiesService) Create(ctx context.Context, name string, value string, pr *PropertyRequest) (responseIsValid bool, resp *APIResponse, err error)
```
Create a specific property for the given job/pipeline/stage run.




### <a name="PropertiesService.Get">func</a> (\*PropertiesService) [Get](https://github.com/beamly/go-gocd/tree/master/gocd/properties.go?s=1115:1237#L42)
``` go
func (ps *PropertiesService) Get(ctx context.Context, name string, pr *PropertyRequest) (*Properties, *APIResponse, error)
```
Get a specific property for the given job/pipeline/stage run.




### <a name="PropertiesService.List">func</a> (\*PropertiesService) [List](https://github.com/beamly/go-gocd/tree/master/gocd/properties.go?s=681:791#L31)
``` go
func (ps *PropertiesService) List(ctx context.Context, pr *PropertyRequest) (*Properties, *APIResponse, error)
```
List the properties for the given job/pipeline/stage run.




### <a name="PropertiesService.ListHistorical">func</a> (\*PropertiesService) [ListHistorical](https://github.com/beamly/go-gocd/tree/master/gocd/properties.go?s=2448:2568#L76)
``` go
func (ps *PropertiesService) ListHistorical(ctx context.Context, pr *PropertyRequest) (*Properties, *APIResponse, error)
```
ListHistorical properties for a given pipeline, stage, job.




## <a name="PropertyCreateResponse">type</a> [PropertyCreateResponse](https://github.com/beamly/go-gocd/tree/master/gocd/properties.go?s=552:618#L25)
``` go
type PropertyCreateResponse struct {
    Name  string
    Value string
}
```
PropertyCreateResponse handles the parsing of the response when creating a property










## <a name="PropertyRequest">type</a> [PropertyRequest](https://github.com/beamly/go-gocd/tree/master/gocd/properties.go?s=251:463#L13)
``` go
type PropertyRequest struct {
    Pipeline        string
    PipelineCounter int
    Stage           string
    StageCounter    int
    Job             string
    LimitPipeline   string
    Limit           int
    Single          bool
}
```
PropertyRequest describes the parameters to be submitted when calling/creating properties.










## <a name="Stage">type</a> [Stage](https://github.com/beamly/go-gocd/tree/master/gocd/stages.go?s=166:781#L7)
``` go
type Stage struct {
    Name                  string                 `json:"name"`
    FetchMaterials        bool                   `json:"fetch_materials"`
    CleanWorkingDirectory bool                   `json:"clean_working_directory"`
    NeverCleanupArtifacts bool                   `json:"never_cleanup_artifacts"`
    Approval              *Approval              `json:"approval,omitempty"`
    EnvironmentVariables  []*EnvironmentVariable `json:"environment_variables,omitempty"`
    Resources             []string               `json:"resource,omitempty"`
    Jobs                  []*Job                 `json:"jobs,omitempty"`
}
```
Stage represents a GoCD Stage object.










### <a name="Stage.Clean">func</a> (\*Stage) [Clean](https://github.com/beamly/go-gocd/tree/master/gocd/resource_stages.go?s=740:763#L39)
``` go
func (s *Stage) Clean()
```
Clean the approvel step.




### <a name="Stage.JSONString">func</a> (\*Stage) [JSONString](https://github.com/beamly/go-gocd/tree/master/gocd/resource_stages.go?s=116:160#L9)
``` go
func (s *Stage) JSONString() (string, error)
```
JSONString returns a string of this stage as a JSON object.




### <a name="Stage.Validate">func</a> (\*Stage) [Validate](https://github.com/beamly/go-gocd/tree/master/gocd/resource_stages.go?s=409:441#L20)
``` go
func (s *Stage) Validate() error
```
Validate ensures the attributes attached to this structure are ready for submission to the GoCD API.




## <a name="StageContainer">type</a> [StageContainer](https://github.com/beamly/go-gocd/tree/master/gocd/resource.go?s=125:310#L6)
``` go
type StageContainer interface {
    GetName() string
    SetStage(stage *Stage)
    GetStage(string) *Stage
    SetStages(stages []*Stage)
    GetStages() []*Stage
    AddStage(stage *Stage)
    Versioned
}
```
StageContainer describes structs which contain stages, eg Pipelines and PipelineTemplates










## <a name="StagesService">type</a> [StagesService](https://github.com/beamly/go-gocd/tree/master/gocd/stages.go?s=97:123#L4)
``` go
type StagesService service
```
StagesService exposes calls for interacting with Stage objects in the GoCD API.










## <a name="StringResponse">type</a> [StringResponse](https://github.com/beamly/go-gocd/tree/master/gocd/gocd.go?s=985:1048#L50)
``` go
type StringResponse struct {
    Message string `json:"message"`
}
```
StringResponse handles the unmarshaling of the single string response from DELETE requests.










## <a name="Tab">type</a> [Tab](https://github.com/beamly/go-gocd/tree/master/gocd/jobs.go?s=2242:2315#L51)
``` go
type Tab struct {
    Name string `json:"name"`
    Path string `json:"path"`
}
```
Tab description in a gocd job










## <a name="Task">type</a> [Task](https://github.com/beamly/go-gocd/tree/master/gocd/jobs.go?s=3471:3579#L91)
``` go
type Task struct {
    Type       string         `json:"type"`
    Attributes TaskAttributes `json:"attributes"`
}
```
Task Describes a Task object in the GoCD api.










### <a name="Task.Validate">func</a> (\*Task) [Validate](https://github.com/beamly/go-gocd/tree/master/gocd/resource_task.go?s=76:107#L6)
``` go
func (t *Task) Validate() error
```
Validate each of the possible task types.




## <a name="TaskAttributes">type</a> [TaskAttributes](https://github.com/beamly/go-gocd/tree/master/gocd/jobs.go?s=3640:4851#L97)
``` go
type TaskAttributes struct {
    RunIf               []string                    `json:"run_if,omitempty"`
    Command             string                      `json:"command,omitempty"`
    WorkingDirectory    string                      `json:"working_directory,omitempty"`
    Arguments           []string                    `json:"arguments,omitempty"`
    BuildFile           string                      `json:"build_file,omitempty"`
    Target              string                      `json:"target,omitempty"`
    NantPath            string                      `json:"nant_path,omitempty"`
    Pipeline            string                      `json:"pipeline,omitempty"`
    Stage               string                      `json:"stage,omitempty"`
    Job                 string                      `json:"job,omitempty"`
    Source              string                      `json:"source,omitempty"`
    IsSourceAFile       bool                        `json:"is_source_a_file,omitempty"`
    Destination         string                      `json:"destination,omitempty"`
    PluginConfiguration *TaskPluginConfiguration    `json:"plugin_configuration,omitempty"`
    Configuration       []PluginConfigurationKVPair `json:"configuration,omitempty"`
}
```
TaskAttributes describes all the properties for a Task.










### <a name="TaskAttributes.ValidateAnt">func</a> (\*TaskAttributes) [ValidateAnt](https://github.com/beamly/go-gocd/tree/master/gocd/jobs_validation.go?s=623:667#L24)
``` go
func (t *TaskAttributes) ValidateAnt() error
```
ValidateAnt checks that the specified values for the Task struct are correct for a an Ant task




### <a name="TaskAttributes.ValidateExec">func</a> (\*TaskAttributes) [ValidateExec](https://github.com/beamly/go-gocd/tree/master/gocd/jobs_validation.go?s=132:177#L6)
``` go
func (t *TaskAttributes) ValidateExec() error
```
ValidateExec checks that the specified values for the Task struct are correct for a cli exec task




## <a name="TaskPluginConfiguration">type</a> [TaskPluginConfiguration](https://github.com/beamly/go-gocd/tree/master/gocd/jobs.go?s=4925:5025#L116)
``` go
type TaskPluginConfiguration struct {
    ID      string `json:"id"`
    Version string `json:"version"`
}
```
TaskPluginConfiguration is for specifying options for pluggable task










## <a name="TimeoutField">type</a> [TimeoutField](https://github.com/beamly/go-gocd/tree/master/gocd/jobs.go?s=6557:6578#L162)
``` go
type TimeoutField int
```
TimeoutField helps manage the marshalling of the timoeut field which can be both "never" and an integer










### <a name="TimeoutField.MarshalJSON">func</a> (TimeoutField) [MarshalJSON](https://github.com/beamly/go-gocd/tree/master/gocd/resource_jobs.go?s=885:943#L48)
``` go
func (tf TimeoutField) MarshalJSON() (b []byte, err error)
```
MarshallJSON of TimeoutField into a string




### <a name="TimeoutField.UnmarshalJSON">func</a> (\*TimeoutField) [UnmarshalJSON](https://github.com/beamly/go-gocd/tree/master/gocd/resource_jobs.go?s=556:615#L31)
``` go
func (tf *TimeoutField) UnmarshalJSON(b []byte) (err error)
```
UnmarshalJSON and handle "never", "null", and integers.




## <a name="Version">type</a> [Version](https://github.com/beamly/go-gocd/tree/master/gocd/configuration.go?s=8090:8365#L207)
``` go
type Version struct {
    Links       *HALLinks `json:"_links"`
    Version     string    `json:"version"`
    BuildNumber string    `json:"build_number"`
    GitSHA      string    `json:"git_sha"`
    FullVersion string    `json:"full_version"`
    CommitURL   string    `json:"commit_url"`
}
```
Version part of cruise-control.xml. @TODO better documentation










## <a name="Versioned">type</a> [Versioned](https://github.com/beamly/go-gocd/tree/master/gocd/resource.go?s=521:598#L23)
``` go
type Versioned interface {
    GetVersion() string
    SetVersion(version string)
}
```
Versioned describes resources which can get and set versions














- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
