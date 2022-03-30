package project

import (
	"os"
	"strings"

	"github.com/yndc/verepo/pkg/git"
)

// Get all applications in the repository
func GetAll() ([]Project, error) {
	return traverse("./cmd", "")
}

func traverse(root string, step string) ([]Project, error) {
	nodes, err := os.ReadDir(joinPath(root, step))
	if err != nil {
		return nil, err
	}

	projects := make([]Project, 0)

	for _, v := range nodes {
		if v.Name() == "main.go" {
			projects = append(projects, Project{
				ID:      step,
				Version: git.Latest(step),
			})
		} else {
			if v.IsDir() {
				r, err := traverse(root, joinPath(step, v.Name()))
				if err != nil {
					return nil, err
				}
				projects = append(projects, r...)
			}
		}
	}

	return projects, nil
}

func joinPath(steps ...string) string {
	switch len(steps) {
	case 0:
		return ""
	case 1:
		return steps[0]
	default:
		bldr := strings.Builder{}
		for _, step := range steps {
			if len(step) > 0 {
				if bldr.Len() > 0 {
					bldr.WriteByte('/')
				}
				bldr.WriteString(step)
			}
		}
		return bldr.String()
	}
}
