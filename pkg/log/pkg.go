/*
Package log wraps the testing package's T.Log(args ...interface{}) method
with a logrus logger ("github.com/sirupsen/logrus").  Since this package
merely provides a logrus Logger, all logging levels and their associated
methods can be used (see the full logrus documentation at
https://github.com/sirupsen/logrus).  Calling the setOutput()or setLevel()
methods may however defeat this wrapper's functionality. A simple example
that creates and uses a log is shown below.

  package main_test

  import ("github.com/PennState/proctor/pkg/log")

  func Test(t *testing.T) {
	  log := log.New(t)
	  log.Trace("-> Test(t *testing.T)")
	  // add some testing code
	  log.Infof("Important key: %v, value: %v", key, val)
	  log.Trace("Test(t *testing.T) ->")
  }

If these tests are executed with:

  go test ./...

then no output is displayed as the default logging level is set to WARN.
If the test is executed using:

  go test ./... -level=INFO

then only as single line is output as shown below:

Output:
  entry.go:259: time="2019-10-18T08:40:59-04:00" level=info msg="Warn-level logging"

If the test is executed using:

  go test ./... -level=TRACE

then three lines are output as expected:

Output:
  entry.go:259: time="2019-10-18T08:44:57-04:00" level=trace msg="-> Test(t *testing.T)"
  entry.go:259: time="2019-10-18T08:44:57-04:00" level=info msg="Important key: MyKey, value: MyValue"
  entry.go:259: time="2019-10-18T08:44:57-04:00" level=trace msg="Test(t *testing.T) ->"

Note that the go test command suppresses all output when the tests succeed,
so the above examples also assume that there is a test failure.  Log
messages (along with the rest of the test output) can be seen by using the
-v (verbose) flag.
*/
package log
