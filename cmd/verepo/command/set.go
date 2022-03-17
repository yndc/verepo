package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yndc/verepo/pkg/config"
	"github.com/yndc/verepo/pkg/git"
	"github.com/yndc/verepo/pkg/semver"
)

func init() {
	RootCmd.AddCommand(setCmd)
	setCmd.SetUsageTemplate(usageTemplate())
}

var setCmd = &cobra.Command{
	Use:   "set <version>",
	Short: "Set the full version of the given project ID in the repository",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		if !config.Global.DryRun && git.HasUncommittedChanges() {
			return fmt.Errorf("this command cannot be run with uncommitted changes")
		}
		project := args[0]
		current := git.Latest(project)
		next, err := semver.Parse(args[1])
		if err != nil {
			return fmt.Errorf("unable to parse target version (%s): %s", args[1], err.Error())
		}

		if err = git.SetVersion(project, current, next); err != nil {
			return err
		}

		fmt.Printf("%s: %s -> %s\n", project, current.String(), next.String())

		return nil
	},
}
