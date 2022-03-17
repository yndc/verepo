package exec

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/yndc/verepo/pkg/config"
)

func Exec(name string, args ...string) ([]byte, error) {
	if config.Global.DryRun {
		fmt.Printf("exec: %s %s\n", name, strings.Join(args, " "))
		return nil, nil
	}
	return exec.Command(name, args...).Output()
}

// execute the given commands sequentially, stops on the first error
func SeqExec(cmds [][]string) ([][]byte, error) {
	outputs := make([][]byte, len(cmds))
	var err error
	for i, cmd := range cmds {
		if len(cmd) > 1 {
			outputs[i], err = Exec(cmd[0], cmd[1:]...)
		} else {
			outputs[i], err = Exec(cmd[0])
		}
		if err != nil {
			return outputs, err
		}
	}
	return outputs, nil
}
