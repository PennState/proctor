package flagged

import (
	"flag"
	"testing"
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
)

func WithAll(t *testing.T, flags ...Flag) {
	for _, flag := range flags {
		if !*flag {
			t.SkipNow()
		}
	}
}

func WithAny(t *testing.T, flags ...Flag) {
	skip := true
	for _, flag := range flags {
		if *flag {
			skip = false
			break
		}
	}

	if skip {
		t.Skip("None of the following flags were present: ")
	}
}
