package command

import (
	"fmt"

	"github.com/flowscan/repomaster-go/pkg/config"
	"github.com/flowscan/repomaster-go/pkg/git"
	"github.com/flowscan/repomaster-go/pkg/semver"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(preCmd)
}

var preCmd = &cobra.Command{
	Use:   "pre",
	Short: "Set the prerelease for the specified app",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		if !config.Global.DryRun && git.HasUncommittedChanges() {
			return fmt.Errorf("this command cannot be run with uncommitted changes")
		}
		app := args[0]
		targetP, err := semver.ParsePrerelease(args[1])
		if err != nil {
			return fmt.Errorf("invalid prerelease: %v", err)
		}
		current := git.Latest(app)
		if current.Invalid {
			return fmt.Errorf("current version (%s) is invalid", current.String())
		}

		if semver.ComparePrerelease(targetP, current.Prerelease) <= 0 {
			fmt.Println(semver.ComparePrerelease(current.Prerelease, targetP))
			return fmt.Errorf(
				"the target pre-release tag (%s) is on a lower precedence over the current pre-release (%s)",
				targetP.String(),
				current.Prerelease.String(),
			)
		}

		next := current
		next.Prerelease = targetP

		if err := git.SetVersion(app, current, next); err != nil {
			return err
		}

		fmt.Printf("%s: %s -> %s\n", app, current.String(), next.String())

		return nil
	},
}
