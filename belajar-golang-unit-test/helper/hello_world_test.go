package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result)
	fmt.Println("Test Hello World Eko Success")
}

func TestHelloWorldEko(t *testing.T) {
	result := HelloWorld("Eko")

	if result != "Hello Eko" {
		//error
		t.Fail()
	}
	fmt.Println("Test Hello World Eko Success")
}

func TestHelloWorldNur(t *testing.T) {
	result := HelloWorld("Nur")

	if result != "Hello Nur" {
		//error
		t.FailNow()
	}
	fmt.Println("Test Hello World Nur Success")
}
