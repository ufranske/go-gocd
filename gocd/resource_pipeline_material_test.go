package gocd

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func testResourceMaterial(t *testing.T) {
	t.Run("Equality", testMaterialEquality)
	t.Run("AttributeEquality", testMaterialAttributeEquality)
	t.Run("FilterUnmarshall", testMaterialAttributeUnmarshall)
}

func testMaterialEquality(t *testing.T) {
	s1 := Material{
		Type: "git",
		Attributes: MaterialAttributesGit{
			URL: "https://github.com/gocd/gocd",
		},
	}

	s2 := Material{
		Type: "git",
		Attributes: MaterialAttributesGit{
			Name: "gocd-src",
			URL:  "https://github.com/gocd/gocd",
		},
	}
	ok, err := s1.Equal(&s2)
	assert.Nil(t, err)
	assert.True(t, ok)
}

func testMaterialAttributeEquality(t *testing.T) {
	for i, test := range []struct {
		a      MaterialAttribute
		b      MaterialAttribute
		result bool
	}{
		{a: MaterialAttributesGit{}, b: MaterialAttributesGit{}, result: true},
		{a: MaterialAttributesSvn{}, b: MaterialAttributesSvn{}, result: true},
		{a: MaterialAttributesHg{}, b: MaterialAttributesHg{}, result: true},
		{a: MaterialAttributesP4{}, b: MaterialAttributesP4{}, result: true},
		{a: MaterialAttributesTfs{}, b: MaterialAttributesTfs{}, result: true},
		{a: MaterialAttributesDependency{}, b: MaterialAttributesDependency{}, result: true},
		{a: MaterialAttributesPackage{}, b: MaterialAttributesPackage{}, result: true},
		{a: MaterialAttributesPlugin{}, b: MaterialAttributesPlugin{}, result: true},
		{
			a:      MaterialAttributesGit{},
			b:      MaterialAttributesGit{URL: "https://github.com/drewsonne/go-gocd"},
			result: false,
		},
		{
			a:      MaterialAttributesGit{URL: "https://github.com/drewsonne/go-gocd"},
			b:      MaterialAttributesGit{URL: "https://github.com/drewsonne/go-gocd", Branch: "feature/branch"},
			result: false,
		},
		{a: MaterialAttributesGit{Branch: ""}, b: MaterialAttributesGit{Branch: "master"}, result: true},
		{a: MaterialAttributesGit{Branch: "master"}, b: MaterialAttributesGit{Branch: ""}, result: true},
		{a: MaterialAttributesGit{Branch: ""}, b: MaterialAttributesGit{Branch: ""}, result: true},
		{a: MaterialAttributesGit{Branch: "master"}, b: MaterialAttributesGit{Branch: "master"}, result: true},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ok, err := test.a.equal(test.b)
			assert.Equal(t, ok, test.result)
			assert.Nil(t, err)
		})
	}
}

func testMaterialAttributeUnmarshall(t *testing.T) {
	m := Material{}
	for i, test := range []struct {
		source   string
		expected MaterialAttribute
	}{
		{source: `{"type": "git"}`, expected: &MaterialAttributesGit{}},
		{source: `{"type": "svn"}`, expected: &MaterialAttributesSvn{}},
		{source: `{"type": "hg"}`, expected: &MaterialAttributesHg{}},
		{source: `{"type": "p4"}`, expected: &MaterialAttributesP4{}},
		{source: `{"type": "tfs"}`, expected: &MaterialAttributesTfs{}},
		{source: `{"type": "dependency"}`, expected: &MaterialAttributesDependency{}},
		{source: `{"type": "package"}`, expected: &MaterialAttributesPackage{}},
		{source: `{"type": "plugin"}`, expected: &MaterialAttributesPlugin{}},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			m.UnmarshalJSON([]byte(test.source))
			assert.IsType(t, test.expected, m.Attributes)
		})
	}
}
