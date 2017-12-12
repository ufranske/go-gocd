package gocd

import "errors"

func (mhg MaterialAttributesHg) equal(mhg2i interface{}) (bool, error) {
	var ok bool
	mhg2, ok := mhg2i.(*MaterialAttributesHg)
	if !ok {
		return false, errors.New("can only compare with same material type")
	}
	urlsEqual := mhg.URL == mhg2.URL
	destEqual := mhg.Destination == mhg2.Destination

	return urlsEqual && destEqual, nil
}

// UnmarshallInterface from a JSON string to a MaterialAttributesHg struct
func unmarshallMaterialAttributesHg(mhg *MaterialAttributesHg, i map[string]interface{}) {
	for key, value := range i {
		if value == nil {
			continue
		}
		switch key {
		case "name":
			mhg.Name = value.(string)
		case "url":
			mhg.URL = value.(string)
		case "destination":
			mhg.Destination = value.(string)
		case "invert_filter":
			mhg.InvertFilter = value.(bool)
		case "auto_update":
			mhg.AutoUpdate = value.(bool)
		case "filter":
			mhg.Filter = unmarshallMaterialFilter(value)
		}
	}
}
