package flagged

//
// Turn concrete testing.T structures into interfaces
// This idea stolen from Testify - https://github.com/stretchr/testify
//

type TestingT interface {
	Log(args ...interface{})
	Skip(msgAndArgs ...interface{})
}

type tHelper interface {
	Helper()
}
