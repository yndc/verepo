package command

import (
	"fmt"

	"github.com/flowscan/repomaster-go/pkg/exec"
	"github.com/flowscan/repomaster-go/pkg/git"
	"github.com/flowscan/repomaster-go/pkg/semver"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setCmd)
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the version of the given app ID in the repository",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		if git.HasUncommittedChanges() {
			return fmt.Errorf("this command cannot be run with un-committed changes")
		}
		app := args[0]
		target, err := semver.Parse(args[1])
		if err != nil {
			return fmt.Errorf("unable to parse target version (%s): ", args[1], err.Error())
		}
		current := git.LatestVer(app)
		if current.Invalid {
			return fmt.Errorf("current version (%s) is invalid", current.String())
		}

		if len(target.Prerelease) == 0 {
			return fmt.Errorf("pre-release tag is required the target version (%s)", target.String())
		}

		if semver.Compare(current, target) <= 0 {
			return fmt.Errorf("the target version (%s) is less than the current version (%s)", target.String(), current.String())
		}

		tag := fmt.Sprintf(`"%s/%s"`, app, target.String())

		exec.MultiExec([][]string{
			{"git", "tag", tag},
			{"git", "push", "origin", tag},
		})

		fmt.Printf("%s: %s -> %s\n", app, current.String(), target.String())

		return nil
	},
}
