package git

import "os/exec"

func HasUncommittedChanges() bool {
	o, err := exec.Command("git", "diff", "--name-only").Output()
	if err != nil {
		return true
	}

	if len(o) == 0 {
		return false
	}
	return true
}
