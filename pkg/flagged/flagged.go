package flagged

import (
	"flag"
	"strings"
)

type Flag struct {
	Name    string
	Default bool
	Help    string
	Flag    *bool
}

// NewFlag creates a new test flag which wraps flag.Bool in order to
// report why a test was skipped.
func NewFlag(name string, def bool, help string) Flag {
	return Flag{
		Name:    name,
		Default: def,
		Help:    help,
		Flag:    flag.Bool(name, def, help),
	}
}

var (
	Small  = NewFlag("small", true, "Run small tests")
	Medium = NewFlag("medium", false, "Run medium tests")
	Large  = NewFlag("large", false, "Run large tests")

	Unit        = NewFlag("unit", true, "Run unit tests")
	Integration = NewFlag("integration", false, "Run integration tests")
	System      = NewFlag("system", false, "Run system tests")
	Acceptance  = NewFlag("acceptance", false, "Run acceptance tests")

	Service = NewFlag("service", false, "Run service tests")
	UI      = NewFlag("ui", false, "Run UI tests")

	Short = NewFlag("short", false, "Run short tests")
	Long  = NewFlag("long", false, "Run long tests")
)

func hasAny(t TestingT, flags ...Flag) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	for _, flag := range flags {
		if *flag.Flag {
			return true
		}
	}
	return false
}

func hasAll(t TestingT, flags ...Flag) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	for _, flag := range flags {
		if !*flag.Flag {
			return false
		}
	}
	return true
}

func flagString(t TestingT, flags ...Flag) string {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	flagNames := []string{}
	for _, flag := range flags {
		flagNames = append(flagNames, flag.Name)
	}
	return strings.Join(flagNames, ", ")
}

// Runs the test only if one of the provided flags is present.
func With(t TestingT, flags ...Flag) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if !hasAny(t, flags...) {
		t.Skip("None of the following flags were present:", flagString(t, flags...))
	}
}

// Runs the test only if all the provided flags are present.
func WithAll(t TestingT, flags ...Flag) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if !hasAll(t, flags...) {
		t.Skip("One (or more) of the following flags was missing:", flagString(t, flags...))
	}
}

// Runs the test only if none of the provided flags are present.
func Without(t TestingT, flags ...Flag) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if hasAny(t, flags...) {
		t.Skip("One (or more) of the following flags were present:", flagString(t, flags...))
	}
}
