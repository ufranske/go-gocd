package gocd

import (
	"testing"
	"encoding/json"
	"net/url"
)

func TestResources_MarshalLinks(t *testing.T) {
	var err error
	var j []byte
	expected_json := `{"self":{"href":"http://test.com/"},"next":{"href":"http://test.com/"},"latest":{"href":"http://test.com/"},"oldest":{"href":"http://test.com/"},"doc":{"href":"http://test.com/"},"find":{"href":"http://test.com/"}}`

	demo_url, err := url.Parse("http://test.com/")
	rl := ResponseLinks{
		Self:   demo_url,
		Next:   demo_url,
		Latest: demo_url,
		Oldest: demo_url,
		Doc:    demo_url,
		Find:   demo_url,
	}

	if err != nil {
		t.Error(err)
	}

	j, err = json.Marshal(rl)

	if err != nil {
		t.Error(err)
	}
	if string(j) != expected_json {
		t.Errorf("Expected '%s'. Got '%s'", expected_json, string(j))
	}
}

func TestResources_MarshalPartialLinks(t *testing.T) {
	var err error
	var j []byte
	expected_json := `{"self":{"href":"http://test.com/"},"next":{"href":"http://test.com/"}}`

	demo_url, err := url.Parse("http://test.com/")
	rl := ResponseLinks{
		Self:   demo_url,
		Next:   demo_url,
	}

	if err != nil {
		t.Error(err)
	}

	j, err = json.Marshal(rl)

	if err != nil {
		t.Error(err)
	}
	if string(j) != expected_json {
		t.Errorf("Expected '%s'. Got '%s'", expected_json, string(j))
	}
}

func TestResources_UnmarshalLinks(t *testing.T) {
	l := &ResponseLinks{}
	err := json.Unmarshal([]byte(`{
"self": {
  "href": "https://ci.example.com/go/api/admin/environments/my_environment"
},
"doc": {
  "href": "https://api.gocd.org/#environment-config"
},
"find": {
  "href": "https://ci.example.com/go/api/admin/environments/:environment_name"
}
}`), l)

	if err != nil {
		t.Error(err)
	}

	// Check "self"
	if l.Self.Host != "ci.example.com" {
		t.Errorf("Expected '%s'. Got '%s'.", "ci.example.com", l.Self.Host)
	}
	if l.Self.Path != "/go/api/admin/environments/my_environment" {
		t.Errorf("Expected '%s'. Got '%s'.", "/go/api/admin/environments/my_environment", l.Self.Path)
	}
	if l.Self.Scheme != "https" {
		t.Errorf("Expected '%s'. Got '%s'.", "https", l.Self.Scheme)
	}

	// Check "doc"
	if l.Doc.Host != "api.gocd.org" {
		t.Errorf("Expected '%s'. Got '%s'.", "api.gocd.org", l.Doc.Host)
	}
	if l.Doc.Path != "/" {
		t.Errorf("Expected '%s'. Got '%s'.", "/", l.Doc.Path)
	}
	if l.Doc.Fragment != "environment-config" {
		t.Errorf("Expected '%s'. Got '%s'.", "#environment-config", l.Doc.Fragment)
	}
	if l.Doc.Scheme != "https" {
		t.Errorf("Expected '%s'. Got '%s'.", "https", l.Doc.Scheme)
	}

	// Check "find"
	if l.Find.Host != "ci.example.com" {
		t.Errorf("Expected '%s'. Got '%s'.", "api.gocd.org", l.Doc.Host)
	}
	if l.Find.Path != "/go/api/admin/environments/:environment_name" {
		t.Errorf("Expected '%s'. Got '%s'.", "/", l.Doc.Path)
	}
	if l.Find.Fragment != "" {
		t.Errorf("Expected '%s'. Got '%s'.", "", l.Doc.Fragment)
	}
	if l.Find.Scheme != "https" {
		t.Errorf("Expected '%s'. Got '%s'.", "https", l.Doc.Scheme)
	}

}

func TestResources_UnmarshalLinksPartial(t *testing.T) {
	l := &ResponseLinks{}
	err := json.Unmarshal([]byte(`{
	"self": {
	  "href": "https://ci.example.com/go/api/admin/environments/my_environment"
	}
}`), l)

	if err != nil {
		t.Error(err)
	}

	if l.Self == nil {
		t.Errorf("Expected 'Self' not to be nil.")
	}
	if l.Next != nil {
		t.Errorf("Expected '%s' to be nil. Got '%s'", "Next", l.Next)
	}
	if l.Latest != nil {
		t.Errorf("Expected '%s' to be nil. Got '%s'", "Latest", l.Latest)
	}
	if l.Oldest != nil {
		t.Errorf("Expected '%s' to be nil. Got '%s'", "Oldest", l.Oldest)
	}
	if l.Doc != nil {
		t.Errorf("Expected '%s' to be nil. Got '%s'", "Doc", l.Doc)
	}
	if l.Find != nil {
		t.Errorf("Expected '%s' to be nil. Got '%s'", "Find", l.Find)
	}

}
