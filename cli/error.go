package cli

import (
	"encoding/json"
	"fmt"
	"github.com/drewsonne/go-gocd/gocd"
)

// JSONCliError describes an error which outputs JSON on the CLI.
type JSONCliError struct {
	ReqType string
	data    dataJSONCliError
	resp    *gocd.APIResponse
}
type dataJSONCliError map[string]interface{}

// NewFlagError creates an error when a flag is missing.
func NewFlagError(flag string) (err error) {
	return fmt.Errorf("'--%s' is missing", flag)
}

// NewCliError creates an error which can be returned from a cli action
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

// Error encodes the error as a JSON string
func (e JSONCliError) Error() string {
	b, err := json.MarshalIndent(e.data, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(b)
}

// ExitCode returns the cli statusin the event of an error
func (e JSONCliError) ExitCode() int {
	if e.resp == nil {
		return 1
	}
	return e.resp.HTTP.StatusCode
}
