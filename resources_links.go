package gocd

import (
	"encoding/json"
	"net/url"
)

type ResponseLinks struct {
	Self   *url.URL
	Next   *url.URL
	Latest *url.URL
	Oldest *url.URL
	Doc    *url.URL
	Find   *url.URL
}

type ResponseEmbedded interface{}

func (l ResponseLinks) MarshalJSON() ([]byte, error) {

	type href struct {
		Href string `json:"href"`
	}
	type linkstruct struct {
		Self   *href `json:"self,omitempty"`
		Next   *href `json:"next,omitempty"`
		Latest *href `json:"latest,omitempty"`
		Oldest *href `json:"oldest,omitempty"`
		Doc    *href `json:"doc,omitempty"`
		Find   *href `json:"find,omitempty"`
	}
	ls := linkstruct{}

	if l.Self != nil {
		ls.Self = &href{Href: l.Self.String()}
	}

	if l.Next != nil {
		ls.Next = &href{Href: l.Next.String()}
	}

	if l.Latest != nil {
		ls.Latest = &href{Href: l.Latest.String()}
	}

	if l.Oldest != nil {
		ls.Oldest = &href{Href: l.Oldest.String()}
	}

	if l.Doc != nil {
		ls.Doc = &href{Href: l.Doc.String()}
	}

	if l.Find != nil {
		ls.Find = &href{Href: l.Find.String()}
	}

	j, err := json.Marshal(ls)
	if err != nil {
		return nil, err
	}

	return j, nil
}

func (l *ResponseLinks) UnmarshalJSON(j []byte) error {
	var dat map[string]map[string]string

	err := json.Unmarshal(j, &dat)

	if err != nil {
		return err
	}

	if dat["self"]["href"] != "" {
		l.Self, err = url.Parse(dat["self"]["href"])
		if err != nil {
			return err
		}
	}
	if dat["next"]["href"] != "" {
		l.Next, err = url.Parse(dat["next"]["href"])
		if err != nil {
			return err
		}
	}
	if dat["latest"]["href"] != "" {
		l.Latest, err = url.Parse(dat["latest"]["href"])
		if err != nil {
			return err
		}
	}
	if dat["oldest"]["href"] != "" {
		l.Oldest, err = url.Parse(dat["oldest"]["href"])
		if err != nil {
			return err
		}
	}
	if dat["doc"]["href"] != "" {
		l.Doc, err = url.Parse(dat["doc"]["href"])
		if err != nil {
			return err
		}
	}
	if dat["find"]["href"] != "" {
		l.Find, err = url.Parse(dat["find"]["href"])
		if err != nil {
			return err
		}
	}

	return nil
}
