package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	value := "Nur Firdaus R"

	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	decode, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(string(decode))
	}
	// fmt.Println(decode)
}
