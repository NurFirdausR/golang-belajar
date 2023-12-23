package belajar_golang_generics

import (
	"fmt"
	"testing"
)

func MultiParams[T1 any, T2 any, T3 any](nama T1, umur T2, kelas T3) {
	fmt.Println(nama)
	fmt.Println(umur)
	fmt.Println(kelas)
}

func TestMultiParams(t *testing.T) {
	MultiParams[string, int16]("nur", 20, "Kuliah")
}
