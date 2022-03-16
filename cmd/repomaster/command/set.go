package command

import (
	"fmt"

	"github.com/flowscan/repomaster-go/pkg/config"
	"github.com/flowscan/repomaster-go/pkg/git"
	"github.com/flowscan/repomaster-go/pkg/semver"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setCmd)
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the full version of the given app ID in the repository",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		if !config.Global.DryRun && git.HasUncommittedChanges() {
			return fmt.Errorf("this command cannot be run with uncommitted changes")
		}
		app := args[0]
		current := git.Latest(app)
		next, err := semver.Parse(args[1])
		if err != nil {
			return fmt.Errorf("unable to parse target version (%s): %s", args[1], err.Error())
		}

		if err = git.SetVersion(app, current, next); err != nil {
			return err
		}

		fmt.Printf("%s: %s -> %s\n", app, current.String(), next.String())

		return nil
	},
}
