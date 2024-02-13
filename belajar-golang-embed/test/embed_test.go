package test

import (
	"embed"
	_ "embed"
	"fmt"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}

//go:embed kehati.png
var logo []byte

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("kehati_new.png", logo, 0644)
	if err != nil {
		panic(err)
	}
	// fmt.Println(logo)
}

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestMultipleFile(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))

}

//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
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
