package goldenfile

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fakeT struct {
	Failed bool
	Helped bool
}

func (f *fakeT) Errorf(fmt string, args ...interface{}) {}

func (f *fakeT) Fatal(args ...interface{}) {
	f.Failed = true
}

func (f *fakeT) Helper() {
	f.Helped = true
}

func (t *fakeT) Log(args ...interface{}) {}

func (f *fakeT) Run(name string, fn func(TestingT)) {
	fn(f)
}

func TestRequireEqualWithFileNotEqual(t *testing.T) {
	fp := GetDefaultFilePath("not-equal-file")
	err := ioutil.WriteFile(fp, []byte("This is **NOT** a test"), 0644)
	require.NoError(t, err)
	var fake fakeT
	f := &fake
	f.Run("", func(f TestingT) {
		RequireStringEq(f, fp, "This is a test")
	})
	assert.True(t, f.Failed)
	assert.True(t, f.Helped)
}

func TestRequireWithMissingFile(t *testing.T) {
	fp := GetDefaultFilePath("missing-file")
	_ = os.Remove(fp)
	var fake fakeT
	f := &fake
	f.Run("", func(f TestingT) {
		RequireStringEq(f, fp, "This is a test")
	})
	assert.True(t, f.Failed)
	assert.True(t, f.Helped)
}

func TestRequireWithUnwriteableFile(t *testing.T) {
	fp := GetDefaultFilePath("unwriteable-file")
	err := ioutil.WriteFile(fp, []byte("This file is not writeable"), 0444)
	if !os.IsPermission(err) {
		require.NoError(t, err)
	}
	_ = os.Chmod(fp, 0444)
	update := true
	Update = &update
	var fake fakeT
	f := &fake
	f.Run("", func(f TestingT) {
		RequireStringEq(f, fp, "This is a test")
	})
	assert.True(t, f.Failed)
	assert.True(t, f.Helped)
}
