package goldenfile

import (
	"io/ioutil"
	"path/filepath"
)

const (
	DefaultTestDirectory    = "./testdata/"
	DefaultGoldenFileSuffix = ".golden"
)

func GetDefaultFilePath(name string) string {
	return filepath.Join(DefaultTestDirectory, name+DefaultGoldenFileSuffix)
}

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
