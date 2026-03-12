package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsSameType[T comparable](value1, value2 T) bool {
	return value1 == value2
}

func TestIsSameType(t *testing.T) {
	assert := assert.New(t)

	assert.True(IsSameType[int](1, 1))
	assert.False(IsSameType[string]("a", "b"))
	assert.False(IsSameType[bool](true, false))
}
