package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkSub(b *testing.B) {
	b.Run("Bayu", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Bayu")
		}
	})
	b.Run("Sugiantoro", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Sugiantoro")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Bayu")
	}
}

func BenchmarkHelloWorldSugiantoro(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Sugiantoro")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Eko",
			request:  "Eko",
			expected: "Hello Eko",
		},
		{
			name:     "Kurniawan",
			request:  "Kurniawan",
			expected: "Hello Kurniawan",
		},
		{
			name:     "Khannedy",
			request:  "Khannedy",
			expected: "Hello Khannedy",
		},
		{
			name:     "Bayu",
			request:  "Bayu",
			expected: "Hello Bayu",
		},
		{
			name:     "Sugiantoro",
			request:  "Sugiantoro",
			expected: "Hello Sugiantoro",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Bayu", func(t *testing.T) {
		result := HelloWorld("Bayu")
		require.Equal(t, "Hello Bayu", result, "Result must be 'Hello Bayu'")
	})
	t.Run("Sugiantoro", func(t *testing.T) {
		result := HelloWorld("Sugiantoro")
		require.Equal(t, "Hello Sugiantoro", result, "Result must be 'Hello Sugiantoro'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

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