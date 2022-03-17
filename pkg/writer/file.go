package writer

import (
	"fmt"
	"os"
	"strings"

	"github.com/yndc/verepo/pkg/config"
	"github.com/yndc/verepo/pkg/repo"
)

func File(relativePath string, contents []byte) error {
	if !strings.HasPrefix(relativePath, "/") {
		relativePath = "/" + relativePath
	}
	dstPath := repo.GetBasePath() + relativePath
	if config.Global.DryRun {
		fmt.Printf("write to %s: %s\n", dstPath, string(contents))
		return nil
	} else {
		return os.WriteFile(dstPath, contents, 0644)
	}
}
