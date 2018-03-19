package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
	"testing"
)

const (
	testString = "test"
	size       = 100
)

func main() {
	fmt.Print("Concat:        ")
	AllocationsConcat()
	fmt.Print("BytesBuffer:   ")
	AllocationsBuffer()
	fmt.Print("StringBuilder: ")
	AllocationsBuilder()
	fmt.Print("Copy:          ")
	AllocationsCopy()
}

func AllocationsConcat() {
	var start runtime.MemStats
	runtime.ReadMemStats(&start)
	Concat(testString, size)
	var end runtime.MemStats
	runtime.ReadMemStats(&end)
	fmt.Printf("there are %d allocated bytes\n", end.Alloc-start.Alloc)
}
func AllocationsBuffer() {
	var start runtime.MemStats
	runtime.ReadMemStats(&start)
	Buffer(testString, size)
	var end runtime.MemStats
	runtime.ReadMemStats(&end)
	fmt.Printf("there are %d allocated bytes\n", end.Alloc-start.Alloc)
}
func AllocationsCopy() {
	var start runtime.MemStats
	runtime.ReadMemStats(&start)
	Copy(testString, 10)
	var end runtime.MemStats
	runtime.ReadMemStats(&end)
	fmt.Printf("there are %d allocated bytes\n", end.Alloc-start.Alloc)
}

func AllocationsBuilder() {
	var start runtime.MemStats
	runtime.ReadMemStats(&start)
	StringBuilder(testString, size)
	var end runtime.MemStats
	runtime.ReadMemStats(&end)
	fmt.Printf("there are %d allocated bytes\n", end.Alloc-start.Alloc)
}

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

func benchmark(size int, testFunc func(string, int) string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		testFunc(testString, size)
	}
}

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
