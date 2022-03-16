package changelog

import (
	"bufio"
	"os"
)

func (d *Document) Write(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(f)

	w.WriteString("# Changelog\n")
	w.WriteString("\n")
	w.WriteString(d.Description)
	w.WriteString("\n\n")
	if d.Unreleased.Count() > 0 {
		w.WriteString("## [Unreleased]\n\n")
		w.WriteString(d.Unreleased.Description)
		if len(d.Unreleased.Additions) > 0 {
			w.WriteString("### Added\n\n")
			for _, v := range d.Unreleased.Additions {
				w.WriteString("- " + v + "\n")
			}
			w.WriteString("\n")
		}
		if len(d.Unreleased.Fixes) > 0 {
			w.WriteString("### Fixed\n\n")
			for _, v := range d.Unreleased.Fixes {
				w.WriteString("- " + v + "\n")
			}
			w.WriteString("\n")
		}
		if len(d.Unreleased.Changes) > 0 {
			w.WriteString("### Changed\n\n")
			for _, v := range d.Unreleased.Changes {
				w.WriteString("- " + v + "\n")
			}
			w.WriteString("\n")
		}
		if len(d.Unreleased.Removals) > 0 {
			w.WriteString("### Removed\n\n")
			for _, v := range d.Unreleased.Removals {
				w.WriteString("- " + v + "\n")
			}
			w.WriteString("\n")
		}
	}
	for _, section := range d.History {
		if section.Count() > 0 {
			w.WriteString("## [" + section.Version.String() + "] - " + section.Date + "\n\n")
			w.WriteString(section.Description)
			if len(section.Additions) > 0 {
				w.WriteString("### Added\n\n")
				for _, v := range section.Additions {
					w.WriteString("- " + v + "\n")
				}
				w.WriteString("\n")
			}
			if len(section.Fixes) > 0 {
				w.WriteString("### Fixed\n\n")
				for _, v := range section.Fixes {
					w.WriteString("- " + v + "\n")
				}
				w.WriteString("\n")
			}
			if len(section.Changes) > 0 {
				w.WriteString("### Changed\n\n")
				for _, v := range section.Changes {
					w.WriteString("- " + v + "\n")
				}
				w.WriteString("\n")
			}
			if len(section.Removals) > 0 {
				w.WriteString("### Removed\n\n")
				for _, v := range section.Removals {
					w.WriteString("- " + v + "\n")
				}
				w.WriteString("\n")
			}
		}
	}

	return nil
}
