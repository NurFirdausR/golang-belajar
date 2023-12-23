package belajar_golang_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee interface {
	Getname() string
}

func Getname[T Employee](parameter T) string {
	return parameter.Getname()
}

type Manager interface {
	Getname() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) Getname() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

type VicePresident interface {
	Getname() string
	GetVicePresidentName() string
}

type MyVicePresident struct {
	Name string
}

func (m *MyVicePresident) Getname() string {
	return m.Name
}

func (m *MyVicePresident) GetVicePresidentName() string {
	return m.Name
}

func TestGetName(t *testing.T) {
	assert.Equal(t, "Nur", Getname[Manager](&MyManager{Name: "Nur"}))
	assert.Equal(t, "Nur", Getname[VicePresident](&MyVicePresident{Name: "Nur"}))
}
