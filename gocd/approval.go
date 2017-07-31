package gocd

type Approval struct {
	Type          string        `json:"type"`
	Authorization *Authorization `json:"authorization,omitempty"`
}

type Authorization struct {
	Users []string `json:"users,omitempty"`
	Roles []string `json:"roles,omitempty"`
}
