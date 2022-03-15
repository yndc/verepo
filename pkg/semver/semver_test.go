package semver_test

import (
	"fmt"
	"testing"

	"github.com/flowscan/repomaster-go/pkg/assert"
	"github.com/flowscan/repomaster-go/pkg/semver"
)

func TestParseVersion(t *testing.T) {
	asserter := assert.New(t)
	p, err := semver.Parse("v1.2.3")
	if err != nil {
		t.Error(err)
	}
	asserter.Equal(p.Major, 1)
	asserter.Equal(p.Minor, 2)
	asserter.Equal(p.Patch, 3)
	asserter.Zero(p.Prerelease)
	asserter.Zero(p.Build)
}

func TestParseVersionFail(t *testing.T) {
	_, err := semver.Parse("v1.2")
	if err == nil {
		t.Error("should error")
	}

	_, err = semver.Parse("1.2.3")
	if err == nil {
		t.Error("should error")
	}

	_, err = semver.Parse("v1.2.3.4")
	if err == nil {
		t.Error("should error")
	}

	_, err = semver.Parse("v1.2.3+")
	if err == nil {
		t.Error("should error")
	}
}

func TestParseVersionLong(t *testing.T) {
	asserter := assert.New(t)
	p, err := semver.Parse("v12.246.3235")
	if err != nil {
		t.Error(err)
	}
	asserter.Equal(p.Major, 12)
	asserter.Equal(p.Minor, 246)
	asserter.Equal(p.Patch, 3235)
	asserter.Zero(p.Prerelease)
	asserter.Zero(p.Build)
}

func TestParseVersionPrerelease(t *testing.T) {
	asserter := assert.New(t)
	p, err := semver.Parse("v1.2.3-dev")
	if err != nil {
		t.Error(err)
	}
	asserter.Equal(p.Major, 1)
	asserter.Equal(p.Minor, 2)
	asserter.Equal(p.Patch, 3)
	asserter.Equal(p.Prerelease.String(), "dev")
	asserter.Zero(p.Build)

	p, err = semver.Parse("v1.2.3-dev.asd")
	if err != nil {
		t.Error(err)
	}
	asserter.True(func() bool {
		return len(p.Prerelease) == 2
	})
	asserter.Equal(p.Prerelease.String(), "dev.asd")
	asserter.Zero(p.Build)

	p, err = semver.Parse("v1.2.3-dev.asd.cvb-3tfh")
	if err != nil {
		t.Error(err)
	}
	asserter.True(func() bool {
		return len(p.Prerelease) == 3
	})
	asserter.Equal(p.Prerelease.String(), "dev.asd.cvb-3tfh")
	asserter.Zero(p.Build)

	p, err = semver.Parse("v1.2.3-14.rc3")
	if err != nil {
		t.Error(err)
	}
	asserter.True(func() bool {
		return len(p.Prerelease) == 2
	})
	asserter.Equal(p.Prerelease.String(), "14.rc3")
	asserter.Zero(p.Build)
}

func TestParseVersionPrereleaseFail(t *testing.T) {
	asserter := assert.New(t)

	p, err := semver.Parse("v1.2.3--14.rc3")
	asserter.Err(err)
	p, err = semver.Parse("v1.2.3-@14.rc3")
	asserter.Err(err)
	p, err = semver.Parse("v1.2.3-ac-2..sa")
	asserter.Err(err)
	p, err = semver.Parse("v1.2.3-")
	asserter.Err(err)
	p, err = semver.Parse("v1.2.3-abv+")
	asserter.Err(err)

	fmt.Println(p)
}
