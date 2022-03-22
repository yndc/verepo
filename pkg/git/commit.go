package git

import (
	"fmt"
	"strings"

	"github.com/yndc/verepo/pkg/exec"
)

func GetCommitsMessages(from string, to string) ([]string, error) {
	o, err := exec.Exec("git", "log", fmt.Sprintf("%s...%s", from, to), "--pretty=format:%s")
	if err != nil {
		return nil, err
	}
	return strings.Split(string(o), "\n"), nil
}

func Commit(m string) error {
	_, err := exec.SeqExec([][]string{
		{"git", "add", "."},
		{"git", "commit", "-a", "-m", m},
	})
	if err != nil {
		return err
	}
	return nil
}
