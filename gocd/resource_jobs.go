package gocd

import (
	"encoding/json"
	"errors"
)

// JSONString returns a string of this stage as a JSON object.
func (j *Job) JSONString() (body string, err error) {
	err = j.Validate()
	if err != nil {
		return
	}

	bdy, err := json.MarshalIndent(j, "", "  ")
	body = string(bdy)

	return
}

// Validate a job structure has non-nil values on correct attributes
func (j *Job) Validate() (err error) {
	if j.Name == "" {
		err = errors.New("`gocd.Jobs.Name` is empty")
	}
	return
}
