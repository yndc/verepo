package changelog

import (
	"github.com/spf13/cobra"
)

func init() {
}

var RootCmd = &cobra.Command{
	Use:   "changelog",
	Short: "Manage changelogs",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
