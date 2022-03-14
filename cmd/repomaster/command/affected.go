package command

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(affectedCmd)
}

var affectedCmd = &cobra.Command{
	Use:   "affected",
	Short: "List the changes files between two git commits or tags",
	Run: func(cmd *cobra.Command, args []string) {
		// git.Tag
		// a, err := affected.Packages("", "", "", affected.WithExcludeFileGlobs(""))
	},
}
