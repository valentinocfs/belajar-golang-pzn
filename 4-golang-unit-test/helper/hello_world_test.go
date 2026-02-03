package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "John",
			request: "John",
		},
		{
			name:    "Jane Doe",
			request: "Jane Doe",
		},
		{
			name:    "John F Kennedy",
			request: "John F Kennedy",
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

func BenchmarkSubBenchmark(b *testing.B) {
	b.Run("Test_1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("John")
		}
	})

	b.Run("Test_2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Jane Doe")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("John Doe")
	}
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "HelloWorld(John Doe)",
			input:    "John Doe",
			expected: "Hello John Doe",
		},
		{
			name:     "HelloWorld(Jane Doe)",
			input:    "Jane Doe",
			expected: "Hello Jane Doe",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.input)
			require.Equal(t, test.expected, result, "Result is not %s", test.expected)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Test_1", func(t *testing.T) {
		result := HelloWorld("John Doe")
		require.Equal(t, "Hello John Doe", result, "Result is not Hello John Doe")
	})

	t.Run("Test_2", func(t *testing.T) {
		result := HelloWorld("John Doe")
		require.Equal(t, "Hello John Doe", result, "Result is not Hello John Doe")
	})
}

func TestMain(m *testing.M) {
	// before test
	fmt.Println("Before Unit Test")

	m.Run()

	// after test
	fmt.Println("After Unit Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on the Mac OS")
	}

	result := HelloWorld("John Doe")
	require.Equal(t, "Hello John Doe", result, "Result is not Hello John Doe")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("John Doe")
	require.Equal(t, "Hello John Doe", result, "Result is not Hello John Doe")
	fmt.Println("TestHelloWorld Require Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("John Doe")
	assert.Equal(t, "Hello John Doe", result, "Result is not Hello John Doe")
	fmt.Println("TestHelloWorld Assert Done")
}

func TestHelloWorldFail(t *testing.T) {
	result := HelloWorld("John Doe")
	if result != "Hello John Doe" {
		t.Fail()
	}
	fmt.Println("TestHelloWorld Fail Done")
}

func TestHelloWorldFailNow(t *testing.T) {
	result := HelloWorld("John Doe")
	if result != "Hello John Doe" {
		t.FailNow()
	}
	fmt.Println("TestHelloWorld FailNow Done")
}

func TestHelloWorldError(t *testing.T) {
	result := HelloWorld("John Doe")
	if result != "Hello John Doe" {
		t.Error("Result is not Hello John Doe")
	}
	fmt.Println("TestHelloWorld Error Done")
}

func TestHelloWorldFatal(t *testing.T) {
	result := HelloWorld("John Doe")
	if result != "Hello John Doe" {
		t.Fatal("Result is not Hello John Doe")
	}
	fmt.Println("TestHelloWorld Fatal Done")
}
