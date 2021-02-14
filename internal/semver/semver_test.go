package semver_test

import (
	"semver/internal/semver"
	"strings"
	"testing"
)

func TestBumpPatch(t *testing.T) {
	ver := "1.0.0"
	expected := "1.0.1"
	bumped, err := semver.BumpPatch(ver)
	if err != nil {
		panic(err)
	}
	if bumped != expected {
		t.Errorf("PatchBumped version of '%s' was '%s', should've been '%s'", ver, bumped, expected)
	}
	strings.Compare("asd", "das")
}

func TestBumpPatchWithPrerelAndMetadata(t *testing.T) {
	ver := "1.0.1-a20210131+myawesomebranch-a01f7f7a"
	expected := "1.0.2-a20210131+myawesomebranch-a01f7f7a"
	bumped, err := semver.BumpPatch(ver)
	if err != nil {
		panic(err)
	}
	if bumped != expected {
		t.Errorf("PatchBumped version of '%s' was '%s', should've been '%s'", ver, bumped, expected)
	}
}

func TestBumpPrerelWithPrerelAndMetadata(t *testing.T) {
	ver := "1.0.1-a20210131.1+myawesomebranch-a01f7f7a"
	expected := "1.0.1-a20210131.2+myawesomebranch-a01f7f7a"
	bumped, err := semver.BumpPrerel(ver)
	if err != nil {
		panic(err)
	}
	if bumped != expected {
		t.Errorf("PatchBumped version of '%s' was '%s', should've been '%s'", ver, bumped, expected)
	}
}

func TestBumpPrerelWithDottedPrerelAndMetadata(t *testing.T) {
	ver := "1.0.1-a2021013.1.1+myawesomebranch-a01f7f7a"
	expected := "1.0.1-a2021013.1.2+myawesomebranch-a01f7f7a"
	bumped, err := semver.BumpPrerel(ver)
	if err != nil {
		panic(err)
	}
	if bumped != expected {
		t.Errorf("PatchBumped version of '%s' was '%s', should've been '%s'", ver, bumped, expected)
	}
}
func TestBumpPrerelWithUnversionedPrerelAndMetadata(t *testing.T) {
	ver := "1.0.1-a20210131+myawesomebranch-a01f7f7a"
	expected := "1.0.1-a20210131.0+myawesomebranch-a01f7f7a"
	bumped, err := semver.BumpPrerel(ver)
	if err != nil {
		panic(err)
	}
	if bumped != expected {
		t.Errorf("PatchBumped version of '%s' was '%s', should've been '%s'", ver, bumped, expected)
	}
}
func TestBumpPrerelWithoutPrerel(t *testing.T) {
	ver := "1.0.1"
	expected := "1.0.2-0"
	bumped, err := semver.BumpPrerel(ver)
	if err != nil {
		panic(err)
	}
	if bumped != expected {
		t.Errorf("PatchBumped version of '%s' was '%s', should've been '%s'", ver, bumped, expected)
	}
}

func TestBumpMinorWithPrerelAndMetadata(t *testing.T) {
	ver := "1.0.1-a20210131+myawesomebranch-a01f7f7a"
	expected := "1.1.0-a20210131+myawesomebranch-a01f7f7a"
	bumped, err := semver.BumpMinor(ver)
	if err != nil {
		panic(err)
	}
	if bumped != expected {
		t.Errorf("PatchBumped version of '%s' was '%s', should've been '%s'", ver, bumped, expected)
	}
}

func TestBumpMajorWithPrerelAndMetadata(t *testing.T) {
	ver := "1.0.1-a20210131+myawesomebranch-a01f7f7a"
	expected := "2.0.0-a20210131+myawesomebranch-a01f7f7a"
	bumped, err := semver.BumpMajor(ver)
	if err != nil {
		panic(err)
	}
	if bumped != expected {
		t.Errorf("PatchBumped version of '%s' was '%s', should've been '%s'", ver, bumped, expected)
	}
}

func TestParseSemver(t *testing.T) {
	ver := "1.0.1-a20210131+myawesomebranch-a01f7f7a"
	expected := semver.Semver{Major: 1, Minor: 0, Patch: 1, Prerelease: "a20210131", Metadata: "myawesomebranch-a01f7f7a"}
	actual, err := semver.ParseSemver(ver)
	if err != nil {
		panic(err)
	}
	if expected != actual {
		t.Errorf("ParseSemver of '%s' produced '%+v', should've been '%+v'", ver, actual, expected)
	}
}
func TestParseSemverNoPatch(t *testing.T) {
	ver := "1.0-a20210131+myawesomebranch-a01f7f7a"
	_, err := semver.ParseSemver(ver)
	if err == nil {
		t.Errorf("Version '%s' was accepted, though it is invalid", ver)

	}
}

func TestParseSemverEmpty(t *testing.T) {
	ver := ""
	_, err := semver.ParseSemver(ver)
	if err == nil {
		t.Errorf("Version '%s' was accepted, though it is invalid", ver)

	}
}

func TestParseSemverMajorOnly(t *testing.T) {
	ver := "1"
	_, err := semver.ParseSemver(ver)
	if err == nil {
		t.Errorf("Version '%s' was accepted, though it is invalid", ver)

	}
	_, err = semver.BumpPatch(ver)
	if err == nil {
		t.Errorf("Version '%s' was accepted, though it is invalid", ver)

	}
	_, err = semver.BumpMinor(ver)
	if err == nil {
		t.Errorf("Version '%s' was accepted, though it is invalid", ver)

	}
	_, err = semver.BumpMajor(ver)
	if err == nil {
		t.Errorf("Version '%s' was accepted, though it is invalid", ver)

	}
	_, err = semver.BumpPrerel(ver)
	if err == nil {
		t.Errorf("Version '%s' was accepted, though it is invalid", ver)

	}

}
