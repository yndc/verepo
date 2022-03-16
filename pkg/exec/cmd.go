package exec

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/flowscan/repomaster-go/pkg/config"
)

func Exec(name string, args ...string) *exec.Cmd {
	if config.Global.DryRun {
		fmt.Printf("exec: %s %s\n", name, strings.Join(args, " "))
		return nil
	}
	return exec.Command(name, args...)
}

func MultiExec(cmds [][]string) []*exec.Cmd {
	res := make([]*exec.Cmd, len(cmds))
	for i, cmd := range cmds {
		if len(cmd) > 1 {
			res[i] = Exec(cmd[0], cmd[1:]...)
		} else {
			res[i] = Exec(cmd[0])
		}
	}
	return res
}
