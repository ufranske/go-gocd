package gocd

import "net/url"

type linkField map[string]map[string]string
type linkHref struct {
	H string `json:"href"`
}

func unmarshallLinkField(d linkField, field string, destination **url.URL) error {
	var e error
	if h := d[field]["href"]; h != "" {
		*destination, e = url.Parse(h)
		return e
	}
	return nil
}
