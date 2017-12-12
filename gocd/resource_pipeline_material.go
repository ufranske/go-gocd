package gocd

import (
	"encoding/json"
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

// UnmarshalJSON string into a Material struct
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
	case "plugin":
		m.Attributes = &MaterialAttributesPlugin{}
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

func unmarshallMaterialFilter(i interface{}) *MaterialFilter {
	return nil
}
