package cli

import (
	"encoding/json"
	"github.com/drewsonne/go-gocd/gocd"
)

type JSONCliError struct {
	data dataJSONCliError
}
type dataJSONCliError map[string]interface{}

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
	return 1
}
