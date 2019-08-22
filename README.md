# proctor

This project provides functions that allow the control of which tests
are run in each stage of a build pipeline.  The conventions and flags
used to control test execution are based on published best practices
as well as the conventions adopted by Google.

## Background

This section describes the resources that were used to determine what
testing best practices should be included in the library.  Each
section provides a brief synopsis as well as one or more references
to the source materials.

### The test pyramid

Martin Fowler has described the "Test Pyramid"<sup>1</sup> in a post
on his blog as a way of categorizing tests.  The idea of a test pyramid
has been around a long time but he's reduced the testing to three
layers which capture the essence of categorizing tests.  A longer essay
on his blog titled "The Practical Test Pyramid"<sup>2</sup> expands on
these ideas with examples and a detailed description of each type of
test.

The three categories of tests in the modern test pyramid are ``unit``,
``service`` and ``ui``.  Note that in this case, UI might be a thick-
client, a CLI interface, a web page or a REST API.  These categories
are classified from fastest to slowest as well as from least to most
expensive.

These articles rework the classic test pyramid  used in many waterfall
projects which contains, from earliest in the project to latest, unit
tests, integration tests, system tests and acceptance tests.  In an
agile world, thinking about when a CI/CD pipeline would run these tests
makes more sense.

### Google testing practices

A blog post titled "Test Sizes"<sup>3</sup> on the Google Testing Blog
provides a useful way of grouping tests together based on size.  Each
of these "sizes" additionally specifies what system and external
resources are allowed.

In Daniel Danilatos "GWT Testing Best Practices"<sup>4</sup> talk at
Google IO 201070-20-10, he casually mentions that test writers at Google
categorize their tests as ``small``, ``medium`` and ``large`` and that,
in general, the ratio of these tests should be 70%, 20% and 10% respect-
ively.  He further states, that ``small`` tests should be very fast
and have no external dependencies

### Test flags

The ``go test`` command includes a feature that allows test flags to
be passed into tests.  This feature and it's usefulness are described
in Mitchell Hashimoto's "Advanced Testing with Go" presentation<sup>5</sup>.

### Black-box testing

### Design for testability

To a large degree, the fraction of your code that can be tested by
small/unit/short tests is determined by the design of that code.  It's
often possible to dramatically increase the amount of code tested
during this phase simply by changing the design of the software-under-
test and/or the test code.  Most of the "Advanced Testing with Go"
<sup>5</sup> presentation mentioned above is techniques that make your
software more testable.

## Rationale

Putting these ideas together, we should be able to characterize each
of our tests and apply one or more CLI ``flags`` to control when
they are run.  Allowing individual tests to be flagged in multiple
ways provides a fine-grained means of selecting tests (e.g. Not all
large tests are UI tests.)

This library provides a very small amount of code that conditionally
skips tests based on these flags

## Implementation

In addition to the Go test command's built-in ``-short`` flag, this
library provides the following additional test control flags:

| Flags       |
| :---------- |
| Small       |
| Medium      |
| Large       |
|             |
| Unit        |
| Integration |
| System      |
| Acceptance  |
|             |
| Service     |
| UI          |
|             |
| Short       |
| Long        |

Note that ``Short`` and ``Long`` are aliases for ``testing.Short()``
and ``!testing.Short()`` respectively.

## Use

## Custom flags

Additionally, custom flags can be added to a project as needed -
include the following code in the global scope of any of the project's
test files as follows:

```Go
var (
    MyFlag = Flag(flag.Bool("myflag", false, "Run my tests"))
)

func TestMySomething(t *testing.T) {
    flagged.With(t, MyFlag)
    ...
}
```

The preceding code would execute if ``-myflag`` as passed as a CLI
flag to the ``go test`` command.

## Coverage

While blackbox testing is recommended as a means to verify that the
exported portions of the code is consistent with its intended use,
the Go coverage tool and Go test command don't make it clear how
the coverage provided by a black-box test can be included in the
code's overall coverage report.  The simplest way to ensure your
coverage report is accurate is to make your Go code a module and
to execute the following Go test command at the top of your module's
directory hierarchy:

```shell
go test -coverprofile coverage.out -coverpkg ./... ./...
```

Remember to pass along the flags that are appropriate for the
current step in your CI/CD pipeline.

## References

1. [Test Pyramid](https://martinfowler.com/bliki/TestPyramid.html)
2. [The Practical Test Pyramid](https://martinfowler.com/articles/practical-test-pyramid.html)
3. [Test Sizes](https://testing.googleblog.com/2010/12/test-sizes.html)
4. [GWT Testing Best Practices](https://youtu.be/T_CLzgEL7FA?t=62)
5. [Advanced Testing with Go](https://youtu.be/8hQG7QlcLBk?t=733)
