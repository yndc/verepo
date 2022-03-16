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

	if len(o) == 0 {
		return semver.Invalid()
	}

	sp := strings.Split(string(o), "/")
	if len(sp) != 2 {
		return semver.Invalid()
	}
	v, _ := semver.Parse(strings.Trim(sp[1], "\n "))
	return v
}
