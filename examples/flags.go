package main

import (
	"flag"
)

func main() {
	// START OMIT
	parameter := flag.Int("parameter", 0, "some parameter\nwith very long description")
	number := flag.Float64("number", 1.1, "some number")
	flag.Parse()
	// END OMIT
	flag.PrintDefaults()
	_, _ = parameter, number
}
