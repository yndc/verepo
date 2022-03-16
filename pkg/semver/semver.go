package semver

import (
	"fmt"
	"strconv"
	"strings"
)

type Parsed struct {
	Major      int
	Minor      int
	Patch      int
	Prerelease Prerelease
	Build      string
	Invalid    bool
}

func Invalid() Parsed {
	return Parsed{
		Invalid: true,
	}
}

func Parse(v string) (p Parsed, err error) {
	if v == "" || v[0] != 'v' {
		return Invalid(), fmt.Errorf("missing v prefix")
	}
	var major, minor, patch int
	var build string
	var prerelease Prerelease
	v = v[1:]

	// extract build
	splitBuild := strings.Split(v, "+")
	switch len(splitBuild) {
	case 2:
		if len(splitBuild[1]) == 0 {
			return Invalid(), fmt.Errorf("build metadata is empty despite using + character")
		}
		build = splitBuild[1]
	case 1:
	default:
		return Invalid(), fmt.Errorf("invalid build version")
	}

	// extract prerelease
	splitPrerelease := strings.Split(splitBuild[0], "-")
	if len(splitPrerelease) > 1 {
		prerelease, err = ParsePrerelease(strings.Join(splitPrerelease[1:], "-"))
		if err != nil {
			return Invalid(), fmt.Errorf("invalid pre-release: %v", err)
		}
	}

	// extract the version
	splitVersion := strings.Split(splitPrerelease[0], ".")
	if len(splitVersion) != 3 {
		return Invalid(), fmt.Errorf("invalid version: %s", splitPrerelease[0])
	}
	major, err = parseInt(splitVersion[0])
	if err != nil {
		return Invalid(), fmt.Errorf("invalid major version: %s", splitVersion[0])
	}
	minor, err = parseInt(splitVersion[1])
	if err != nil {
		return Invalid(), fmt.Errorf("invalid minor version: %s", splitVersion[1])
	}
	patch, err = parseInt(splitVersion[2])
	if err != nil {
		return Invalid(), fmt.Errorf("invalid patch version: %s", splitVersion[2])
	}

	return Parsed{
		Major:      major,
		Minor:      minor,
		Patch:      patch,
		Prerelease: prerelease,
		Build:      build,
		Invalid:    false,
	}, nil
}

func (p Parsed) String() string {
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("v%d.%d.%d", p.Major, p.Minor, p.Patch))
	if len(p.Prerelease) > 0 {
		b.WriteRune('-')
		b.WriteString(p.Prerelease.String())
	}
	if len(p.Build) > 0 {
		b.WriteRune('+')
		b.WriteString(p.Build)
	}
	return b.String()
}

func (p Parsed) MarshalJSON() ([]byte, error) {
	if p.Invalid {
		return nil, fmt.Errorf("semver is invalid")
	}
	return []byte("\"" + p.String() + "\""), nil
}

func (p Parsed) UnmarshalJSON(s []byte) error {
	p, err := Parse(strings.Trim(string(s), "\""))
	return err
}

func (p Parsed) BumpPatch() Parsed {
	return Parsed{
		Major:      p.Major,
		Minor:      p.Minor,
		Patch:      p.Patch + 1,
		Prerelease: p.Prerelease,
		Build:      p.Build,
	}
}

func (p Parsed) BumpMinor() Parsed {
	return Parsed{
		Major:      p.Major,
		Minor:      p.Minor + 1,
		Patch:      0,
		Prerelease: p.Prerelease,
		Build:      p.Build,
	}
}

func (p Parsed) BumpMajor() Parsed {
	return Parsed{
		Major:      p.Major + 1,
		Minor:      0,
		Patch:      0,
		Prerelease: p.Prerelease,
		Build:      p.Build,
	}
}

func Compare(left Parsed, right Parsed) int {
	verCmp := CompareVersion(left, right)
	if verCmp == 0 {
		if len(left.Prerelease) > 0 {
			if len(right.Prerelease) > 0 {
				return ComparePrerelease(left.Prerelease, right.Prerelease)
			} else {
				return -1
			}
		} else if len(right.Prerelease) > 0 {
			return +1
		}
		return 0
	}
	return verCmp
}

func CompareVersion(left Parsed, right Parsed) int {
	// compare major
	if left.Major > right.Major {
		return +1
	} else if left.Major < right.Major {
		return -1
	} else {
		// compare minor
		if left.Minor > right.Minor {
			return +1
		} else if left.Minor < right.Minor {
			return -1
		} else {
			// compare patch
			if left.Patch > right.Patch {
				return +1
			} else if left.Patch < right.Patch {
				return -1
			} else {
				return 0
			}
		}
	}
}

func parseInt(v string) (int, error) {
	i, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return 0, err
	}
	return int(i), nil
}

func min(left int, right int) int {
	if left > right {
		return right
	}
	return left
}
