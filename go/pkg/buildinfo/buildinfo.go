package buildinfo

import (
	"strings"
)

var (
	GitDescribe    = "unknown"
	ProgramVersion = "unknown"
)

func getProgramVersion(gitDescribe string) string {
	programVersion := gitDescribe

	for _, prefix := range []string{
		"heads/",
		"remotes/",
		"tags/",
	} {
		if strings.HasPrefix(programVersion, prefix) {
			programVersion = strings.TrimPrefix(programVersion, prefix)
			break
		}
	}

	if strings.HasSuffix(programVersion, "-dirty") {
		programVersion = strings.TrimSuffix(programVersion, "-dirty")
		programVersion += "-modified"
	}

	return programVersion
}

func init() {
	ProgramVersion = getProgramVersion(GitDescribe)
}
