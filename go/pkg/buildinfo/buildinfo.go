package buildinfo

import (
	"strings"
)

var (
	GitDescribe = "unknown"
	Version     = "unknown"
)

func getVersion(gitDescribe string) string {
	version := gitDescribe

	if version != "unknown" {
		for _, prefix := range []string{
			"heads/",
			"remotes/",
			"tags/",
		} {
			if strings.HasPrefix(version, prefix) {
				version = strings.TrimPrefix(version, prefix)
				break
			}
		}

		if strings.Contains(version, "-dirty") {
			version = strings.TrimSuffix(version, "-dirty")
			version += "-modified"
		}
	}

	return version
}

func init() {
	Version = getVersion(GitDescribe)
}
