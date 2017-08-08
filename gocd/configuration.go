package gocd

import (
	"context"
)

// ConfigurationService describes the HAL _link resource for the api response object for a pipelineconfig
type ConfigurationService service

type ConfigXML struct {
	Repositories []ConfigRepository `xml:"repositories>repository"`
	Server       ConfigServer `xml:"server"`
}

type ConfigRepository struct {
	ID   string `xml:"id"`
	Name string `xml:"name"`
	PluginConfiguration
}

type ConfigServer struct {
	MailHost                  MailHost `xml:"mailhost"`
	Security                  ConfigSecurity `xml:"security"`
	Elastic                   ConfigElastic  `xml:"elastic"`
	ArtifactsDir              string         `xml:"artifactsdir,attr"`
	SiteUrl                   string         `xml:"siteUrl,attr"`
	SecureSiteUrl             string         `xml:"secureSiteUrl,attr"`
	PurgeStart                string         `xml:"purgeStart,attr"`
	PurgeUpTo                 string         `xml:"purgeUpto,attr"`
	JobTimeout                int            `xml:"jobTimeout,attr"`
	AgentAutoRegisterKey      string         `xml:"agentAutoRegisterKey,attr"`
	WebhookSecret             string         `xml:"webhookSecret,attr"`
	CommandRepositoryLocation string         `xml:"commandRepositoryLocation,attr"`
	ServerId                  string         `xml:"serverId,attr"`
}

type MailHost struct {
	Hostname string `xml:"hostname,attr"`
	Port     int `xml:"port,attr"`
	TLS      bool `xml:"tls,attr"`
	From     string `xml:"from,attr"`
	Admin    string `xml:"admin,attr"`
}

type ConfigSecurity struct {
	AuthConfigs []ConfigAuthConfig `xml:"authConfigs>authConfig"`
	Roles       []ConfigRole `xml:"roles>role"`
	Admins      []string `xml:"admins>user"`
}

type ConfigRole struct {
	Name  string `xml:"name,attr"`
	Users []string `xml:"users>user"`
}

type ConfigAuthConfig struct {
	ID         string `xml:"id,attr"`
	PluginId   string `xml:"pluginId,attr"`
	Properties []ConfigProperty `xml:"property"`
}

type ConfigElastic struct {
	Profiles []ConfigElasticProfile `xml:"profiles>profile"`
}

type ConfigElasticProfile struct {
	ID         string           `xml:"id,attr"`
	PluginID   string           `xml:"pluginId,attr"`
	Properties []ConfigProperty `xml:"property"`
}

type ConfigProperty struct {
	Key   string `xml:"key"`
	Value string `xml:"value"`
}

// Get will retrieve all agents, their status, and metadata from the GoCD Server.
// Get returns a list of pipeline instanves describing the pipeline history.
func (cd *ConfigurationService) Get(ctx context.Context) (*ConfigXML, *APIResponse, error) {
	req, err := cd.client.NewRequest("GET", "admin/config.xml", nil, "")
	if err != nil {
		return nil, nil, err
	}

	cx := ConfigXML{}
	resp, err := cd.client.Do(ctx, req, &cx, responseTypeXML)
	if err != nil {
		return nil, resp, err
	}

	return &cx, resp, nil
}
