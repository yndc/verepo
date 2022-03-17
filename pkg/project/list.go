package project

import (
	"os"

	"github.com/flowscan/repomaster-go/pkg/git"
	"github.com/flowscan/repomaster-go/pkg/repo"
)

// Get all applications in the repository
func GetAll() ([]Project, error) {
	folders, err := os.ReadDir(repo.GetBasePath() + "/cmd")
	if err != nil {
		return nil, err
	}

	projects := make([]Project, 0)

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

		project := Project{
			ID:      v.Name(),
			Version: git.Latest(v.Name()),
		}
		if err != nil {
			return nil, err
		}

		projects = append(projects, project)
	}

	return projects, nil
}
