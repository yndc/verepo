package command

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/flowscan/repomaster-go/pkg/config"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.PersistentFlags().BoolVarP(&config.Global.DryRun, "dry-run", "d", false, "Dry run")
	RootCmd.PersistentFlags().BoolVarP(&config.Global.Verbose, "verbose", "v", false, "Verbose output")
	RootCmd.SetUsageTemplate(usageTemplate())
}

var RootCmd = &cobra.Command{
	Short: "RepoMaster: A tool for managing Golang repositories",
	Args:  cobra.MinimumNArgs(2),
}

func Execute() {
	wrapCmd := &cobra.Command{
		Use: "repomaster <app> <command>",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("<app> is required")
			} else {
				if _, err := os.Stat("./cmd/" + args[0]); errors.Is(err, os.ErrNotExist) {
					return fmt.Errorf("app \"%s\" not found at %s", args[0], "./cmd/"+args[0])
				}
				if len(args) == 1 {
					args = nil
				} else {
					args = []string{args[1], args[0]}
					if len(args) > 2 {
						args = append(args, args[2:]...)
					}
				}
			}

			RootCmd.SetArgs(args)
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			RootCmd.Execute()
		},
		Example: strings.Join([]string{
			`repomaster appname bump`,
			`repomaster appname log --fix "Fixes bug #1234"`,
			`repomaster appname log --add "Added feature XYZ"`,
			`repomaster appname release`,
		}, "\n"),
	}
	if err := wrapCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
