package hiro

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/peterhellberg/iglo"
)

// Errors
var (
	errFileMissing = errors.New("file does not exist")
	errFileIsDir   = errors.New("file is a directory")
)

// Generate generates HTML output from Markdown input
func Generate(inputFn, outputFn, templateFn string) error {
	// Get the template string
	tmpl, err := readTemplate(templateFn)
	if err != nil {
		return err
	}
	// Set the iglo template
	iglo.Tmpl = string(tmpl)

	// Open the input file
	input, err := openFile(inputFn)
	if err != nil {
		return err
	}
	// Make sure the input file is closed later.
	defer input.Close()

	// Create the output file
	output, err := os.Create(outputFn)
	if err != nil {
		return err
	}
	// Make sure the output file is closed later.
	defer output.Close()

	// Convert the API Blueprint Markdown to HTML
	return iglo.MarkdownToHTML(output, input)
}

func readTemplate(fn string) ([]byte, error) {
	// Return the default template if passed empty filename
	if fn == "" {
		return []byte(DefaultTemplate), nil
	}

	return ioutil.ReadFile(fn)
}

func openFile(fn string) (*os.File, error) {
	// Check if the file exists
	finfo, err := os.Stat(fn)
	if err != nil {
		return nil, errFileMissing
	}

	if finfo.IsDir() {
		return nil, errFileIsDir
	}

	// Open the file
	return os.Open(fn)
}
