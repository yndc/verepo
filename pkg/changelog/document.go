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

type HistoricalSection struct {
	Section
	Version semver.Parsed
	Date    string
}
