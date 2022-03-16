package changelog_test

import (
	"fmt"
	"testing"

	"github.com/flowscan/repomaster-go/pkg/changelog"
)

func TestParser(t *testing.T) {
	c, _ := changelog.Parse("../../cmd/" + "repomaster" + "/CHANGELOG.md")
	fmt.Println(c)
}
