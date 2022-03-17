package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yndc/verepo/pkg/git"
)

func init() {
	RootCmd.AddCommand(getCmd)
	getCmd.SetUsageTemplate(usageTemplate())
}

var getCmd = &cobra.Command{
	Use:   "version",
	Short: "Get the latest version of the given project ID in the repository",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		v := git.Latest(args[0])
		if v.Invalid {
			fmt.Printf("Version for %s not found or invalid\n", args[0])
		} else {
			fmt.Println(v.String())
		}
	},
}
