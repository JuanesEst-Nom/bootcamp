package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const header = `<!DOCTYPE html>
  <html>
    <head>
      <meta http-equiv="content-type" content="text/html; charset=utf-8" />
      <title>Markdown Preview Tool</title>
    </head>
    <body>
`

const footer = `
    </body>
  </html>
`

func parseContent(input []byte) []byte {
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)
	return append([]byte(header), append(body, []byte(footer)...)...)
}

func run(in, out string, writer io.Writer) error {
	if in == "" {
		return fmt.Errorf("input file is required")
	}

	input, err := os.ReadFile(in)
	if err != nil {
		return fmt.Errorf("failed to read input file: %v", err)
	}

	var outFile string
	if out != "" {
		outFile = out + ".html"
	} else {
		tf, err := os.CreateTemp(".", "md*.html")
		if err != nil {
			return fmt.Errorf("failed to create temp file: %v", err)
		}
		tf.Close()
		outFile = tf.Name()
	}

	data := parseContent(input)

	if err := saveHTML(data, outFile); err != nil {
		return fmt.Errorf("failed to save HTML: %v", err)
	}

	fmt.Fprintln(writer, outFile)
	return nil
}

func saveHTML(data []byte, nameFile string) error {
	err := os.WriteFile(nameFile, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}
	return nil
}

func main() {
	in := flag.String("in", "", "Path to the input markdown file")
	out := flag.String("out", "", "Name of the output HTML file (optional)")
	flag.Parse()

	if err := run(*in, *out, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
