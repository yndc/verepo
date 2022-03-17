package changelog_test

import (
	"testing"

	"github.com/yndc/verepo/pkg/changelog"
)

func TestParser(t *testing.T) {
	c, _ := changelog.Parse("../../cmd/" + "verepo" + "/CHANGELOG.md")

	c.Write("./omg.md", "ayy")
}
