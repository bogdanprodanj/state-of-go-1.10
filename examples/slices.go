package main

import "fmt"

func main() {
	// START OMIT
	text := []byte("Hello Lviv!")
	fmt.Printf("text:  %s", desc(text))

	hello := text[0:5]
	fmt.Printf("hello: %s", desc(hello))

	hello = append(hello, '#')
	fmt.Printf("hello: %s", desc(hello))

	fmt.Printf("text:  %s", desc(text))
	// END OMIT
}
