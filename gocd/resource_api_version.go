package gocd

import "regexp"

const (
	apiV0 = ""
	// Version 1 of the GoCD API.
	apiV1 = "application/vnd.go.cd.v1+json"
	// Version 2 of the GoCD API.
	apiV2 = "application/vnd.go.cd.v2+json"
	// Version 3 of the GoCD API.
	apiV3 = "application/vnd.go.cd.v3+json"
	// Version 4 of the GoCD API.
	apiV4 = "application/vnd.go.cd.v4+json"
	// Version 5 of the GoCD API.
	apiV5 = "application/vnd.go.cd.v5+json"
	// Version 6 of the GoCD API.
	apiV6 = "application/vnd.go.cd.v6+json"
)

type ApiVersion string

type ApiVersionRange struct {
	min ApiVersion
	max ApiVersion
}

var (
	ApiVersion0     = ApiVersion(apiV0)
	ApiVersion1     = ApiVersion(apiV1)
	ApiVersion2     = ApiVersion(apiV2)
	ApiVersion3     = ApiVersion(apiV3)
	ApiVersion4     = ApiVersion(apiV4)
	ApiVersion5     = ApiVersion(apiV5)
	ApiVersion6     = ApiVersion(apiV6)
	apiVersionRegex = regexp.MustCompile("application/vnd.go.cd.v(\\d+)\\+json")
)

func (av ApiVersion) Int() int {
	return 0
}

func (av ApiVersion) LessThan(version ApiVersion) bool {
	return true
}

func (av ApiVersion) Equal(version ApiVersion) bool {
	return string(av) == string(version)
}

func (av ApiVersion) String() string {
	return string(av)
}

func NewApiVersionRange(min ApiVersion, max ApiVersion) *ApiVersionRange {
	return &ApiVersionRange{
		min: min,
		max: max,
	}
}

func (avr *ApiVersionRange) Contains(version ApiVersion) bool {
	return true
}
