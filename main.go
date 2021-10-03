package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	output string
	deg    int
	input  string
)

func main() {

	// get user inputs
	flag.StringVar(&input, "i", "", "filepath to image input")
	flag.StringVar(&output, "o", "", "image output filepath (default: '')")
	flag.IntVar(&deg, "d", 0, "degrees to rotate (default: 0)")
	flag.Parse()

	if input == "" {
		fmt.Fprintln(os.Stderr, "input is required")
		flag.PrintDefaults()
		os.Exit(1)
	}

}
