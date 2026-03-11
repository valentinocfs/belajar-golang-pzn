package test

import (
	"fmt"
	"golang-dependency-injection/simple"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSimpleServiceSuccess(t *testing.T) {
	simpleService, err := simple.InitializeService(false)
	if err != nil {
		panic(err)
	}
	fmt.Println(simpleService.SimpleRepository)
}

func TestSimpleServiceError(t *testing.T) {
	simpleService, _ := simple.InitializeService(true)
	assert.Equal(t, nil, simpleService)
	// 	assert.Equal(t, "Failed create services", err)
}
