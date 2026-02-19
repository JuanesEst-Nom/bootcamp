package main

import (
	"strings"
	"testing"
)

func TestCountLogic(t *testing.T) {

	// Test cases for word counting
	wordTests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Single sentence",
			input:    "Hello world\nexit",
			expected: 2,
		},
		{
			name:     "Multiple sentences",
			input:    "First line\nSecond line\nexit",
			expected: 4,
		},
		{
			name:     "Single word",
			input:    "Golang\nexit",
			expected: 1,
		},
		{
			name:     "Composed word",
			input:    "This is a read-only task\nexit",
			expected: 5,
		},
		{
			name:     "Multiple blank lines",
			input:    "Hello\n\n\nworld\nexit",
			expected: 2,
		},
		{
			name:     "Case insensitive exit",
			input:    "Hello\nEXIT",
			expected: 1,
		},
	}

	for _, tc := range wordTests {
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			result := CountLogic(reader, false)
			if result != tc.expected {
				t.Errorf("Test [%s] failed: Expected %d words, but got %d", tc.name, tc.expected, result)
			}
		})
	}

	// Test cases for line counting
	lineTests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Single line",
			input:    "Hello world\nexit",
			expected: 1,
		},
		{
			name:     "Multiple lines no breaks",
			input:    "Line1\nLine2\nexit",
			expected: 2,
		},
		{
			name:     "Multiple lines with breaks",
			input:    "Line1\n\nLine2\nexit",
			expected: 3,
		},
		{
			name:     "Exit at start",
			input:    "exit\nHello",
			expected: 0,
		},
		{
			name:     "Exit in middle",
			input:    "Hello\nexit\nWorld",
			expected: 1,
		},
		{
			name:     "Exit case variants",
			input:    "Line1\nExiT",
			expected: 1,
		},
	}

	for _, tc := range lineTests {
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			result := CountLogic(reader, true)
			if result != tc.expected {
				t.Errorf("Test [%s] failed: Expected %d lines, but got %d", tc.name, tc.expected, result)
			}
		})
	}
}
