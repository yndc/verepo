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
	releaseCmd.SetUsageTemplate(usageTemplate())
}

var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Remove the pre-release tag of the current version of the given project ID in the repository",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if !config.Global.DryRun && git.HasUncommittedChanges() {
			return fmt.Errorf("this command cannot be run with uncommitted changes")
		}
		project := args[0]
		current := git.Latest(project)
		if len(current.Prerelease) == 0 {
			return fmt.Errorf("current version (%s) is already released", current.String())
		}

		// update the changelog
		changelogPath := "./cmd/" + project + "/changelog.md"
		doc, err := changelog.Parse(changelogPath)
		if err != nil {
			doc = &changelog.Document{}
		}
		doc.Release(current)
		err = doc.Write(changelogPath, project)
		if err != nil {
			return err
		}

		// commit the changelog change
		err = git.Commit(fmt.Sprintf("Release %s/%s", project, current.VersionString()))
		if err != nil {
			return err
		}

		// add the git tag
		if err := git.ReleaseVersion(project, current); err != nil {
			return err
		}
		fmt.Printf("Released %s/%s\n", project, current.VersionString())

		return nil
	},
}
