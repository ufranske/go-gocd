package gocd

import (
	"encoding/json"
	"errors"
	"strconv"
)

// JSONString returns a string of this stage as a JSON object.
func (j *Job) JSONString() (string, error) {
	err := j.Validate()
	if err != nil {
		return "", err
	}

	bdy, err := json.MarshalIndent(j, "", "  ")
	return string(bdy), err
}

// Validate a job structure has non-nil values on correct attributes
func (j *Job) Validate() (err error) {
	if j.Name == "" {
		err = errors.New("`gocd.Jobs.Name` is empty")
	}
	return
}

// UnmarshalJSON and handle integers, null, and never
func (tf *TimeoutField) UnmarshalJSON(b []byte) (err error) {
	value := string(b)
	var valInt int

	if value == `"never"` || value == `"null"` {
		valInt = 0
	} else {
		valInt, err = strconv.Atoi(value)

	}
	return nil
}
