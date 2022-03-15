package app

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/flowscan/repomaster-go/pkg/repo"
	"golang.org/x/mod/semver"
)

type App struct {
	ID          string `json:"-"`
	Name        string `json:"name"`
	Description string `json:"-"`
	Version     string `json:"version"`
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
	if _, err := os.Stat(a.Path() + "/app.json"); err != nil {
		return false
	}

	// check if semver is valid
	if !semver.IsValid(a.Version) {
		return false
	}

	return true
}

func (a *App) Semver() (int, int, int, string, string, error) {
	if !semver.IsValid(a.Version) {
		return 0, 0, 0, "", "", fmt.Errorf("invalid semver version: %s", a.Version)
	}
	return parseInt(semver.Major(a.Version)),
		parseInt(semver.Major(a.Version)),
		parseInt(semver.Major(a.Version)),
		semver.Prerelease(a.Version),
		semver.Build(a.Version),
		nil
}

func parseInt(s string) int {
	i, _ := strconv.ParseInt(s, 10, 32)
	return int(i)
}
