package changelog

import (
	"bufio"
	"os"
	"strings"

	"github.com/yndc/verepo/pkg/git"
	"github.com/yndc/verepo/pkg/semver"
)

func (d *Document) Write(path string, project string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(f)
	vers := make([]semver.Parsed, 0)

	w.WriteString("# Changelog\n\n")
	if len(d.Description) > 0 {
		w.WriteString(strings.Trim(d.Description, " \n"))
		w.WriteString("\n\n")
		w.Flush()
	}
	if d.Unreleased.Count() > 0 || len(d.Unreleased.Description) > 0 {
		w.WriteString("## [Unreleased]\n\n")
		if len(d.Unreleased.Description) > 0 {
			w.WriteString(strings.Trim(d.Unreleased.Description, " \n") + "\n\n")
		}
		if len(d.Unreleased.Additions) > 0 {
			w.WriteString("### Added\n\n")
			for _, v := range d.Unreleased.Additions {
				w.WriteString("- " + v + "\n")
			}
			w.WriteString("\n")
			w.Flush()
		}
		if len(d.Unreleased.Fixes) > 0 {
			w.WriteString("### Fixed\n\n")
			for _, v := range d.Unreleased.Fixes {
				w.WriteString("- " + v + "\n")
			}
			w.WriteString("\n")
			w.Flush()
		}
		if len(d.Unreleased.Changes) > 0 {
			w.WriteString("### Changed\n\n")
			for _, v := range d.Unreleased.Changes {
				w.WriteString("- " + v + "\n")
			}
			w.WriteString("\n")
			w.Flush()
		}
		if len(d.Unreleased.Removals) > 0 {
			w.WriteString("### Removed\n\n")
			for _, v := range d.Unreleased.Removals {
				w.WriteString("- " + v + "\n")
			}
			w.WriteString("\n")
			w.Flush()
		}
	}
	for _, section := range d.History {
		vers = append(vers, section.Version)
		w.WriteString("## [" + section.Version.VersionStringNoV() + "] - " + section.Date + "\n\n")
		if len(section.Description) > 0 {
			w.WriteString(strings.Trim(section.Description, " \n") + "\n\n")
		}
		if len(section.Additions) > 0 {
			w.WriteString("### Added\n\n")
			for _, v := range section.Additions {
				w.WriteString("- " + v + "\n")
			}
			w.WriteString("\n")
			w.Flush()
		}
		if len(section.Fixes) > 0 {
			w.WriteString("### Fixed\n\n")
			for _, v := range section.Fixes {
				w.WriteString("- " + v + "\n")
			}
			w.WriteString("\n")
			w.Flush()
		}
		if len(section.Changes) > 0 {
			w.WriteString("### Changed\n\n")
			for _, v := range section.Changes {
				w.WriteString("- " + v + "\n")
			}
			w.WriteString("\n")
			w.Flush()
		}
		if len(section.Removals) > 0 {
			w.WriteString("### Removed\n\n")
			for _, v := range section.Removals {
				w.WriteString("- " + v + "\n")
			}
			w.WriteString("\n")
			w.Flush()
		}
	}
	w.Flush()

	origin := git.GetOrigin()
	if d.Unreleased.Count() > 0 && len(d.History) > 0 {
		w.WriteString("[unreleased]: " + origin + "/compare/" + d.History[len(d.History)-1].Version.String() + "...HEAD\n")
		w.Flush()
	}
	for i, v := range vers {
		if i < len(vers)-1 {
			w.WriteString("[" + v.VersionStringNoV() + "]: " + origin + "/compare/" + formatVersion(project, vers[i+1]) + "..." + formatVersion(project, v) + "\n")
		} else {
			w.WriteString("[" + v.VersionStringNoV() + "]: " + origin + "/releases/tag/" + formatVersion(project, v) + "\n")
		}
		w.Flush()
	}

	return nil
}

func formatVersion(project string, version semver.Parsed) string {
	return project + "/" + version.String()
}
