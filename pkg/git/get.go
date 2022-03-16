package git

import (
	"os/exec"
	"strings"

	"github.com/flowscan/repomaster-go/pkg/semver"
)

func LatestVer(app string) semver.Parsed {
	o, err := exec.Command("git", "tag", "-l", app+"/*").Output()
	if err != nil {
		return semver.Invalid()
	}

	if tag, ok := getLatestTag(o); ok {
		return tag
	}
	return semver.Invalid()
}

func getLatestTag(o []byte) (semver.Parsed, bool) {
	sp := strings.Split(string(o), "\n")
	found := false
	latest := semver.Parsed{
		Major:      0,
		Minor:      0,
		Patch:      0,
		Prerelease: []string{"dev"},
	}
	for _, s := range sp {
		if len(s) == 0 {
			continue
		}
		sp := strings.Split(s, "/")
		p, err := semver.Parse(sp[1])
		if err != nil {
			continue
		}
		if semver.Compare(latest, p) > 0 {
			latest = p
			found = true
		}
	}
	return latest, found
}
