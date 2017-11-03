package gocd

type MaterialAttributes interface {
	equal(attributes interface{}) (bool, error)
	UnmarshallInterface(map[string]interface{})
}

// MaterialAttributesGit describes a git material
type MaterialAttributesGit struct {
	Name        string `json:"name,omitempty"`
	URL         string `json:"url,omitempty"`
	Branch      string `json:"branch,omitempty"`
	Destination string `json:"destination,omitempty"`
	AutoUpdate  bool   `json:"auto_update,omitempty"`

	Filter       *MaterialFilter `json:"filter,omitempty"`
	InvertFilter bool            `json:"invert_filter"`

	SubmoduleFolder string `json:"submodule_folder,omitempty"`
	ShallowClone    bool   `json:"shallow_clone,omitempty"`
}

// MaterialAttributes describes a material type
type MaterialAttributesSvn struct {
	Name              string `json:"name,omitempty"`
	URL               string `json:"url,omitempty"`
	Username          string `json:"username"`
	Password          string `json:"password"`
	EncryptedPassword string `json:"encrypted_password"`
	Destination       string `json:"destination,omitempty"`

	Filter         *MaterialFilter `json:"filter,omitempty"`
	InvertFilter   bool            `json:"invert_filter"`
	AutoUpdate     bool            `json:"auto_update,omitempty"`
	CheckExternals bool            `json:"check_externals"`
}
