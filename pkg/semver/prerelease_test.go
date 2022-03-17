package semver_test

import (
	"testing"

	"github.com/yndc/verepo/pkg/assert"
	"github.com/yndc/verepo/pkg/semver"
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
