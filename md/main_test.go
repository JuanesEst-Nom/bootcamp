package main

import (
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"Valid input", "output", false},
		{"Empty input", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := run(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				os.Remove(tt.input + ".html")
			}
		})
	}
}
