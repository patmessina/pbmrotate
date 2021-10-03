package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/patmessina/pbmrotate/pkg/p1"
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

	image, err := p1.NewImageFromFile(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	image.Rotate(deg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if output != "" {
		image.WriteToFile(output)
	} else {
		image.Print()
	}

}
