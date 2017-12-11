package gocd

import "errors"

func (mapk *MaterialAttributesPackage) equal(mapk2i interface{}) (bool, error) {
	var ok bool
	mapk2, ok := mapk2i.(*MaterialAttributesPackage)
	if !ok {
		return false, errors.New("Can only compare with same material type.")
	}

	return mapk.Ref == mapk2.Ref, nil
}

func (mapk *MaterialAttributesPackage) UnmarshallInterface(i map[string]interface{}) {
	for key, value := range i {
		if value == nil {
			continue
		}
		switch key {
		case "ref":
			mapk.Ref = value.(string)
		}
	}
}
