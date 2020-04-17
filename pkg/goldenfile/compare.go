package goldenfile

import (
	"github.com/stretchr/testify/assert"
)

//
// Define comparators
//

type ByteArrayComparisonAssertionFunc func(t TestingT, expected []byte, actual []byte, msgAndArgs ...interface{}) bool
type StringComparisonAssertionFunc func(t assert.TestingT, expected string, actual string, msgAndArgs ...interface{}) bool

func adaptInterfacesToByteArrays(f assert.ComparisonAssertionFunc) ByteArrayComparisonAssertionFunc {
	return func(t TestingT, expected []byte, actual []byte, msgAndArgs ...interface{}) bool {
		return f(t.(assert.TestingT), interface{}(expected), interface{}(actual), msgAndArgs...)
	}
}

func adaptStringsToByteArrays(f StringComparisonAssertionFunc) ByteArrayComparisonAssertionFunc {
	return func(t TestingT, expected []byte, actual []byte, msgAndArgs ...interface{}) bool {
		return f(t.(assert.TestingT), string(expected), string(actual), msgAndArgs...)
	}
}

var bytesEq ByteArrayComparisonAssertionFunc = adaptInterfacesToByteArrays(assert.Equal)
var jsonEq ByteArrayComparisonAssertionFunc = adaptStringsToByteArrays(assert.JSONEq)
var stringEq ByteArrayComparisonAssertionFunc = adaptInterfacesToByteArrays(assert.Equal)
var yamlEq ByteArrayComparisonAssertionFunc = adaptStringsToByteArrays(assert.YAMLEq)
