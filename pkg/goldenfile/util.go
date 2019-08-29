package goldenfile

import "flag"

var Update = flag.Bool("update", false, "Update golden files")

//
// Turn concrete testing.T structures into interfaces
// This idea stolen from Testify - https://github.com/stretchr/testify
//

type TestingT interface {
	Errorf(format string, args ...interface{})
	FailNow()
	Fatal(msgAndArgs ...interface{})
	Log(args ...interface{})
	Skip(msgAndArgs ...interface{})
}

type tHelper interface {
	Helper()
}
