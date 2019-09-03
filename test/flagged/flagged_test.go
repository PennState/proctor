package blackboxflagged

import (
	"testing"

	"github.com/PennState/proctor/pkg/flagged"
	"github.com/stretchr/testify/assert"
)

type skipper func(flagged.TestingT, ...flagged.Flag)

func Test(t *testing.T) {
	set := true
	flagged.Medium.Flag = &set
	flagged.Integration.Flag = &set
	tests := []struct {
		Name    string
		Flags   []flagged.Flag
		Skip    bool
		Skipper skipper
	}{
		{Name: "FlaggedWithRuns", Flags: []flagged.Flag{flagged.Medium, flagged.System}, Skip: false, Skipper: flagged.With},
		{Name: "FlaggedWithSkips", Flags: []flagged.Flag{flagged.Large, flagged.System}, Skip: true, Skipper: flagged.With},
		{Name: "FlaggedWithAllRuns", Flags: []flagged.Flag{flagged.Medium, flagged.Integration}, Skip: false, Skipper: flagged.WithAll},
		{Name: "FlaggedWithAllSkips", Flags: []flagged.Flag{flagged.Medium, flagged.System}, Skip: true, Skipper: flagged.WithAll},
		{Name: "FlaggedWithOutRuns", Flags: []flagged.Flag{flagged.Large, flagged.System}, Skip: false, Skipper: flagged.Without},
		{Name: "FlaggedWithOutSkips", Flags: []flagged.Flag{flagged.Medium, flagged.System}, Skip: true, Skipper: flagged.Without},
	}
	for _, test := range tests {
		skipped := true
		t.Run(test.Name, func(t *testing.T) {
			test.Skipper(t, test.Flags...)
			skipped = false
		})
		assert.Equal(t, test.Skip, skipped)
	}
}
