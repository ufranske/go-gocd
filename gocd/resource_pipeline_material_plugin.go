package gocd

import "errors"

func (mapp *MaterialAttributesPlugin) equal(mapp2i interface{}) (bool, error) {
	var ok bool
	mapp2, ok := mapp2i.(*MaterialAttributesPlugin)
	if !ok {
		return false, errors.New("Can only compare with same material type.")
	}

	return mapp.Ref == mapp2.Ref &&
			mapp.Destination == mapp2.Destination,
		nil
}

func (mapp *MaterialAttributesPlugin) UnmarshallInterface(i map[string]interface{}) {
	for key, value := range i {
		if value == nil {
			continue
		}
		switch key {
		case "ref":
			mapp.Ref = value.(string)
		case "destination":
			mapp.Destination = value.(string)
		case "filter":
			mapp.Filter = unmarshallMaterialFilter(value)
		case "invert_filter":
			mapp.InvertFilter = value.(bool)
		}
	}
}
