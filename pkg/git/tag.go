package git

import (
	"fmt"

	"github.com/flowscan/repomaster-go/pkg/exec"
	"github.com/flowscan/repomaster-go/pkg/semver"
)

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

	exec.MultiExec([][]string{
		{"git", "tag", tag},
		{"git", "push", "origin", tag},
	})

	return nil
}
