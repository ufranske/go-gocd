package gocd

import (
	"context"
	"strings"
	"strconv"
)

type ServerVersionService service

type ServerVersionParts struct {
	Major int
	Minor int
	Patch int
}

type ServerVersion struct {
	Version      string `json:"version"`
	VersionParts *ServerVersionParts
	BuildNumber  string `json:"build_number"`
	GitSha       string `json:"git_sha"`
	FullVersion  string `json:"full_version"`
	CommitURL    string `json:"commit_url"`
}

// Get retrieves information about a specific plugin.
func (svs *ServerVersionService) Get(ctx context.Context) (v *ServerVersion, resp *APIResponse, err error) {
	v = &ServerVersion{}
	_, resp, err = svs.client.getAction(ctx, &APIClientRequest{
		Path:         "version",
		ResponseBody: v,
		APIVersion:   apiV1,
	})

	if err == nil {
		var major, minor, patch int
		versionParts := strings.Split(v.Version, ".")

		if major, err = strconv.Atoi(versionParts[0]); err != nil {
			return
		}

		if minor, err = strconv.Atoi(versionParts[1]); err != nil {
			return
		}

		if patch, err = strconv.Atoi(versionParts[2]); err != nil {
			return
		}

		v.VersionParts = &ServerVersionParts{
			Major: major,
			Minor: minor,
			Patch: patch,
		}
	}

	return
}
