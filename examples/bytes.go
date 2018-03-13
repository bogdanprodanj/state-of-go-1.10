package main

import (
	"bytes"
	"fmt"
)

func main() {
	// START OMIT
	slice1 := []byte("Hello Lviv!")
	fmt.Printf("slice1: %s", desc(slice1))

	// Fields splits the slice2 around each instance of one or more consecutive white space
	// characters, returning a slice of subslices of slice2 or
	// an empty list if s contains only white space.
	slice2 := bytes.Fields(slice1)[0]
	fmt.Printf("slice2: %s", desc(slice2))

	slice2 = append(slice2, '#')
	fmt.Printf("slice2: %s", desc(slice2))
	// END OMIT
}

func desc(b []byte) string {
	return fmt.Sprintf("len: %2d | cap: %2d | %q\n", len(b), cap(b), b)
}
