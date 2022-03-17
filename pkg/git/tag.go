package git

import (
	"fmt"
	"strings"

	"github.com/yndc/verepo/pkg/exec"
	"github.com/yndc/verepo/pkg/semver"
)

func Latest(project string) semver.Parsed {
	o, err := exec.Exec("git", "tag", "-l", project+"/*")
	if err != nil {
		return semver.Zero()
	}

	if tag, ok := getLatestTag(o); ok {
		return tag
	}
	return semver.Zero()
}

func SetVersion(project string, from semver.Parsed, to semver.Parsed) error {
	if from.Invalid {
		return fmt.Errorf("current version (%s) is invalid", from.String())
	}
	if semver.Compare(to, from) <= 0 {
		return fmt.Errorf("the target version (%s) is less than the current version (%s)", to.String(), from.String())
	}
	if len(to.Prerelease) == 0 {
		return fmt.Errorf("pre-release tag is required for the target version (%s)", to.String())
	}

	tag := fmt.Sprintf(`%s/%s`, project, to.String())

	exec.SeqExec([][]string{
		{"git", "tag", tag},
		{"git", "push", "origin", tag},
	})

	return nil
}

func ReleaseVersion(project string, current semver.Parsed) error {
	if current.Invalid {
		return fmt.Errorf("current version (%s) is invalid", current.String())
	}

	current.Prerelease = semver.Prerelease{}

	tag := fmt.Sprintf(`%s/%s`, project, current.String())

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
