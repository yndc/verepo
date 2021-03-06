package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yndc/verepo/pkg/config"
	"github.com/yndc/verepo/pkg/git"
	"github.com/yndc/verepo/pkg/semver"
)

func init() {
	RootCmd.AddCommand(bumpCmd)
	bumpCmd.Flags().Bool("minor", false, "")
	bumpCmd.Flags().Bool("major", false, "")
	bumpCmd.Flags().StringP("prerelease", "p", "", "")
	bumpCmd.SetUsageTemplate(usageTemplate())
}

var bumpCmd = &cobra.Command{
	Use:   "bump",
	Short: "Bump the project version",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if !config.Global.DryRun && git.HasUncommittedChanges() {
			return fmt.Errorf("this command cannot be run with uncommitted changes")
		}
		project := args[0]
		current := git.Latest(project)
		var next semver.Parsed
		if major, _ := cmd.Flags().GetBool("major"); major {
			next = current.BumpMajor()
		} else if minor, _ := cmd.Flags().GetBool("minor"); minor {
			next = current.BumpMinor()
		} else {
			next = current.BumpPatch()
		}

		if prerelease, _ := cmd.Flags().GetString("prerelease"); len(prerelease) > 0 {
			p, err := semver.ParsePrerelease(prerelease)
			if err != nil {
				return err
			}
			next.Prerelease = p
		} else if len(next.Prerelease) == 0 {
			next.Prerelease = []string{"dev"}
		}

		if err := git.SetVersion(project, current, next); err != nil {
			return err
		}

		fmt.Printf("%s: %s -> %s\n", project, current.String(), next.String())

		return nil
	},
}
