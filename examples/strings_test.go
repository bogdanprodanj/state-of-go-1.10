package main

import (
	"bytes"
	"strings"
	"testing"
)

const testString = "test"

func Concat(input string, size int) string {
	var str string
	for n := 0; n < size; n++ {
		str += input
	}
	return str
}

func Append(input string, size int) string {
	var data []byte
	for n := 0; n < size; n++ {
		data = append(data, input...)
	}
	return string(data)
}

func Copy(input string, size int) string {
	data := make([]byte, 0, 64)
	var off int
	for n := 0; n < size; n++ {
		off = len(data)
		if off+len(input) > cap(data) {
			temp := make([]byte, 2*cap(data)+len(input))
			copy(temp, data)
			data = temp
		}
		data = data[0 : off+len(input)]
		copy(data[off:], input)
	}
	return string(data)
}

func Buffer(input string, size int) string {
	var buffer bytes.Buffer
	for n := 0; n < size; n++ {
		buffer.WriteString(input)
	}
	return buffer.String()
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

func BenchmarkConcat2(b *testing.B)  { benchmark(2, Concat, b) }
func BenchmarkAppend2(b *testing.B)  { benchmark(2, Append, b) }
func BenchmarkCopy2(b *testing.B)    { benchmark(2, Copy, b) }
func BenchmarkBuffer2(b *testing.B)  { benchmark(2, Buffer, b) }
func BenchmarkBuilder2(b *testing.B) { benchmark(2, StringBuilder, b) }

func BenchmarkConcat100(b *testing.B)  { benchmark(100, Concat, b) }
func BenchmarkAppend100(b *testing.B)  { benchmark(100, Append, b) }
func BenchmarkCopy100(b *testing.B)    { benchmark(100, Copy, b) }
func BenchmarkBuffer100(b *testing.B)  { benchmark(100, Buffer, b) }
func BenchmarkBuilder100(b *testing.B) { benchmark(100, StringBuilder, b) }

func BenchmarkConcat100000(b *testing.B)  { benchmark(100000, Concat, b) }
func BenchmarkAppend100000(b *testing.B)  { benchmark(100000, Append, b) }
func BenchmarkCopy100000(b *testing.B)    { benchmark(100000, Copy, b) }
func BenchmarkBuffer100000(b *testing.B)  { benchmark(100000, Buffer, b) }
func BenchmarkBuilder100000(b *testing.B) { benchmark(100000, StringBuilder, b) }
