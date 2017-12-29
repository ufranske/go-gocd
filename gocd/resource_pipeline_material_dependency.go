package gocd

import (
	"errors"
	"fmt"
)

func (mad MaterialAttributesDependency) equal(mad2i interface{}) (bool, error) {
	var ok bool
	mad2, ok := mad2i.(MaterialAttributesDependency)
	if !ok {
		return false, errors.New("can only compare with same material type")
	}
	return mad.Pipeline == mad2.Pipeline &&
			mad.Stage == mad2.Stage,
		nil
}

// UnmarshallInterface for a MaterialAttribute struct to be turned into a json string
func unmarshallMaterialAttributesDependency(mad *MaterialAttributesDependency, i map[string]interface{}) {
	for key, value := range i {
		if value == nil {
			continue
		}
		switch key {
		case "name":
			mad.Name = value.(string)
		case "pipeline":
			mad.Pipeline = value.(string)
		case "stage":
			mad.Stage = value.(string)
		case "auto_update":
			mad.AutoUpdate = value.(bool)
		default:
			fmt.Println(value)
			fmt.Println(key)
		}
	}
}
