package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTableHelloWorld(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Rida",
			request: "Rida",
		},
		{
			name:    "Adila",
			request: "Adila",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Rida", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Rida")
		}
	})

	b.Run("Adila", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Adila")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Rida")
	}
}

func BenchmarkHelloWorldAdila(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Adila")
	}
}

func TestSubTest(t *testing.T) {

	t.Run("Rida", func(t *testing.T) {
		result := HelloWorld("Rida")

		assert.Equal(t, "Hello Rida", result, "Role must be 'Hello Rida'")
		fmt.Println("dieksekusi")
	})

	t.Run("Adila", func(t *testing.T) {
		result := HelloWorld("Adil")

		assert.Equal(t, "Hello Adila", result, "Role must be 'Hello Adila'")
		fmt.Println("dieksekusi")
	})
}

func TestTableHelloWorld(t *testing.T) {

	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Rida",
			request:  "Rida",
			expected: "Hello Rida",
		},
		{
			name:     "Adila",
			request:  "Adila",
			expected: "Hello Adila",
		},
		{
			name:     "Eka",
			request:  "Eka",
			expected: "Hello Eka",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)

			assert.Equal(t, test.expected, result, "Role must be "+test.expected)
			fmt.Println("dieksekusi")
		})
	}
}

func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")
	m.Run()
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	t.Skip("Ini diskip")
	fmt.Println("test skip")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Rid")

	require.Equal(t, "Hello Rida", result, "Role must be 'Hello Rida'")
	fmt.Println("tidak dieksekusi")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Rid")

	assert.Equal(t, "Hello Rida", result, "Role must be 'Hello Rida'")
	fmt.Println("dieksekusi")
}

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
