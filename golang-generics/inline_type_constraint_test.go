package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func FindMin[T interface{ int64 | float64 }](first, second T) T {
	if first < second {
		return first
	}
	return second
}

func TestFindMin(t *testing.T) {
	assert.Equal(t, int64(1), FindMin(int64(1), int64(2)))
	assert.Equal(t, float64(1.5), FindMin(float64(1.5), float64(2.5)))
}

func GetFirstValue[T []E, E any](data T) E {
	return data[0]
}

func TestGetFirstValue(t *testing.T) {
	assert.Equal(t, int64(1), GetFirstValue([]int64{1, 2, 3}))
	assert.Equal(t, float64(1.5), GetFirstValue([]float64{1.5, 2.5, 3.5}))
	assert.Equal(t, "Windah", GetFirstValue([]string{"Windah", "Basudara"}))
}
