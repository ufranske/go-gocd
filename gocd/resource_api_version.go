package gocd

import (
	"regexp"
	"strconv"
)

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

func (av ApiVersion) Int() (i int) {
	if string(av) == "" {
		return 0
	}
	matches := apiVersionRegex.FindStringSubmatch(av.String())
	// We match for integers, so this shouldn't fail
	i, _ = strconv.Atoi(matches[1])
	return
}

func (av ApiVersion) LessThan(version ApiVersion) bool {
	return av.Int() < version.Int()
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
	upperBound := version.LessThan(avr.max)
	lowerBound := avr.min.LessThan(version)
	equal := avr.max.Equal(version)
	return upperBound && lowerBound || equal
}
