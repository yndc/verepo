package git

import (
	"github.com/flowscan/repomaster-go/pkg/exec"
)

func Commit(m string) error {
	_, err := exec.Exec("git", "commit", "-a", "-m", m)
	if err != nil {
		return err
	}
	return nil
}
