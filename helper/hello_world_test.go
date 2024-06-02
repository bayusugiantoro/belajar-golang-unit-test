package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on Windows OS")
	}

	result := HelloWorld("Bayu")
	require.Equal(t, "Hello Bayu", result, "Result must be 'Hello Bayu'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Bayu")
	require.Equal(t, "Hello Bayu", result, "Result must be 'Hello Bayu'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Bayu")
	assert.Equal(t, "Hello Bayu", result, "Result must be 'Hello Bayu'")
	fmt.Println("TestHelloWorld with Assert Done")
}

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