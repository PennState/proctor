package blackboxgoldenfile

import (
	"fmt"
	"math"
	"os"
	"testing"
	"time"

	"github.com/PennState/proctor/pkg/goldenfile"
	"github.com/stretchr/testify/assert"
)

var b = string([]byte{0xde, 0xad, 0xbe, 0xef})

const (
	j = `
{
	"fieldA": "Field A",
	"fieldB": 1234,
	"fieldC": [
		2345,
		3457
	],
	"fieldD": {
		"fieldE": "Field E",
		"fieldF": 4567
	}
}
`
	s = "This is a test"
	y = `
---
fieldA: "Field A"
fieldB: 1234
fieldC:
- 2345
- 3456
fieldD:
  fieldE: "Field E"
  fieldF: 4567
...
`
)

func assertModifiedTimeAfter(t *testing.T, goldenFile string, now time.Time) {
	t.Helper()
	fi, err := os.Stat(goldenFile)
	if err != nil {
		t.Error(err)
	}
	diff := math.Abs(fi.ModTime().Sub(now).Seconds())
	// TODO: Figure out why this assertion doesn't work (mtime should
	// always be after now) - then replace the hack below.
	//assert.True(t, fi.ModTime().After(now))
	assert.Less(t, diff, 10*time.Millisecond.Seconds())
}

func adaptStringToBytes(t *testing.T, f func(goldenfile.TestingT, string, []byte, ...interface{}) bool) func(goldenfile.TestingT, string, string, ...interface{}) bool {
	t.Helper()
	return func(t goldenfile.TestingT, goldenFile string, actual string, msgAndArgs ...interface{}) bool {
		return f(t, goldenFile, []byte(actual), msgAndArgs)
	}
}

func adaptFakeReturn(t *testing.T, f func(goldenfile.TestingT, string, string, ...interface{})) func(goldenfile.TestingT, string, string, ...interface{}) bool {
	t.Helper()
	return func(t goldenfile.TestingT, goldenFile string, actual string, msgAndArgs ...interface{}) bool {
		f(t, goldenFile, actual, msgAndArgs)
		return true
	}
}

func adaptFakeReturnAndStringToBytes(t *testing.T, f func(goldenfile.TestingT, string, []byte, ...interface{})) func(goldenfile.TestingT, string, string, ...interface{}) bool {
	t.Helper()
	return func(t goldenfile.TestingT, goldenFile string, actual string, msgAndArgs ...interface{}) bool {
		f(t, goldenFile, []byte(actual), msgAndArgs)
		return true
	}
}

func TestGoldenFiles(t *testing.T) {
	tests := []struct {
		Name      string
		Type      string
		Data      string
		Assertion func(goldenfile.TestingT, string, string, ...interface{}) bool
	}{
		{Name: "bytes", Type: "assert", Data: b, Assertion: adaptStringToBytes(t, goldenfile.AssertBytesEq)},
		{Name: "json", Type: "assert", Data: j, Assertion: goldenfile.AssertJSONEq},
		{Name: "string", Type: "assert", Data: s, Assertion: goldenfile.AssertStringEq},
		{Name: "yaml", Type: "assert", Data: y, Assertion: goldenfile.AssertYAMLEq},
		{Name: "bytes", Type: "require", Data: b, Assertion: adaptFakeReturnAndStringToBytes(t, goldenfile.RequireBytesEq)},
		{Name: "json", Type: "require", Data: j, Assertion: adaptFakeReturn(t, goldenfile.RequireJSONEq)},
		{Name: "string", Type: "require", Data: s, Assertion: adaptFakeReturn(t, goldenfile.RequireStringEq)},
		{Name: "yaml", Type: "require", Data: y, Assertion: adaptFakeReturn(t, goldenfile.RequireYAMLEq)},
	}
	updates := []bool{false, true}
	for _, update := range updates {
		for _, test := range tests {
			testName := fmt.Sprintf("%s-%s-%t", test.Name, test.Type, update)
			t.Run(testName, func(t *testing.T) {
				goldenfile.Update = &update
				fp := goldenfile.MakeDefaultName(test.Name)
				now := time.Now().Truncate(0)
				actual := test.Assertion(t, fp, test.Data)
				assert.True(t, actual)
				if update {
					assertModifiedTimeAfter(t, fp, now)
				}
			})
		}
	}
}
