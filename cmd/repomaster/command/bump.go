package command

import (
	"fmt"

	"github.com/flowscan/repomaster-go/pkg/exec"
	"github.com/flowscan/repomaster-go/pkg/git"
	"github.com/flowscan/repomaster-go/pkg/semver"
	"github.com/spf13/cobra"
)

var Level int

func init() {
	rootCmd.AddCommand(bumpCmd)
	bumpCmd.Flags().Bool("minor", false, "")
	bumpCmd.Flags().Bool("major", false, "")
}

var bumpCmd = &cobra.Command{
	Use:   "bump",
	Short: "Bump the version of the given app ID in the repository",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if git.HasUncommittedChanges() {
			return fmt.Errorf("this command cannot be run with uncommitted changes")
		}
		app := args[0]
		current := git.LatestVer(app)
		if current.Invalid {
			return fmt.Errorf("current version (%s) is invalid", current.String())
		}

		var next semver.Parsed
		if major, _ := cmd.Flags().GetBool("major"); major {
			next = current.BumpMajor()
		} else if minor, _ := cmd.Flags().GetBool("minor"); minor {
			next = current.BumpMinor()
		} else {
			next = current.BumpPatch()
		}

		tag := fmt.Sprintf(`%s/%s`, app, next.String())

		res := exec.MultiExec([][]string{
			{"git", "tag", tag},
			{"git", "push", "origin", tag},
		})
		for _, r := range res {
			o, err := r.Output()
			fmt.Println(o)
			fmt.Println(err)
		}

		fmt.Printf("%s: %s -> %s\n", app, current.String(), next.String())

		return nil
	},
}
