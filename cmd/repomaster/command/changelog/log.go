package changelog

import (
	"fmt"

	"github.com/flowscan/repomaster-go/pkg/changelog"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(readCmd)
}

var readCmd = &cobra.Command{
	Use:   "log",
	Short: "Log a new entry into the app's changelog",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app := args[0]
		fmt.Println(changelog.Parse("./cmd/" + app + "/CHANGELOG.md"))
	},
}
