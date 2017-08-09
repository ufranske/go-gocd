package gocd

import (
	"encoding/xml"
)

type ConfigTasks struct {
	Tasks []ConfigTask `xml:",any"`
}

type ConfigTask struct {
	// Because we need to preserve the order of tasks, and we have an array of elements with mixed types,
	// we need to use this generic xml type for tasks.
	XMLName  TaskTypeName `json:",omitempty"`
	Type     string `xml:"type,omitempty"`
	RunIf    ConfigTaskRunIf `xml:"runif"`
	Command  string          `xml:"command,attr,omitempty"  json:",omitempty"`
	Args     []string        `xml:"arg,omitempty"  json:",omitempty"`
	Pipeline string          `xml:"pipeline,attr,omitempty"  json:",omitempty"`
	Stage    string          `xml:"stage,attr,omitempty"  json:",omitempty"`
	Job      string          `xml:"job,attr,omitempty"  json:",omitempty"`
	SrcFile  string          `xml:"srcfile,attr,omitempty"  json:",omitempty"`
	SrcDir   string          `xml:"srcdir,attr,omitempty"  json:",omitempty"`
}

type ConfigTaskRunIf struct {
	Status string `xml:"status,attr"`
}

type TaskTypeName xml.Name

//func (ttn *ConfigTask) MarshalJSON() ([]byte, error) {
//	ttn.Type = ttn.XMLName.Local
//	j, e := json.Marshal(ttn)
//	if e != nil {
//		return nil, e
//	}
//	return j, nil
//}
//
//func (ttn *TaskTypeName) UnmarshalJSON(j []byte) error {
//	return nil
//}
