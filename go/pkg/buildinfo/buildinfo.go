package buildinfo

import (
	"strings"
)

var (
	GitDescribe = "unknown"
	Version     = "unknown"
)

func init() {
	if GitDescribe != "unknown" {
		Version = GitDescribe

		for _, prefix := range []string{
			"heads/",
			"tags/",
		} {
			Version = strings.TrimPrefix(Version, prefix)
		}

		if strings.Contains(Version, "-dirty") {
			Version = strings.TrimSuffix(Version, "-dirty")
			Version += "-modified"
		}
	}
}
