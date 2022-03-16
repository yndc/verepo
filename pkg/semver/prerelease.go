package semver

import (
	"fmt"
	"strconv"
	"strings"
)

type Prerelease []string

func (p Prerelease) String() string {
	return strings.Join(p, ".")
}

// bumps a pre-release tag, if the last identifier is numeric, then the number will be incremented by 1
// otherwise a new numeric identifier "1" will be appended at the end of the pre-release tag
func (p Prerelease) Bump() Prerelease {
	n := p
	if i, err := parseInt(n[len(n)-1]); err == nil {
		n[len(n)-1] = strconv.FormatInt(int64(i+1), 10)
	} else {
		n = append(n, "1")
	}
	return n
}

func ParsePrerelease(p string) (Prerelease, error) {
	splitted := strings.Split(p, ".")
	for _, identifier := range splitted {
		if err := validateIdentifier(identifier); err != nil {
			return nil, err
		}
	}
	return splitted, nil
}

func ComparePrerelease(left Prerelease, right Prerelease) int {
	if left.String() == right.String() {
		return 0
	}

	// compare each prerelease identifiers, stops once a precedence difference is found
	limit := min(len(left), len(right))
	for i := 0; i < limit; i++ {
		cmp := comparePrereleaseIdentifier(left[i], right[i])
		if cmp != 0 {
			return cmp
		}
	}

	// the longer identifiers have higher precedence
	if len(left) > len(right) {
		return +1
	}
	return -1
}

func validateIdentifier(identifier string) error {
	// no empty identifiers
	if len(identifier) == 0 {
		return fmt.Errorf("pre-release identifiers cannot be empty")
	}

	// must not start with --
	if identifier[0] == '-' {
		return fmt.Errorf("must not start with -")
	}

	// must only ascii characters
	for _, c := range identifier {
		if c < 48 || c > 122 ||
			(c > 57 && c < 65) ||
			(c > 90 && c < 97) {
			if c != '-' {
				return fmt.Errorf("pre-release identifiers must comprise only ASCII alphanumerics and hyphens")
			}
		}
	}

	// no leading zeroes on numeric identifiers
	if isNumeric(identifier) {
		if identifier[0] == '0' {
			return fmt.Errorf("numeric identifiers MUST NOT include leading zeroes")
		}
	}

	return nil
}

func comparePrereleaseIdentifier(left string, right string) int {
	if leftInt, err := parseInt(left); err == nil {
		if rightInt, err := parseInt(right); err == nil {
			// if both are numeric, compare numerically
			if leftInt > rightInt {
				return +1
			} else if leftInt < rightInt {
				return -1
			} else {
				return 0
			}
		}
		// if left is numeric but right is not, string always wins
		return -1
	} else if isNumeric(right) {
		// if left is string but right is numeric, string always wins
		return +1
	}

	// compare lexigraphically
	return strings.Compare(left, right)
}

func isNumeric(s string) bool {
	_, err := parseInt(s)
	return err == nil
}
