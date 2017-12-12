package gocd

import "errors"

func (mas MaterialAttributesSvn) equal(mas2i interface{}) (isEqual bool, err error) {
	var ok bool
	mas2, ok := mas2i.(*MaterialAttributesGit)
	if !ok {
		return false, errors.New("can only compare with same material type")
	}
	urlsEqual := mas.URL == mas2.URL
	destinationEqual := mas.Destination == mas2.Destination

	return urlsEqual && destinationEqual, nil
}

// UnmarshallInterface from a JSON string to a MaterialAttributesSvn struct
func unmarshallMaterialAttributesSvn(mas *MaterialAttributesSvn, i map[string]interface{}) {
	for key, value := range i {
		if value == nil {
			continue
		}
		switch key {
		case "name":
			mas.Name = value.(string)
		case "url":
			mas.URL = value.(string)
		case "username":
			mas.Username = value.(string)
		case "password":
			mas.Password = value.(string)
		case "encrypted_password":
			mas.EncryptedPassword = value.(string)
		case "check_externals":
			mas.CheckExternals = value.(bool)
		case "destination":
			mas.Destination = value.(string)
		case "invert_filter":
			mas.InvertFilter = value.(bool)
		case "auto_update":
			mas.AutoUpdate = value.(bool)
		}
	}
}
