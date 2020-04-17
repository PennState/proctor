package goldenfile

import (
	"io/ioutil"
	"path/filepath"
)

const (
	DefaultTestDirectory    = "./testdata/"
	DefaultGoldenFileSuffix = ".golden"
)

/*
GetDefaultFilePath creates a file path from a file name by prepending
the default test directory name and appending the default golden file
suffix.  If the passed name is "foo.json", the resulting file path
will be "./testdata/foo.json.goldenfile".
*/
func GetDefaultFilePath(name string) string {
	return filepath.Join(DefaultTestDirectory, name+DefaultGoldenFileSuffix)
}

/*
PrependDefaultDirectoryPath creates a file path from a file name by
prepending the default test directory name.  If the provided name is
"foo.json", the resulting file path will be "./testdata/foo.json".
*/
func PrependDefaultDirectoryPath(name string) string {
	return filepath.Join(DefaultTestDirectory, name)
}

//
// Golden file I/O
//

func readGoldenFile(t TestingT, goldenFile string) []byte {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	data, err := ioutil.ReadFile(goldenFile)
	if err != nil {
		t.Fatal("Failed to read golden file: ", goldenFile)
	}
	return data
}

func writeGoldenFile(t TestingT, goldenFile string, actual []byte) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	err := ioutil.WriteFile(goldenFile, actual, 0644)
	if err != nil {
		t.Fatal("Failed to write golden file: ", goldenFile)
	}
	t.Log("Updated golden file: ", goldenFile)
}
