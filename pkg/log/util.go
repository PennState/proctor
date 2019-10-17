package log

import (
	"flag"
)

var level = flag.String("level", "WARN", "Sets the logging level")

//
// Turn concrete testing.T structures into interfaces
// This idea stolen from Testify - https://github.com/stretchr/testify
//

type TestingT interface {
	Log(args ...interface{})
}

type tHelper interface {
	Helper()
}
