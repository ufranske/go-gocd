package gocd

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// Equal is true if the two materials are logically equivalent. Not neccesarily literally equal.
func (m Material) Equal(a *Material) (isEqual bool, err error) {
	if m.Type != a.Type {
		return false, nil
	}
	switch m.Type {
	case "git":
		if isEqual, err = m.Attributes.equal(a.Attributes); err != nil {
			return false, err
		}
	default:
		panic(fmt.Errorf("Material comparison not implemented for '%s'", m.Type))
	}
	return isEqual, nil
}

func (m *Material) UnmarshalJSON(b []byte) error {
	temp := map[string]interface{}{}
	json.Unmarshal(b, &temp)

	switch m.Type = temp["type"].(string); strings.ToLower(m.Type) {
	case "git":
		m.Attributes = &MaterialAttributesGit{}
	case "svn":
		m.Attributes = &MaterialAttributesSvn{}
	case "hg":
		m.Attributes = &MaterialAttributesHg{}
	case "p4":
		m.Attributes = &MaterialAttributesP4{}
	case "tfs":
		m.Attributes = &MaterialAttributesTfs{}
	case "dependency":
		m.Attributes = &MaterialAttributesDependency{}
	case "package":
		m.Attributes = &MaterialAttributesPackage{}
	default:
		return fmt.Errorf("Unexpected Material type: '%s'", m.Type)
	}

	for key, value := range temp {
		if value == nil {
			continue
		}
		switch key {
		case "attributes":
			m.Attributes.UnmarshallInterface(temp["attributes"].(map[string]interface{}))
		case "fingerprint":
			m.Fingerprint = value.(string)
		case "description":
			m.Description = value.(string)
		case "type":
			continue
		default:
			return fmt.Errorf("Unexpected key: '%s'", key)
		}
	}

	return nil
}

func (mag *MaterialAttributesGit) equal(a2i interface{}) (bool, error) {
	//var a2 MaterialAttributesGit
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
