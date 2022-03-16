package semver_test

import (
	"testing"

	"github.com/flowscan/repomaster-go/pkg/assert"
	"github.com/flowscan/repomaster-go/pkg/semver"
)

func TestComparePrerelease(t *testing.T) {
	asserter := assert.New(t)
	p1, err := semver.ParsePrerelease("dev")
	if err != nil {
		t.Error(err)
	}
	p2, err := semver.ParsePrerelease("lmao")
	if err != nil {
		t.Error(err)
	}
	asserter.Equal(semver.ComparePrerelease(p1, p2), 1)
}
