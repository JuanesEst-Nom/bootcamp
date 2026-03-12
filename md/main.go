package main

import (
	"flag"
	"fmt"
	"os"
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

func run(name string) error {

	if name == "" {
		return fmt.Errorf("output file name is required")
	}

	ext := name + ".html"

	body := header + footer
	data := []byte(body)

	if err := saveHTML(data, ext); err != nil {
		return fmt.Errorf("failed to save HTML: %v", err)
	}

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

	name := flag.String("out", "", "Name of the output HTML file")
	flag.Parse()

	if err := run(*name); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
