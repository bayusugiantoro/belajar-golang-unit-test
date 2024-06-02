package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Bayu")

	if result != "Hello Bayu" {
		// error
		panic("Result is not 'Hello Bayu'")
	}

}

func TestHelloWorldSugiantoro(t *testing.T) {
	result := HelloWorld("Sugiantoro")

	if result != "Hello Sugiantoro" {
		// error
		panic("Result is not 'Hello Sugiantoro'")
	}

}