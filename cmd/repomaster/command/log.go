package command

import (
	"github.com/flowscan/repomaster-go/pkg/changelog"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(logCmd)
	logCmd.Flags().Bool("add", false, "")
	logCmd.Flags().Bool("remove", false, "")
	logCmd.Flags().Bool("deprecate", false, "")
	logCmd.Flags().Bool("fix", false, "")
	logCmd.Flags().Bool("security", false, "")
	logCmd.Flags().Bool("summary", false, "")
	logCmd.SetUsageTemplate(usageTemplate())
}

var logCmd = &cobra.Command{
	Use:   "log <message>",
	Short: "Log a new entry into the project's changelog",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		project := args[0]
		msg := args[1]

		doc, err := changelog.Parse("./cmd/" + project + "/changelog.md")
		if err != nil {
			doc = &changelog.Document{}
		}

		if add, _ := cmd.Flags().GetBool("add"); add {
			doc.Unreleased.Additions = append(doc.Unreleased.Additions, msg)
		} else if remove, _ := cmd.Flags().GetBool("remove"); remove {
			doc.Unreleased.Removals = append(doc.Unreleased.Removals, msg)
		} else if deprecate, _ := cmd.Flags().GetBool("deprecate"); deprecate {
			doc.Unreleased.Changes = append(doc.Unreleased.Changes, msg)
		} else if fix, _ := cmd.Flags().GetBool("fix"); fix {
			doc.Unreleased.Fixes = append(doc.Unreleased.Fixes, msg)
		} else if security, _ := cmd.Flags().GetBool("security"); security {
			doc.Unreleased.Fixes = append(doc.Unreleased.Fixes, msg)
		} else if security, _ := cmd.Flags().GetBool("summary"); security {
			doc.Unreleased.Description = msg + "\n\n"
		} else {
			doc.Unreleased.Changes = append(doc.Unreleased.Changes, msg)
		}

		doc.Write("./cmd/"+project+"/changelog.md", project)
	},
}
