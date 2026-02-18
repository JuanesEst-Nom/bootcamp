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
		{"Single sentence", "Hello world\nexit", 2},
		{"Multiple sentences", "First line\nSecond line\nexit", 4},
		{"Single word", "Golang\nexit", 1},
		{"Composed word", "This is a read-only task\nexit", 5},
		{"Multiple blank lines", "Hello\n\n\nworld\nexit", 2},
		{"Case insensitive exit", "Hello\nEXIT", 1},
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
		{"Single line", "Hello world\nexit", 1},
		{"Multiple lines no breaks", "Line1\nLine2\nexit", 2},
		{"Multiple lines with breaks", "Line1\n\nLine2\nexit", 3},
		{"Exit at start", "exit\nHello", 0},
		{"Exit in middle", "Hello\nexit\nWorld", 1},
		{"Exit case variants", "Line1\nExiT", 1},
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
