package command

import (
	"fmt"

	"github.com/flowscan/repomaster-go/pkg/app"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all applications with their versions under the `cmd` folder",
	Run: func(cmd *cobra.Command, args []string) {
		apps, err := app.GetAll()
		if err != nil {
			PrintError(err)
		}
		for _, a := range apps {
			fmt.Printf("- %s\n", a.ID)
			if a.Name != a.ID {
				fmt.Printf("  name: %s\n", a.Name)
			}
			fmt.Printf("  version: %s\n", a.Version)
		}
	},
}
