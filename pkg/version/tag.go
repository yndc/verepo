package version

import (
	"fmt"
	"os/exec"
)

func generateTag() string {
	appName := "graphql"
	version := "v0.2.4"
	branch := exec.Command("git", "branch", "--show-current")
	commit := exec.Command("git", "rev-parse", "--short", "HEAD")
	return fmt.Sprintf("%s:%s-%s-%s", appName, version, branch, commit)
}
