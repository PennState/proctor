package goldenfile

//
// Assertions and requirements (in the same fashion as Testify)
//

func assertEqual(t TestingT, eq GoldenFileComparisonAssertionFunc, goldenFile string, actual []byte, msgAndArgs ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if *Update {
		writeGoldenFile(t, goldenFile, actual)
		return true
	}
	expected := readGoldenFile(t, goldenFile)
	return eq(t, expected, actual, msgAndArgs...)
}

func AssertBytesEq(t TestingT, goldenFile string, actual []byte, msgAndArgs ...interface{}) bool {
	return assertEqual(t, bytesEq, goldenFile, actual, msgAndArgs...)
}

func AssertJSONEq(t TestingT, goldenFile string, actual string, msgAndArgs ...interface{}) bool {
	return assertEqual(t, jsonEq, goldenFile, []byte(actual), msgAndArgs...)
}

func AssertStringEq(t TestingT, goldenFile string, actual string, msgAndArgs ...interface{}) bool {
	return assertEqual(t, stringEq, goldenFile, []byte(actual), msgAndArgs...)
}

func AssertYAMLEq(t TestingT, goldenFile string, actual string, msgAndArgs ...interface{}) bool {
	return assertEqual(t, yamlEq, goldenFile, []byte(actual), msgAndArgs...)
}
