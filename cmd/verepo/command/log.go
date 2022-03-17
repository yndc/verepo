package command

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yndc/verepo/pkg/changelog"
	"github.com/yndc/verepo/pkg/git"
)

func init() {
	RootCmd.AddCommand(logCmd)
	logCmd.Flags().Bool("add", false, "")
	logCmd.Flags().Bool("remove", false, "")
	// logCmd.Flags().Bool("deprecate", false, "")
	logCmd.Flags().Bool("fix", false, "")
	// logCmd.Flags().Bool("security", false, "")
	logCmd.Flags().Bool("summary", false, "")
	logCmd.Flags().Bool("commits", false, "")
	logCmd.SetUsageTemplate(usageTemplate())
}

var logCmd = &cobra.Command{
	Use:   "log <message>",
	Short: "Log a new entry into the project's changelog",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		project := args[0]
		doc, err := changelog.Parse("./cmd/" + project + "/changelog.md")
		if err != nil {
			doc = &changelog.Document{}
		}

		if commits, _ := cmd.Flags().GetBool("commits"); commits {
			if doc.Unreleased.Count() > 0 {
				return fmt.Errorf("cannot use --commits since there's unreleased changes in the changelog")
			}
			from := ""
			if len(doc.History) > 0 {
				from = project + "/" + doc.History[len(doc.History)-1].Version.VersionString()
			} else {
				latest := git.Latest(project)
				if latest.Invalid {
					return fmt.Errorf("unable to get previous version")
				}
				from = latest.VersionString()
			}
			logs, err := git.GetCommitsMessages(from, "HEAD")
			if err != nil {
				return fmt.Errorf("trying to get commit messages from %s to HEAD: ", err.Error())
			}
			for _, msg := range logs {
				f := firstWordLowercase(msg)
				msg := capitalize(msg)
				switch f {
				case "add", "added":
					doc.Unreleased.Additions = append(doc.Unreleased.Additions, msg)
				case "fix", "fixed":
					doc.Unreleased.Fixes = append(doc.Unreleased.Fixes, msg)
				case "remove", "removed", "delete", "deleted":
					doc.Unreleased.Removals = append(doc.Unreleased.Removals, msg)
				default:
					doc.Unreleased.Changes = append(doc.Unreleased.Changes, msg)
				}
			}
		} else if len(args) > 1 {
			msg := args[1]
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
			} else if summary, _ := cmd.Flags().GetBool("summary"); summary {
				doc.Unreleased.Description = msg + "\n\n"
			} else {
				doc.Unreleased.Changes = append(doc.Unreleased.Changes, msg)
			}
		} else {
			return fmt.Errorf("message is required, or use --commits flag to add the uncommitted messages into the changelog")
		}

		return doc.Write("./cmd/"+project+"/changelog.md", project)
	},
}

func firstWordLowercase(line string) string {
	return strings.ToLower(strings.Split(line, " ")[0])
}

func capitalize(line string) string {
	sp := strings.Split(line, " ")
	sp[0] = strings.Title(sp[0])
	return strings.Join(sp, " ")
}
