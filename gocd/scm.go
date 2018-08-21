package gocd

// SCMsService exposes calls for interacting with SCM objects in the GoCD API.
type SCMsService service

type SCM struct {
	Links          *HALLinks           `json:"links"`
	ID             string              `json:"id"`
	Name           string              `json:"name"`
	AutoUpdate     bool                `json:"auto_update"`
	PluginMetadata *SCMMetadata        `json:"plugin_metadata"`
	Configuration  []*SCMConfiguration `json:"configuration"`
}

type SCMMetadata struct {
	ID      string `json:"id"`
	Version string `json:"version"`
}

type SCMConfiguration struct {
	Key            string `json:"key"`
	Value          string `json:"value"`
	EncryptedValue string `json:"encrypted_value"`
}
