package main

import (
	"bytes"
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		out     string
		wantErr bool
	}{
		{"Valid input with out flag", "testdata/test.md", "testdata/output", false},
		{"Valid input without out flag", "testdata/test.md", "", false},
		{"Empty input", "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := run(tt.in, tt.out)
			if (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				if tt.out != "" {
					os.Remove(tt.out + ".html")
				} else {
					os.Remove("test.md.html")
				}
			}
		})
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
