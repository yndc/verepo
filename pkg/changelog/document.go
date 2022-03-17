package changelog

import (
	"time"

	"github.com/flowscan/repomaster-go/pkg/semver"
)

type Document struct {
	Description string
	Unreleased  Section
	History     []HistoricalSection
}

type Section struct {
	Description string
	Additions   []string
	Fixes       []string
	Removals    []string
	Changes     []string
}

func (s Section) Count() int {
	return len(s.Additions) + len(s.Fixes) + len(s.Removals) + len(s.Changes)
}

type HistoricalSection struct {
	Section
	Version semver.Parsed
	Date    string
}

func (d *Document) Release(ver semver.Parsed) {
	// newSection := Section{}
	// newSection.Description = d.Unreleased.Description
	// newSection.Additions = d.Unreleased.Additions
	// newSection.Changes = d.Unreleased.Changes
	// newSection.Fixes = d.Unreleased.Fixes
	// newSection.Removals = d.Unreleased.Removals
	prev := d.Unreleased
	d.History = append([]HistoricalSection{{
		Version: ver.Version(),
		Date:    time.Now().Format("2006-01-02"),
		Section: prev,
	}}, d.History...)
	d.Unreleased = Section{}
}
