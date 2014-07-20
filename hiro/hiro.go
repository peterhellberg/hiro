package hiro

import (
	"errors"
	"os"

	"github.com/subosito/iglo"
)

func Generate(inputFile, outputFile string) error {
	// Open the input file
	input, err := openInputFile(inputFile)
	if err != nil {
		return err
	}
	// Make sure the input file is closed later.
	defer input.Close()

	// Create the output file
	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	// Make sure the output file is closed later.
	defer output.Close()

	// Convert the API Blueprint Markdown to HTML
	return iglo.MarkdownToHTML(output, input)
}

func openInputFile(inputFile string) (*os.File, error) {
	// Check if the input file exists
	finfo, err := os.Stat(inputFile)
	if err != nil {
		return nil, errors.New("The input file does not exist.")
	}

	if finfo.IsDir() {
		return nil, errors.New("The input file is a directory.")
	}

	// Open the input file
	return os.Open(inputFile)
}
