package command

import (
	"fmt"
	"os"

	"github.com/flowscan/repomaster-go/pkg/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&config.Global.DryRun, "dry-run", "d", false, "Dry run")
	rootCmd.PersistentFlags().BoolVarP(&config.Global.Verbose, "verbose", "v", false, "Verbose output")
}

var rootCmd = &cobra.Command{
	Use:   "repomaster",
	Short: "RepoMaster: A tool for managing Golang repositories",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
