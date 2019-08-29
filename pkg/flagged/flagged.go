package flagged

import (
	"flag"
)

type Flag *bool

var (
	Small  = Flag(flag.Bool("small", true, "Run small tests"))
	Medium = Flag(flag.Bool("medium", false, "Run medium tests"))
	Large  = Flag(flag.Bool("large", false, "Run large tests"))

	Unit        = Flag(flag.Bool("unit", true, "Run unit tests"))
	Integration = Flag(flag.Bool("integration", false, "Run integration tests"))
	System      = Flag(flag.Bool("system", false, "Run system tests"))
	Acceptance  = Flag(flag.Bool("acceptance", false, "Run acceptance tests"))

	Service = Flag(flag.Bool("service", false, "Run service tests"))
	UI      = Flag(flag.Bool("ui", false, "Run UI tests"))

	Short = Flag(flag.Bool("short", false, "Run short tests"))
	long  = !(bool(*Short))
	Long  = Flag(&long)
)

func hasAny(t TestingT, flags ...Flag) bool {
	for _, flag := range flags {
		if *flag {
			return true
		}
	}
	return false
}

func hasAll(t TestingT, flags ...Flag) bool {
	for _, flag := range flags {
		if !*flag {
			return false
		}
	}
	return true
}

// Runs the test only if one of the provided flags is present.
func With(t TestingT, flags ...Flag) {
	if !hasAny(t, flags...) {
		// TODO: figure out how to generate a flag list
		t.Skip("None of the following flags were present: ")
	}
}

// Runs the test only if all the provided flags are present.
func WithAll(t TestingT, flags ...Flag) {
	if !hasAll(t, flags...) {
		// TODO: figure out how to generate a flag list
		t.Skip("One (or more) of the following flags was missing: ")
	}
}

// Runs the test only if none of the provided flags are present.
func Without(t TestingT, flags ...Flag) {
	if hasAny(t, flags...) {
		// TODO: figure out how to generate a flag list
		t.Skip("One (or more) of the following flags were present: ")
	}
}
