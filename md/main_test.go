package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestRunUsingOutFlag(t *testing.T) {
	var buf bytes.Buffer
	outFile := "testdata/output"

	if err := run("testdata/test.md", outFile, &buf); err != nil {
		t.Fatalf("run() error = %v", err)
	}
	defer os.Remove(outFile + ".html")

	if _, err := os.Stat(outFile + ".html"); os.IsNotExist(err) {
		t.Errorf("expected file %s.html to exist", outFile)
	}

	printed := strings.TrimSpace(buf.String())
	if printed != outFile+".html" {
		t.Errorf("expected printed file name '%s.html', got '%s'", outFile, printed)
	}
}

func TestRunWithoutOutFlag(t *testing.T) {
	var buf bytes.Buffer

	if err := run("testdata/test.md", "", &buf); err != nil {
		t.Fatalf("run() error = %v", err)
	}

	outFile := strings.TrimSpace(buf.String())
	defer os.Remove(outFile)

	if outFile == "" {
		t.Fatal("expected file name to be printed, got empty string")
	}

	if _, err := os.Stat(outFile); os.IsNotExist(err) {
		t.Errorf("expected temp file %s to exist", outFile)
	}
}

func TestParseContent(t *testing.T) {
	input, err := os.ReadFile("testdata/test.md")
	if err != nil {
		t.Fatalf("failed to read test markdown file: %v", err)
	}

	got := parseContent(input)

	golden, err := os.ReadFile("testdata/test.md.html")
	if err != nil {
		t.Fatalf("failed to read golden file: %v", err)
	}

	if !bytes.Equal(got, golden) {
		t.Errorf("parseContent() output does not match golden file.\nGot:\n%s\nWant:\n%s", got, golden)
	}
}
