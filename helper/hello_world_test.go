package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldBayu(t *testing.T) {
	result := HelloWorld("Bayu")

	if result != "Hello Bayu" {
		// error
		t.Error("Result must be 'Hello Bayu'")
	}

	fmt.Println("TestHelloWorldBayu Done")

}

func TestHelloWorldSugiantoro(t *testing.T) {
	result := HelloWorld("Sugiantoro")

	if result != "Hello Sugiantoro" {
		// error
		t.Fatal("Result must be 'Hello Sugiantoro'")
	}

	fmt.Println("TestHelloWorldSugiantoro Done")

}