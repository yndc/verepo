package project

import (
	"os"

	"github.com/yndc/verepo/pkg/repo"
	"github.com/yndc/verepo/pkg/semver"
)

type Project struct {
	ID      string        `json:"-"`
	Version semver.Parsed `json:"version"`
}

func (a *Project) Path() string {
	return repo.Path("/cmd/" + a.ID)
}

func (a *Project) Valid() bool {
	// check if path exists
	if _, err := os.Stat(a.Path() + "/main.go"); err != nil {
		return false
	}

	// check if semver is valid
	if a.Version.Invalid {
		return false
	}

	return true
}
