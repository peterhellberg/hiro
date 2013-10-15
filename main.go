package main

import (
	"flag"
	"fmt"
	"github.com/subosito/iglo"
	"os"
)

// Command line flags
var inputFile = flag.String("input", "blueprint.md", "Input file (.md)")
var outputFile = flag.String("output", "blueprint.html", "Output file (.html)")

// Simple error check
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Command line usage information
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\nCommand line arguments:\n\n")
		flag.PrintDefaults()
		fmt.Println()
		os.Exit(1)
	}

	// Parse the command line flags
	flag.Parse()

	// Require that both flags are provided
	if flag.NFlag() != 2 {
		flag.Usage()
	}

	// Check if the input file exists
	finfo, err := os.Stat(*inputFile)
	if err != nil {
		// no such file or dir
		fmt.Println("The input file does not exist.")
		return
	}

	if finfo.IsDir() {
		fmt.Println("The input file is a directory.")
		return
	} else {
		// Open the input file
		i, err := os.Open(*inputFile)
		check(err)

		// Make sure the input file is closed later.
		defer i.Close()

		// Create the output file
		o, err := os.Create(*outputFile)
		check(err)

		// Make sure the output file is closed later.
		defer o.Close()

		// Convert the API Blueprint Markdown to HTML
		err = iglo.MarkdownToHTML(o, i)
		check(err)

		// Issue a Sync to flush writes to stable storage.
		o.Sync()
	}
}
