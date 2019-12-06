/*
Package proctor provides utilities to conditionally control the execution
of Go tests.

The Go language provides an awesome built-in facility for writing tests,
which, when combined with Testify's assertions, allow a developer to
test pretty much every facet of their code.  The proctor library provides
utilities to control tests - which tests are executed in a CI/CD pipeline
stage as well as the environment in which the tests are executed.  There
are also utility packages that make writing effective tests much easier.

Much of the inspiration for this library cam from Mitchell Hashimoto's
"Advanced Testing with Go" presentation -
https://www.youtube.com/watch?v=8hQG7QlcLBk.  Every Go programmer should
watch this talk for inspiration on how to make test writing pleasant.

There is actually no package named "proctor" (beyond these paragraphs) -
the utilities described above are packaged as follows:

  - flagged: provides utilities for controlling test execution through the
             use of command-line flags.

  - goldenfile: provides utilities for naming, comparing (reading) and
                updating (writing) goldenfiles.

  - log: returns a logrus logger that wraps testing.T.Log() as the output
         stream.
*/
package proctor
