package main

import (
	"fmt"
)

func main() {
	names := []string{"Nur", "john", "febry"}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Max(names))
}
