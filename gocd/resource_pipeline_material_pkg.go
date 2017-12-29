package gocd

import "errors"

func (mapk MaterialAttributesPackage) equal(mapk2i MaterialAttribute) (bool, error) {
	var ok bool
	mapk2, ok := mapk2i.(MaterialAttributesPackage)
	if !ok {
		return false, errors.New("can only compare with same material type")
	}

	return mapk.Ref == mapk2.Ref, nil
}

// UnmarshallInterface from a JSON string to a MaterialAttributesPackage struct
func unmarshallMaterialAttributesPackage(mapk *MaterialAttributesPackage, i map[string]interface{}) {
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
