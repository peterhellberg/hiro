package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/peterhellberg/hiro/hiro"
)

// Command line flags
var inputFile = flag.String("input", "blueprint.md", "Input file (.md)")
var outputFile = flag.String("output", "index.html", "Output file (.html)")

func main() {
	// Command line usage information
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\nCommand line arguments:\n\n")
		flag.PrintDefaults()
		os.Exit(0)
	}

	// Parse the command line flags
	flag.Parse()

	// Generate the output file
	err := hiro.Generate(*inputFile, *outputFile)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
