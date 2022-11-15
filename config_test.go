package main

import (
	"path"
	"strings"
	"testing"
)

func TestInvalidConfig(t *testing.T) {
	testCases := []struct {
		path    string
		message string
	}{
		{
			path:    "missing-title.toml",
			message: "index.title",
		},
		{
			path:    "missing-domain.toml",
			message: "domain",
		},
		{
			path:    "missing-repo.toml",
			message: "repo",
		},
		{
			path:    "missing-vcs.toml",
			message: "vcs",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.path, func(t *testing.T) {
			_, err := newConfig(path.Join("fixtures", tC.path))
			if err != nil && !strings.Contains(err.Error(), tC.message) {
				t.Errorf("'%s' does not contains '%s'", err, tC.message)
			}
		})
	}
}

func TestSampleConfigIsValid(t *testing.T) {
	err := initSampleConfig()
	if err != nil {
		t.Error(err)
	}
}

func TestValidConfig(t *testing.T) {
	_, err := newConfig(path.Join("fixtures", "valid.toml"))
	if err != nil {
		t.Error(err)
	}
}
