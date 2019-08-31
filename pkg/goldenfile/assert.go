package goldenfile

//
// Assertions and requirements (in the same fashion as Testify)
//

func assertEqual(t TestingT, eq GoldenFileComparisonAssertionFunc, filepath string, actual []byte, msgAndArgs ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if *Update {
		writeGoldenFile(t, filepath, actual)
		return true
	}
	expected := readGoldenFile(t, filepath)
	return eq(t, expected, actual, msgAndArgs...)
}

func AssertBytesEq(t TestingT, filepath string, actual []byte, msgAndArgs ...interface{}) bool {
	return assertEqual(t, bytesEq, filepath, actual, msgAndArgs...)
}

func AssertJSONEq(t TestingT, filepath string, actual string, msgAndArgs ...interface{}) bool {
	return assertEqual(t, jsonEq, filepath, []byte(actual), msgAndArgs...)
}

func AssertStringEq(t TestingT, filepath string, actual string, msgAndArgs ...interface{}) bool {
	return assertEqual(t, stringEq, filepath, []byte(actual), msgAndArgs...)
}

func AssertYAMLEq(t TestingT, filepath string, actual string, msgAndArgs ...interface{}) bool {
	return assertEqual(t, yamlEq, filepath, []byte(actual), msgAndArgs...)
}
