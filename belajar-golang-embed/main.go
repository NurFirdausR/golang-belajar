package main

import (
	"embed"
	"fmt"
	"io/ioutil"
)

//go:embed version.txt
var version string

//go:embed kehati.png
var logo []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	err := ioutil.WriteFile("kehati_new.png", logo, 0644)
	if err != nil {
		panic(err)
	}

	dirEntries, _ := path.ReadDir("files")
	// dirEntries.
	for _, entry := range dirEntries {
		// entry.
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			a, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(a))

		}
	}

}
