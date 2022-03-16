package app

import (
	"os"

	"github.com/flowscan/repomaster-go/pkg/git"
	"github.com/flowscan/repomaster-go/pkg/repo"
)

// Get all applications in the repository
func GetAll() ([]App, error) {
	folders, err := os.ReadDir(repo.GetBasePath() + "/cmd")
	if err != nil {
		return nil, err
	}

	apps := make([]App, 0)

	for _, v := range folders {
		files, err := os.ReadDir(repo.GetBasePath() + "/cmd/" + v.Name())
		if err != nil {
			return nil, err
		}

		hasMain := false
		for _, f := range files {
			switch f.Name() {
			case "main.go":
				hasMain = true
			}
		}
		if !hasMain {
			continue
		}

		app := App{
			ID:      v.Name(),
			Version: git.Latest(v.Name()),
		}
		if err != nil {
			return nil, err
		}

		apps = append(apps, app)
	}

	return apps, nil
}
