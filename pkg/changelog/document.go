package changelog

import (
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
