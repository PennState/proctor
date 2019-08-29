package goldenfile

import (
	"github.com/stretchr/testify/assert"
)

//
// Define comparators
//

type GoldenFileComparisonAssertionFunc func(t TestingT, expected []byte, actual []byte, msgAndArgs ...interface{}) bool
type StringComparisonAssertionFunc func(t assert.TestingT, expected string, actual string, msgAndArgs ...interface{}) bool

func adaptInterfacesToByteArrays(f assert.ComparisonAssertionFunc) GoldenFileComparisonAssertionFunc {
	return func(t TestingT, expected []byte, actual []byte, msgAndArgs ...interface{}) bool {
		return f(t.(assert.TestingT), interface{}(expected), interface{}(actual), msgAndArgs...)
	}
}

func adaptStringsToByteArrays(f StringComparisonAssertionFunc) GoldenFileComparisonAssertionFunc {
	return func(t TestingT, expected []byte, actual []byte, msgAndArgs ...interface{}) bool {
		return f(t.(assert.TestingT), string(expected), string(actual), msgAndArgs...)
	}
}

var bytesEq GoldenFileComparisonAssertionFunc = adaptInterfacesToByteArrays(assert.Equal)
var jsonEq GoldenFileComparisonAssertionFunc = adaptStringsToByteArrays(assert.JSONEq)
var stringEq GoldenFileComparisonAssertionFunc = adaptInterfacesToByteArrays(assert.Equal)
var yamlEq GoldenFileComparisonAssertionFunc = adaptStringsToByteArrays(assert.YAMLEq)
