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

	var rawAttributes map[string]interface{}
	for key, value := range temp {
		if value == nil {
			continue
		}
		switch key {
		case "attributes":
			rawAttributes = value.(map[string]interface{})
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

	switch m.Type = temp["type"].(string); strings.ToLower(m.Type) {
	case "git":
		mag := MaterialAttributesGit{}
		unmarshallMaterialAttributesGit(&mag, rawAttributes)
		m.Attributes = mag
	case "svn":
		mas := &MaterialAttributesSvn{}
		unmarshallMaterialAttributesSvn(mas, rawAttributes)
		m.Attributes = mas
	case "hg":
		mah := &MaterialAttributesHg{}
		unmarshallMaterialAttributesHg(mah, rawAttributes)
		m.Attributes = mah
	case "p4":
		map4 := &MaterialAttributesP4{}
		unmarshallMaterialAttributesP4(map4, rawAttributes)
		m.Attributes = map4
	case "tfs":
		mat := &MaterialAttributesTfs{}
		unmarshallMaterialAttributesTfs(mat, rawAttributes)
		m.Attributes = mat
	case "dependency":
		mad := &MaterialAttributesDependency{}
		unmarshallMaterialAttributesDependency(mad, rawAttributes)
		m.Attributes = mad
	case "package":
		mapp := &MaterialAttributesPackage{}
		unmarshallMaterialAttributesPackage(mapp, rawAttributes)
		m.Attributes = mapp
	case "plugin":
		mapl := &MaterialAttributesPlugin{}
		unmarshallMaterialAttributesPlugin(mapl, rawAttributes)
		m.Attributes = mapl
	default:
		return fmt.Errorf("Unexpected Material type: '%s'", m.Type)
	}

	return nil
}

func unmarshallMaterialFilter(i interface{}) *MaterialFilter {
	return nil
}
