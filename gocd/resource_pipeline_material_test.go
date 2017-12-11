package gocd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testResourceMaterial(t *testing.T) {
	t.Run("Equality", testMaterialEquality)
	t.Run("AttributeEquality", testMaterialAttributeEquality)
	t.Run("FilterUnmarshall", testMaterialFilterUnmarshall)
}

func testMaterialEquality(t *testing.T) {
	s1 := Material{
		Type: "git",
		Attributes: &MaterialAttributesGit{
			URL: "https://github.com/gocd/gocd",
		},
	}

	s2 := Material{
		Type: "git",
		Attributes: &MaterialAttributesGit{
			Name: "gocd-src",
			URL:  "https://github.com/gocd/gocd",
		},
	}
	ok, err := s1.Equal(&s2)
	assert.Nil(t, err)
	assert.True(t, ok)
}

func testMaterialAttributeEquality(t *testing.T) {
	a1 := MaterialAttributesGit{}
	a2 := MaterialAttributesGit{}
	ok, err := a1.equal(&a2)
	assert.Nil(t, err)
	assert.True(t, ok)

	a2.URL = "https://github.com/drewsonne/go-gocd"
	ok, err = a1.equal(&a2)
	assert.Nil(t, err)
	assert.False(t, ok)

	a1.URL = "https://github.com/drewsonne/go-gocd"
	a2.Branch = "feature/branch"
	ok, err = a1.equal(&a2)
	assert.Nil(t, err)
	assert.False(t, ok)

	for _, branchCombo := range [][]string{
		{"", "master"},
		{"master", ""},
		{"", ""},
		{"master", "master"},
	} {
		a1.Branch = branchCombo[0]
		a2.Branch = branchCombo[1]
		ok, err = a1.equal(&a2)
		assert.Nil(t, err)
		assert.True(t, ok)
	}
}


func testMaterialFilterUnmarshall(t *testing.T) {

}