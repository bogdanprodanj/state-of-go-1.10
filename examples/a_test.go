package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
	"testing"
)

func main() {
	fmt.Println("Concat")
	a1()
	fmt.Println("Buffer")
	a2()
	fmt.Println("Copy")
	a3()
	fmt.Println("StringBuilder")
	a4()
}

func a1() {
	var start runtime.MemStats
	runtime.ReadMemStats(&start)
	fmt.Println(Concat(testString, 10))
	var end runtime.MemStats
	runtime.ReadMemStats(&end)
	fmt.Printf("there are %d allocated bytes\n", end.Alloc-start.Alloc)
}
func a2() {
	var start runtime.MemStats
	runtime.ReadMemStats(&start)
	fmt.Println(Buffer(testString, 10))
	var end runtime.MemStats
	runtime.ReadMemStats(&end)
	fmt.Printf("there are %d allocated bytes\n", end.Alloc-start.Alloc)
}
func a3() {
	var start runtime.MemStats
	runtime.ReadMemStats(&start)
	fmt.Println(Copy(testString, 10))
	var end runtime.MemStats
	runtime.ReadMemStats(&end)
	fmt.Printf("there are %d allocated bytes\n", end.Alloc-start.Alloc)
}

func a4() {
	var start runtime.MemStats
	runtime.ReadMemStats(&start)
	fmt.Println(StringBuilder(testString, 10))
	var end runtime.MemStats
	runtime.ReadMemStats(&end)
	fmt.Printf("there are %d allocated bytes\n", end.Alloc-start.Alloc)
}

// START1 OMIT
func Concat(input string, size int) string {
	var str string
	for n := 0; n < size; n++ {
		str += input
	}
	return str
}

func Buffer(input string, size int) string {
	var buffer bytes.Buffer
	for n := 0; n < size; n++ {
		buffer.WriteString(input)
	}
	return buffer.String()
}

// END1 OMIT

// START2 OMIT
func Copy(input string, size int) string {
	bs := make([]byte, size*len(input))
	length := 0
	for n := 0; n < size; n++ {
		length += copy(bs[length:], input)
	}
	return string(bs)
}

func StringBuilder(input string, size int) string {
	var strBuilder strings.Builder
	for n := 0; n < size; n++ {
		strBuilder.WriteString(input)
	}
	return strBuilder.String()
}

// END2 OMIT

const testString = "test"

// START3 OMIT
func benchmark(size int, testFunc func(string, int) string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		testFunc("test", size)
	}
}

// END3 OMIT

func BenchmarkConcat2(b *testing.B)       { benchmark(2, Concat, b) }
func BenchmarkConcat100(b *testing.B)     { benchmark(100, Concat, b) }
func BenchmarkConcat100000(b *testing.B)  { benchmark(100000, Concat, b) }
func BenchmarkBuffer2(b *testing.B)       { benchmark(2, Buffer, b) }
func BenchmarkBuffer100(b *testing.B)     { benchmark(100, Buffer, b) }
func BenchmarkBuffer100000(b *testing.B)  { benchmark(100000, Buffer, b) }
func BenchmarkCopy2(b *testing.B)         { benchmark(2, Copy, b) }
func BenchmarkCopy100(b *testing.B)       { benchmark(100, Copy, b) }
func BenchmarkCopy100000(b *testing.B)    { benchmark(100000, Copy, b) }
func BenchmarkBuilder2(b *testing.B)      { benchmark(2, StringBuilder, b) }
func BenchmarkBuilder100(b *testing.B)    { benchmark(100, StringBuilder, b) }
func BenchmarkBuilder100000(b *testing.B) { benchmark(100000, StringBuilder, b) }
