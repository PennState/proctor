package flagged_test

import (
	"flag"
	"os"
	"testing"

	"github.com/PennState/proctor/pkg/flagged"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	flag.Set("medium", "true")
	flag.Set("integration", "true")
	flag.Parse()
	os.Exit(m.Run())
}

func TestFlaggedWithAllRuns(t *testing.T) {
	ran := false
	t.Run("", func(t *testing.T) {
		flagged.WithAll(t, flagged.Medium, flagged.Integration)
		ran = true
	})
	assert.True(t, ran)
}

func TestFlaggedWithAllSkips(t *testing.T) {
	skipped := true
	t.Run("", func(t *testing.T) {
		flagged.WithAll(t, flagged.Medium, flagged.System)
		skipped = false
	})
	assert.True(t, skipped)
}

func TestFlaggedWithAnyRuns(t *testing.T) {
	ran := false
	t.Run("", func(t *testing.T) {
		flagged.WithAny(t, flagged.Medium, flagged.System)
		ran = true
	})
	assert.True(t, ran)
}

func TestFlaggedWithAnySkips(t *testing.T) {
	skipped := true
	t.Run("", func(t *testing.T) {
		flagged.WithAny(t, flagged.Large, flagged.System)
		skipped = false
	})
	assert.True(t, skipped)
}
