package belajar_golang_generics

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Compare[T comparable](val1 T, val2 T) bool {
	return reflect.DeepEqual(val1, val2)
}

func TestCompare(t *testing.T) {
	assert.Equal(t, true, Compare[string]("nur", "nur"))
}
