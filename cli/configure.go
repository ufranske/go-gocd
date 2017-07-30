package cli

import (
	"github.com/segmentio/go-prompt"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/user"
	"strings"
)

const (
	ConfigureCommandName  = "configure"
	ConfigureCommandUsage = "Generate configuration file ~/.gocd.conf"
	ConfigDirectoryPath   = "~/.gocd.conf"
	EnvVarServer          = "GOCD_SERVER"
	EnvVarUsername        = "GOCD_USERNAME"
	EnvVarPassword        = "GOCD_PASSWORD"
)

func ConfigureAction(c *cli.Context) error {

	s, err := generateConfigFile()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(configFilePath(), []byte(s), 0644)
	if err != nil {
		return err
	}

	return nil
}

type Configuration struct {
	Server   string `yaml:"server"`
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
	SslCheck bool   `yaml:"ssl_check,omitempty"`
}

func (c *Configuration) HasAuth() bool {
	return (c.Username != "") && (c.Password != "")
}

// Build a default template
func generateConfigFile() (string, error) {
	cfg := Configuration{}
	cfg.Server = prompt.StringRequired("GoCD Server (should contain '/go/' suffix)")
	if u := prompt.String("Client Username"); u != "" {
		cfg.Username = u
	}
	if p := prompt.PasswordMasked("Client Password"); p != "" {
		cfg.Password = p
	}

	s, err := yaml.Marshal(cfg)
	if err != nil {
		return "", err
	}

	return string(s), nil
}
func configFilePath() string {
	// @TODO Make it work for windows. Maybe...
	usr, _ := user.Current()
	return strings.Replace(ConfigDirectoryPath, "~", usr.HomeDir, 1)
}

func loadConfig() (*Configuration, error) {
	s, err := ioutil.ReadFile(configFilePath())
	if err != nil {
		return nil, err
	}

	cfg := Configuration{}
	err = yaml.Unmarshal(s, &cfg)
	if err != nil {
		return nil, err
	}

	if server := os.Getenv(EnvVarServer); server != "" {
		cfg.Server = server
	}

	if username := os.Getenv(EnvVarUsername); username != "" {
		cfg.Username = username
	}

	if password := os.Getenv(EnvVarPassword); password != "" {
		cfg.Password = password
	}

	return &cfg, nil
}

func ConfigureCommand() *cli.Command {
	return &cli.Command{
		Name:   ConfigureCommandName,
		Usage:  ConfigureCommandUsage,
		Action: ConfigureAction,
	}
}
