package app

import (
	"encoding/json"
	"os"

	"github.com/flowscan/repomaster-go/pkg/repo"
	"github.com/flowscan/repomaster-go/pkg/semver"
)

type App struct {
	ID      string        `json:"-"`
	Version semver.Parsed `json:"version"`
}

func ParseJson(app string, b []byte) (*App, error) {
	var a *App
	err := json.Unmarshal(b, &a)
	if err != nil {
		return nil, err
	}
	a.ID = app
	return a, nil
}

func (a *App) Path() string {
	return repo.Path("/cmd/" + a.ID)
}

func (a *App) Valid() bool {
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
