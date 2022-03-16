package command

import (
	"fmt"

	"github.com/flowscan/repomaster-go/pkg/git"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the latest version of the given app ID in the repository",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		v := git.LatestVer(args[0])
		if v.Invalid {
			fmt.Println("invalid version format")
		} else {
			fmt.Println(v.String())
		}
	},
}
