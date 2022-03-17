package command

import (
	"github.com/flowscan/repomaster-go/pkg/config"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.PersistentFlags().BoolVarP(&config.Global.DryRun, "dry-run", "d", false, "Dry run")
	RootCmd.PersistentFlags().BoolVarP(&config.Global.Verbose, "verbose", "v", false, "Verbose output")
}

var RootCmd = &cobra.Command{
	Args: cobra.MinimumNArgs(2),
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd:   true,
		DisableNoDescFlag:   true,
		DisableDescriptions: true,
	},
}
