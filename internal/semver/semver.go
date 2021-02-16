package semver

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Semver is a struct representing a semeantic version
type Semver struct {
	Major      int
	Minor      int
	Patch      int
	Prerelease string
	Metadata   string
}

func (s *Semver) String() string {
	ss := ""
	ss = ss + fmt.Sprintf("%d.%d.%d", s.Major, s.Minor, s.Patch)
	if s.Prerelease != "" {
		ss = ss + fmt.Sprintf("-%s", s.Prerelease)
	}
	if s.Metadata != "" {
		ss = ss + fmt.Sprintf("+%s", s.Metadata)
	}
	return ss
}

//ValidSemver checks if a semver string is valid
func ValidSemver(ver string) (match bool) {
	r := `^([0-9]+)\.([0-9]+)\.([0-9]+)(?:-([0-9A-Za-z-]+(?:\.[0-9A-Za-z-]+)*))?(?:\+[0-9A-Za-z-]+)?$`
	match, _ = regexp.Match(r, []byte(ver))
	return
}

// ParseSemver takes a string representation of a semantic version,
// and generates a Semver struct.
func ParseSemver(ver string) (Semver, error) {
	if !ValidSemver(ver) {
		return Semver{}, fmt.Errorf("Invalid semver: %s", ver)
	}
	s := Semver{}
	temp := strings.SplitN(ver, "+", 2)
	ver = temp[0]
	if len(temp) == 2 {
		s.Metadata = temp[1]
	}
	temp = strings.SplitN(ver, "-", 2)
	ver = temp[0]
	if len(temp) == 2 {
		s.Prerelease = temp[1]
	}

	temp = strings.Split(ver, ".")
	s.Major = 0
	s.Minor = 0
	s.Patch = 0
	if m, err := strconv.Atoi(temp[0]); err == nil {
		s.Major = m
	}
	if m, err := strconv.Atoi(temp[1]); err == nil {
		s.Minor = m
	}

	if m, err := strconv.Atoi(temp[2]); err == nil {
		s.Patch = m
	}
	return s, nil
}

// BumpPatch takes a semver string, bumps the Patch version N.N.X up by one
// clearing the prerel and metadata portions.
func BumpPatch(ver string) (string, error) {
	s, err := ParseSemver(ver)
	if err != nil {
		return "", err
	}
	s.Patch = s.Patch + 1
	return s.String(), nil
}

// BumpMinor takes a semver string, bumps the Minor version N.X.0 up by one
// and zeroes the patch version, and clearing the prerel and metadata.
func BumpMinor(ver string) (string, error) {
	s, err := ParseSemver(ver)
	if err != nil {
		return "", err
	}
	s.Minor = s.Minor + 1
	s.Patch = 0
	s.Prerelease = ""
	s.Metadata = ""
	return s.String(), nil
}

// BumpMajor takes a semver string, bumps the Major version X.0.0 up by one
// zeroing the minor and patch versions, as well as clearing the prerel and metadata
func BumpMajor(ver string) (string, error) {
	s, err := ParseSemver(ver)
	if err != nil {
		return "", err
	}
	s.Major = s.Major + 1
	s.Minor = 0
	s.Patch = 0
	s.Prerelease = ""
	s.Metadata = ""
	return s.String(), nil
}

// BumpPrerel takes a semver string, and bumps the prerel version.
// If the prerel is unversioned (not numeric or ending in a .N)
// the prerel version is set to 0.
// 1.0.0.a2021 -> 1.0.0.a2021.0
// 1.0.0.a2021.1 -> 1.0.0.a2021.2
// 1.0.0 -> 1.0.0-0
// 1.0.0-1 -> 1.0.0-2
func BumpPrerel(ver string) (string, error) {
	s, err := ParseSemver(ver)
	if err != nil {
		return "", err
	}
	s.Metadata = ""
	if s.Prerelease == "" {
		s.Patch = s.Patch + 1
		s.Prerelease = "0"
		return s.String(), nil
	}
	parts := strings.SplitAfter(s.Prerelease, ".")
	last := parts[len(parts)-1]
	if prerelnum, err := strconv.Atoi(last); err == nil {
		prerelnum = prerelnum + 1
		parts[len(parts)-1] = fmt.Sprintf("%d", prerelnum)

		s.Prerelease = strings.Join(parts, "")
		return s.String(), nil
	}
	s.Prerelease = s.Prerelease + ".0"
	return s.String(), nil
}
