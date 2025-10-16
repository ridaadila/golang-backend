package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Rida")

	if result != "Hello Rida" {
		t.Error("Result must be Hello Rida")
	}

	fmt.Println("Fail tetap jalan")
}

func TestHelloWorldNew(t *testing.T) {
	result := HelloWorld("Adila")

	if result != "Hello Adila" {
		t.Fatal("Result must be Hello Adila")
	}

	fmt.Println("Fail now tidak jalan")
}
