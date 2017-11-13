package gocd

func (matp *MaterialAttributesPlugin) UnmarshallInterface(i map[string]interface{}) {
	for key, value := range i {
		if value == nil {
			continue
		}
		switch key {
		case "ref":
			matp.Ref = value.(string)
		case "destination":
			matp.Destination = value.(string)
		case "filter":
			continue
		}
	}
}
