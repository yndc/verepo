package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yndc/verepo/pkg/project"
)

func init() {
	multiCmd.Flags().BoolP("list", "l", false, "List all projects in this repository")
}

var multiCmd = &cobra.Command{
	Use: "verepo <project> <command>",
	Args: func(cmd *cobra.Command, args []string) error {

		return nil
	},
	DisableFlagParsing: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		projects, err := project.GetAll()
		if err != nil {
			return fmt.Errorf("unable to list projects, make sure you're on the project root directory: %s", err.Error())
		}
		if len(args) == 0 {
			return fmt.Errorf("<project> is required")
		} else {
			if args[0] == "--list" || args[0] == "-l" {
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
