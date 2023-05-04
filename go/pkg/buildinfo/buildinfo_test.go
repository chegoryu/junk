package buildinfo

import (
	"testing"
)

func TestGetProgramVersion(t *testing.T) {
	var tests = []struct {
		GitDescribe            string
		ExpectedProgramVersion string
	}{
		{"unknown", "unknown"},
		{"random", "random"},

		{"heads/main/something", "main/something"},
		{"remotes/main/something", "main/something"},
		{"tags/v1.0.0/something", "v1.0.0/something"},
		{
			"heads/tags/remotes/tags/heads/v1.0.0/something",
			"tags/remotes/tags/heads/v1.0.0/something",
		},
		{
			"tags/heads/remotes/heads/tags/v1.0.0/something",
			"heads/remotes/heads/tags/v1.0.0/something",
		},
		{
			"remotes/heads/tags/heads/remotes/v1.0.0/something",
			"heads/tags/heads/remotes/v1.0.0/something",
		},

		{"main-dirty", "main-modified"},
		{"main-dirty-x", "main-dirty-x"},
		{"heads/main-dirty", "main-modified"},
	}

	for _, test := range tests {
		t.Run(test.GitDescribe, func(t *testing.T) {
			programVersion := getProgramVersion(test.GitDescribe)
			if programVersion != test.ExpectedProgramVersion {
				t.Errorf("expected '%s', got '%s'", test.ExpectedProgramVersion, programVersion)
			}
		})
	}
}
