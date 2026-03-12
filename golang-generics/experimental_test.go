package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

func ExperimentalMin[T constraints.Ordered](first, second T) T {
	if first < second {
		return first
	}
	return second
}

func TestExperimentalMin(t *testing.T) {
	assert.Equal(t, 100, ExperimentalMin(100, 200))
	assert.Equal(t, 100.5, ExperimentalMin(100.5, 200.5))
}

func TestExperimentalMaps(t *testing.T) {
	first := map[string]string{
		"Name": "Windah",
	}
	second := map[string]string{
		"Name": "Basudara",
	}
	assert.Equal(t, "Windah", first["Name"])
	assert.Equal(t, "Basudara", second["Name"])
}

func TestExperimentalSlice(t *testing.T) {
	first := []int{1, 2, 3}
	second := []int{4, 5, 6}
	assert.Equal(t, []int{1, 2, 3}, first)
	assert.Equal(t, []int{4, 5, 6}, second)
}
