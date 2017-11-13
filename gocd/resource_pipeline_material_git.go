package gocd

import (
	"fmt"
	"errors"
)

func (mag *MaterialAttributesGit) equal(a2i interface{}) (bool, error) {
	var ok bool
	a2, ok := a2i.(*MaterialAttributesGit)
	if !ok {
		return false, errors.New("Can only compare with same material type.")
	}
	urlsEqual := mag.URL == a2.URL
	branchesEqual := mag.Branch == a2.Branch ||
		mag.Branch == "" && a2.Branch == "master" ||
		mag.Branch == "master" && a2.Branch == ""

	if !urlsEqual {
		return false, nil
	}
	return branchesEqual, nil
}

func (mag *MaterialAttributesGit) UnmarshallInterface(i map[string]interface{}) {
	for key, value := range i {
		if value == nil {
			continue
		}
		switch key {
		case "name":
			mag.Name = value.(string)
		case "url":
			mag.URL = value.(string)
		case "auto_update":
			mag.AutoUpdate = value.(bool)
		case "branch":
			mag.Branch = value.(string)
		case "submodule_folder":
			mag.SubmoduleFolder = value.(string)
		case "destination":
			mag.Destination = value.(string)
		case "shallow_clone":
			mag.ShallowClone = value.(bool)
		case "invert_filter":
			mag.InvertFilter = value.(bool)
			//case "filter":
			//	mag.Filter = value.(string)
		default:
			fmt.Println(value)
			fmt.Println(key)
		}
	}
}
