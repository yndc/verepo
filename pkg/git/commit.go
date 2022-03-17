package git

import (
	"fmt"

	"github.com/flowscan/repomaster-go/pkg/exec"
)

func Commit(m string) error {
	o, err := exec.Exec("git", "commit", "-a", "-m", m)
	fmt.Println(string(o))
	fmt.Println(err)
	if err != nil {
		return err
	}
	return nil
}
