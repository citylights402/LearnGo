package e

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

	new_assert := assert.New(t) //New makes a new Assertions object for the specified TestingT.

	var a string = "Hello"
	var b string = "Hell"
	assert.Equal(t, a, b, "The two words should be the same1.")
	new_assert.Equal(a, b, "The two words should be the same2.")
}
