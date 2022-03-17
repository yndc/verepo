package git

import (
	"regexp"
	"strings"

	"github.com/yndc/verepo/pkg/exec"
)

func GetOrigin() string {
	o, _ := exec.Exec("git", "remote", "get-url", "origin")
	s := strings.Trim(string(o), " \n")

	s = regexp.MustCompile(`oauth(.+)@`).ReplaceAllString(s, "")
	s = regexp.MustCompile(`\.git$`).ReplaceAllString(s, "")
	return s
}
