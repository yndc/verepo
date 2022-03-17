package command

import (
	"fmt"

	"github.com/flowscan/repomaster-go/pkg/changelog"
	"github.com/flowscan/repomaster-go/pkg/config"
	"github.com/flowscan/repomaster-go/pkg/git"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(releaseCmd)
}

var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Remove the pre-release tag of the current version of the given app ID in the repository",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if !config.Global.DryRun && git.HasUncommittedChanges() {
			return fmt.Errorf("this command cannot be run with uncommitted changes")
		}
		app := args[0]
		current := git.Latest(app)
		if len(current.Prerelease) == 0 {
			return fmt.Errorf("current version (%s) is already released", current.String())
		}

		// update the changelog
		changelogPath := "./cmd/" + app + "/changelog.md"
		doc, err := changelog.Parse(changelogPath)
		if err != nil {
			doc = &changelog.Document{}
		}

		err = doc.Write(changelogPath, app)
		if err != nil {
			return err
		}

		// add the git tag
		if err := git.ReleaseVersion(app, current); err != nil {
			return err
		}

		fmt.Printf("Released version %s:%s\n", app, current.VersionString())

		return nil
	},
}
