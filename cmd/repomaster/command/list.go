package command

import (
	"fmt"

	"github.com/flowscan/repomaster-go/pkg/app"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all applications with their versions under the `cmd` folder",
	RunE: func(cmd *cobra.Command, args []string) error {
		apps, err := app.GetAll()
		if err != nil {
			return err
		}
		if len(apps) > 0 {
			fmt.Printf("apps:\n")
			for _, a := range apps {
				fmt.Printf("- %s:%s\n", a.ID, a.Version.String())
			}
		} else {
			fmt.Println("no apps found")
		}
		return nil
	},
}
