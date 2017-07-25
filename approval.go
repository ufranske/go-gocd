package gocd

type Approval struct {
	Type          string `json:"type"`
	Authorization Authorization `json:"authorization"`
}

type Authorization struct {
	Users []string `json:"users"`
	Roles []string `json:"roles"`
}
