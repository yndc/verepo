package command

import (
	"fmt"
	"os"

	"github.com/flowscan/repomaster-go/pkg/config"
	"github.com/flowscan/repomaster-go/pkg/project"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.PersistentFlags().BoolVarP(&config.Global.DryRun, "dry-run", "d", false, "Dry run")
	RootCmd.PersistentFlags().BoolVarP(&config.Global.Verbose, "verbose", "v", false, "Verbose output")
	RootCmd.SetUsageTemplate(usageTemplate())
	topCmd.Flags().Bool("list", false, "List all projects in this repository")
}

var RootCmd = &cobra.Command{
	Short: "RepoMaster: A tool for managing Golang repositories",
	Args:  cobra.MinimumNArgs(2),
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd:   true,
		DisableNoDescFlag:   true,
		DisableDescriptions: true,
	},
}

var topCmd = &cobra.Command{
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
				if project.ID == args[0] {
					if len(args) == 1 {
						args = nil
					} else {
						args = []string{args[1], args[0]}
						if len(args) > 2 {
							args = append(args, args[2:]...)
						}
					}
					RootCmd.SetArgs(args)
					return RootCmd.Execute()
				}
			}
			return fmt.Errorf("project \"%s\" not found at %s", args[0], "./cmd/"+args[0])
		}
	},
}

func Execute() {
	if err := topCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
