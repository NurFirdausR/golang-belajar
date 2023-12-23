package belajar_golang_generics

import (
	"fmt"
	"testing"
)

func Length[T any](params T) T {
	// fmt.Println(params)
	return params
}

func TestLength(t *testing.T) {
	result := Length[string]("hello cuk")
	result2 := Length[int16](20)

	fmt.Println(result)
	fmt.Println(result2)
}
