package version

import "os/exec"

func generateTag() string {
	branch := exec.Command("git", "branch", "--show-current")
}
