package goldenfile

//
// Requirements (in the same fashion as Testify)
//

func requireEqual(t TestingT, eq GoldenFileComparisonAssertionFunc, filepath string, actual []byte, msgAndArgs ...interface{}) {
	if assertEqual(t, eq, filepath, actual, msgAndArgs...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.Fatal(msgAndArgs)
}

func RequireBytesEq(t TestingT, filepath string, actual []byte, msgAndArgs ...interface{}) {
	requireEqual(t, bytesEq, filepath, actual, msgAndArgs...)
}

func RequireJSONEq(t TestingT, filepath string, actual string, msgAndArgs ...interface{}) {
	requireEqual(t, jsonEq, filepath, []byte(actual), msgAndArgs...)
}

func RequireStringEq(t TestingT, filepath string, actual string, msgAndArgs ...interface{}) {
	requireEqual(t, stringEq, filepath, []byte(actual), msgAndArgs...)
}

func RequireYAMLEq(t TestingT, filepath string, actual string, msgAndArgs ...interface{}) {
	requireEqual(t, yamlEq, filepath, []byte(actual), msgAndArgs...)
}
