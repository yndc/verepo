package git

import (
	"fmt"
	"strings"

	"github.com/flowscan/repomaster-go/pkg/exec"
	"github.com/flowscan/repomaster-go/pkg/semver"
)

func Latest(app string) semver.Parsed {
	o, err := exec.Exec("git", "tag", "-l", app+"/*")
	if err != nil {
		return semver.Invalid()
	}

	if tag, ok := getLatestTag(o); ok {
		return tag
	}
	return semver.Invalid()
}

func SetVersion(app string, from semver.Parsed, to semver.Parsed) error {
	if from.Invalid {
		return fmt.Errorf("current version (%s) is invalid", from.String())
	}
	if semver.Compare(to, from) <= 0 {
		return fmt.Errorf("the target version (%s) is less than the current version (%s)", to.String(), from.String())
	}
	if len(to.Prerelease) == 0 {
		return fmt.Errorf("pre-release tag is required for the target version (%s)", to.String())
	}

	tag := fmt.Sprintf(`%s/%s`, app, to.String())

	exec.SeqExec([][]string{
		{"git", "tag", tag},
		{"git", "push", "origin", tag},
	})

	return nil
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
		if semver.Compare(p, latest) > 0 {
			latest = p
			found = true
		}
	}
	return latest, found
}
