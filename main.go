package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/peterhellberg/hiro/hiro"
)

// Command line flags
var (
	inputFn    = flag.String("input", "blueprint.md", "Input file (.md)")
	outputFn   = flag.String("output", "index.html", "Output file (.html)")
	templateFn = flag.String("template", "", "Iglo template file")
)

func main() {
	// Command line usage information
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\nCommand line arguments:\n\n")
		flag.PrintDefaults()
	}

	// Parse the command line flags
	flag.Parse()

	// Generate the output file
	err := hiro.Generate(*inputFn, *outputFn, *templateFn)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
