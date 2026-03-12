package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Number interface {
	~int | int8 | int16 | int32 | int64 | float32 | float64
}

type Age int

func Min[T Number](first, second T) T {
	if first < second {
		return first
	}
	return second
}

func TestTypeSets(t *testing.T) {
	assert.Equal(t, 0.5, Min(1, 0.5))
	assert.Equal(t, 0.5, Min(3.14, 0.5))
}

func TestTypeDeclaration(t *testing.T) {
	assert.Equal(t, Age(100), Min(Age(100), Age(200)))
}

func TestTypeInference(t *testing.T) {
	assert.Equal(t, int(100), Min(int(100), int(200)))
	assert.Equal(t, Age(100), Min(Age(100), Age(200)))
}
