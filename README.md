# test-control

This project provides functions that allow the control of which tests
are run in each stage of a build pipeline.  The conventions and flags
used to control test execution are based on published best practices
as well as the conventions adopted by Google.

## The test pyramid

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
agile world, thinking about when in a CI/CD pipeline these tests should
be run makes more sense.

## Google testing practices

A blog post titled "Test Sizes"<sup>1</sup> on the Google Testing Blog
provides a useful way of grouping tests together based on size.

In Daniel Danilatos "GWT Testing Best Practices"<sup>2</sup> talk at
Google IO 201070-20-10, he casually mentions that test writers at Google
categorize their tests as ``small``, ``medium`` and ``large`` and that,
in general, the ratio of these tests should be 70%, 20% and 10% respect-
ively.  He further states, that<sup>2</sup>

References

1. [Test Pyramid](https://martinfowler.com/bliki/TestPyramid.html)
2. [The Practical Test Pyramid](https://martinfowler.com/articles/practical-test-pyramid.html)
3. [Test Sizes](https://testing.googleblog.com/2010/12/test-sizes.html)
4. [GWT Testing Best Practices](https://youtu.be/T_CLzgEL7FA?t=62)
5. [Advanced Testing with Go](https://youtu.be/8hQG7QlcLBk?t=733)
