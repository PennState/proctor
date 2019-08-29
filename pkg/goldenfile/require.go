package goldenfile

//
// Requirements (in the same fashion as Testify)
//

func requireEqual(t TestingT, eq GoldenFileComparisonAssertionFunc, goldenFile string, actual []byte, msgAndArgs ...interface{}) {
	if assertEqual(t, eq, goldenFile, actual, msgAndArgs...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

func RequireBytesEq(t TestingT, goldenFile string, actual []byte, msgAndArgs ...interface{}) {
	requireEqual(t, bytesEq, goldenFile, actual, msgAndArgs...)
}

func RequireJSONEq(t TestingT, goldenFile string, actual string, msgAndArgs ...interface{}) {
	requireEqual(t, jsonEq, goldenFile, []byte(actual), msgAndArgs...)
}

func RequireStringEq(t TestingT, goldenFile string, actual string, msgAndArgs ...interface{}) {
	requireEqual(t, stringEq, goldenFile, []byte(actual), msgAndArgs...)
}

func RequireYAMLEq(t TestingT, goldenFile string, actual string, msgAndArgs ...interface{}) {
	requireEqual(t, yamlEq, goldenFile, []byte(actual), msgAndArgs...)
}
