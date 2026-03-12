package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee interface {
	GetName() string
}

func GetName[T Employee](param T) string {
	return param.GetName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

type VicePresident interface {
	GetName() string
	GetVicePresidentName() string
}

type MyVicePresident struct {
	Name string
}

func (v *MyVicePresident) GetName() string {
	return v.Name
}

func (v *MyVicePresident) GetVicePresidentName() string {
	return v.Name
}

func TestGetName(t *testing.T) {
	assert.Equal(t, "Windah", GetName[Manager](&MyManager{Name: "Windah"}))
	assert.Equal(t, "Basudara", GetName[VicePresident](&MyVicePresident{Name: "Basudara"}))
}
