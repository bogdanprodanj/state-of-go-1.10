package main

import (
	"bytes"
	"fmt"
)

func main() {
	// START OMIT
	abc := []byte("Hello World")
	fmt.Printf("abc: %s", desc(abc))

	xyz := bytes.Fields(abc)
	fmt.Println("Before change")
	fmt.Printf("[0]: %s", desc(xyz[0]))
	fmt.Printf("[1]: %s", desc(xyz[1]))

	xyz[0] = append(xyz[0], byte('X'))
	xyz[0] = append(xyz[0], byte('Y'))
	xyz[0] = append(xyz[0], byte('Z'))
	fmt.Println("After change")
	fmt.Printf("[0]: %s", desc(xyz[0]))
	fmt.Printf("[1]: %s", desc(xyz[1]))
	// END OMIT
}
