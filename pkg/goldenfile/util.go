package goldenfile

import "flag"

var Update = flag.Bool("update", false, "Update golden files")

//
// Turn concrete testing.T structures into interfaces
// This idea stolen from Testify - https://github.com/stretchr/testify
//

type TestingT interface {
	Fatal(msgAndArgs ...interface{})
	Log(args ...interface{})
}

type tHelper interface {
	Helper()
}
