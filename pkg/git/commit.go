package git

import (
	"strings"

	"github.com/flowscan/repomaster-go/pkg/exec"
)

func GetCommitsMessages(from string, to string) ([]string, error) {
	o, err := exec.Exec("git", "log", "repomaster/v1.0.3-dev...HEAD", "--pretty=format:%%s")
	if err != nil {
		return nil, err
	}
	return strings.Split(string(o), "\n"), nil
}

func Commit(m string) error {
	_, err := exec.Exec("git", "commit", "-a", "-m", m)
	if err != nil {
		return err
	}
	return nil
}
