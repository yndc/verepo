package app

import (
	"encoding/json"
	"os"

	"github.com/flowscan/repomaster-go/pkg/repo"
	"github.com/flowscan/repomaster-go/pkg/semver"
	"github.com/flowscan/repomaster-go/pkg/writer"
)

// Get all applications in the repository
func GetAll() ([]*App, error) {
	folders, err := os.ReadDir(repo.GetBasePath() + "/cmd")
	if err != nil {
		return nil, err
	}

	apps := make([]*App, 0)

	for _, v := range folders {
		files, err := os.ReadDir(repo.GetBasePath() + "/cmd/" + v.Name())
		if err != nil {
			return nil, err
		}

		hasMain := false
		hasAppJson := false
		for _, f := range files {
			switch f.Name() {
			case "main.go":
				hasMain = true
			case "app.json":
				hasAppJson = true
			}
		}
		if !hasMain {
			continue
		}

		if !hasAppJson {
			newAppJson := generateAppJson(v.Name())
			jsonBytes, err := json.MarshalIndent(newAppJson, "", "\t")
			if err != nil {
				return nil, err
			}
			err = writer.File("/cmd/"+v.Name()+"/app.json", jsonBytes)
			if err != nil {
				return nil, err
			}
		}

		application, err := Get(v.Name())
		if err != nil {
			return nil, err
		}

		apps = append(apps, application)
	}

	return apps, nil
}

func Get(app string) (*App, error) {
	src, err := os.ReadFile(repo.Path("/cmd/" + app + "/app.json"))
	if err != nil {
		return nil, err
	}
	return ParseJson(app, src)
}

func generateAppJson(app string) *App {
	v, _ := semver.Parse("v0.0.0-dev")
	return &App{
		ID:          app,
		Name:        app,
		Description: "",
		Version:     v,
	}
}
