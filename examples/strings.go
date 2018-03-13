package main

import (
	"bytes"
	"fmt"
	"runtime"
)

func main() {
	var start runtime.MemStats
	runtime.ReadMemStats(&start)

	// START OMIT
	var buf bytes.Buffer // HL
	fmt.Fprintln(&buf, "Hello, Lviv gophers!")
	fmt.Print(buf.String())
	// END OMIT
	var end runtime.MemStats
	runtime.ReadMemStats(&end)
	fmt.Printf("there are %d allocated bytes\n", end.Alloc-start.Alloc)
}