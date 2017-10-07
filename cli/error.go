package cli

import (
	"encoding/json"
	"github.com/drewsonne/go-gocd/gocd"
	"fmt"
)

type JSONCliError struct {
	ReqType string
	data    dataJSONCliError
	resp    *gocd.APIResponse
}
type dataJSONCliError map[string]interface{}

func NewFlagError(flag string) (err error) {
	return fmt.Errorf("'--%s' is missing", flag)
}

func NewCliError(reqType string, hr *gocd.APIResponse, err error) (jerr JSONCliError) {
	data := dataJSONCliError{
		"Error": err.Error(),
	}
	if hr != nil {
		b1, _ := json.Marshal(hr.HTTP.Header)
		b2, _ := json.Marshal(hr.Request.HTTP.Header)
		data["Error"] = "An error occurred while retrieving the resource."
		data["Status"] = hr.HTTP.StatusCode
		data["ResponseHeader"] = string(b1)
		data["ResponseBody"] = hr.Body
		data["RequestBody"] = hr.Request.Body
		data["RequestEndpoint"] = hr.Request.HTTP.URL.String()
		data["RequestHeader"] = string(b2)

	} else {
		data["Request"] = reqType
	}
	return JSONCliError{
		data: data,
		resp: hr,
	}
}

func (e JSONCliError) Error() string {
	b, err := json.MarshalIndent(e.data, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(b)
}

func (e JSONCliError) ExitCode() int {
	if e.resp == nil {
		return 1
	}
	return e.resp.HTTP.StatusCode
}
