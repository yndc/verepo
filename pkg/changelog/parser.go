package changelog

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/flowscan/repomaster-go/pkg/semver"
)

type section int
type subSection int

const (
	sectionInvalid = iota
	sectionHead
	sectionUnreleased
	sectionHistory
)

const (
	subSectionInvalid = -1
	subSectionHeader  = iota
	subSectionAdd
	subSectionFix
	subSectionChange
	subSectionRemove
)

type parserState struct {
	doc         *Document
	section     section
	subSection  subSection
	pendingItem string
}

func (s *parserState) tryAddChange(line string) bool {
	if s.subSection != subSectionInvalid {
		sp := strings.Split(line, "- ")
		if strings.HasPrefix(line, "- ") && len(sp) == 2 {
			// new entry
			item := sp[1]
			if len(s.pendingItem) > 0 {
				s.addItem(s.pendingItem)
				s.pendingItem = ""
			}
			s.pendingItem = item
			return true
		} else if len(s.pendingItem) > 0 {
			if line == "" {
				s.addItem(s.pendingItem)
				s.pendingItem = ""
			} else {
				s.pendingItem = s.pendingItem + " " + strings.Trim(line, " \n")
			}
			return true
		}
	}
	return false
}

func (s *parserState) addItem(item string) {
	var section *Section
	if s.section == sectionUnreleased {
		section = &s.doc.Unreleased
	} else {
		section = &s.doc.History[len(s.doc.History)-1].Section
	}
	switch s.subSection {
	case subSectionAdd:
		section.Additions = append(section.Additions, item)
	case subSectionFix:
		section.Fixes = append(section.Fixes, item)
	case subSectionChange:
		section.Changes = append(section.Changes, item)
	case subSectionRemove:
		section.Removals = append(section.Removals, item)
	}
}

func (s *parserState) tryParseSectionHeader(line string) (bool, error) {
	if strings.Trim(line, " ") == "## [Unreleased]" {
		s.section = sectionUnreleased
		return true, nil
	}
	matches := regexp.MustCompile(`## \[(.*)\] - ([0-9]{4}-[0-9]{2}-[0-9]{2})`).FindStringSubmatch(line)
	if len(matches) == 3 {
		verStr := matches[1]
		date := matches[2]
		ver, err := semver.Parse("v" + verStr)
		if err != nil {
			return false, err
		}
		s.section = sectionHistory
		s.doc.History = append(s.doc.History, HistoricalSection{
			Version: ver,
			Date:    date,
			Section: Section{},
		})
		return true, nil
	}
	return false, nil
}

func (s *parserState) tryParseSubsectionHeader(line string) bool {
	sp := strings.Split(line, "### ")
	if len(sp) == 2 {
		r := true
		switch strings.Trim(sp[1], " ") {
		case "Added":
			s.subSection = subSectionAdd
		case "Changed":
			s.subSection = subSectionChange
		case "Fixed":
			s.subSection = subSectionFix
		case "Removed":
			s.subSection = subSectionRemove
		default:
			s.subSection = subSectionInvalid
			r = false
		}
		return r
	}
	return false
}

func (s *parserState) tryParseReferences(line string) bool {
	matched, _ := regexp.MatchString(`\[(.+)\]: (.+)`, line)
	return matched
}

func Parse(path string) (*Document, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	state := parserState{
		doc: &Document{},
	}
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		switch state.section {
		case sectionInvalid:
			if line == "# Changelog" {
				state.section = sectionHead
			}
		case sectionHead:
			if ok, err := state.tryParseSectionHeader(line); ok {
				continue
			} else if err != nil {
				return nil, parserErr(i, err)
			}
			if state.tryParseReferences(line) {
				continue
			}
			writeString(&state.doc.Description, line)
		case sectionUnreleased:
			if ok, err := state.tryParseSectionHeader(line); ok {
				continue
			} else if err != nil {
				return nil, parserErr(i, err)
			}
			if ok := state.tryParseSubsectionHeader(line); ok {
				continue
			}
			if ok := state.tryAddChange(line); ok {
				continue
			}
			if state.tryParseReferences(line) {
				continue
			}
			writeString(&state.doc.Unreleased.Description, line)
		case sectionHistory:
			if ok, err := state.tryParseSectionHeader(line); ok {
				continue
			} else if err != nil {
				return nil, parserErr(i, err)
			}
			if ok := state.tryParseSubsectionHeader(line); ok {
				continue
			}
			if ok := state.tryAddChange(line); ok {
				continue
			}
			if state.tryParseReferences(line) {
				continue
			}
			writeString(&state.doc.History[len(state.doc.History)-1].Section.Description, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return state.doc, nil
}

func parserErr(line int, err error) error {
	return fmt.Errorf("line %d: %v", line, err)
}

func writeString(dst *string, line string) {
	if line == "" {
		if len(*dst) > 0 {
			n := *dst + "\n\n"
			*dst = n
		}
	} else {
		if len(*dst) > 0 {
			n := *dst + " " + line
			*dst = n
		} else {
			*dst = line
		}
	}
}
