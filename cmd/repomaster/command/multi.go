package command

import (
	"fmt"
	"os"

	"github.com/flowscan/repomaster-go/pkg/project"
	"github.com/spf13/cobra"
)

func init() {
	multiCmd.Flags().Bool("list", false, "List all projects in this repository")
}

var multiCmd = &cobra.Command{
	Use: "repomaster <project> <command>",
	Args: func(cmd *cobra.Command, args []string) error {

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		projects, err := project.GetAll()
		if list, _ := cmd.Flags().GetBool("list"); list {
			if err != nil {
				return err
			}
			if len(projects) > 0 {
				fmt.Printf("projects:\n")
				for _, a := range projects {
					fmt.Printf("- %s:%s\n", a.ID, a.Version.String())
				}
			} else {
				fmt.Println("no projects found in this repository")
			}
			return nil
		}
		if len(args) == 0 {
			return fmt.Errorf("<project> is required")
		} else {
			for _, project := range projects {
				var childArgs []string
				if project.ID == args[0] {
					if len(args) == 1 {
						childArgs = nil
					} else {
						childArgs = []string{args[1], args[0]}
						if len(args) > 2 {
							childArgs = append(childArgs, args[2:]...)
						}
					}
					RootCmd.SetArgs(childArgs)
					RootCmd.SetUsageTemplate(usageTemplate())
					for _, c := range RootCmd.Commands() {
						c.SetUsageTemplate(usageTemplate())
					}
					err = RootCmd.Execute()
					if err != nil {
						os.Exit(1)
					}
					return nil
				}
			}
			return fmt.Errorf("project \"%s\" not found at %s", args[0], "./cmd/"+args[0])
		}
	},
}

func Execute() {
	if err := multiCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
