package repo

import "strings"

func GetBasePath() string {
	return "/home/yondercode/projects/go-repomaster"
}

func Path(to string) string {
	if !strings.HasPrefix(to, "/") {
		to = "/" + to
	}
	return GetBasePath() + to
}
