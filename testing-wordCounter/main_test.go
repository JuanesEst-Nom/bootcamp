package main

import (
	"strings"
	"testing"
)

func TestWordCounter(t *testing.T) {

	// test word cases
	wordTests := []struct {
		name     string
		input    string
		expected int
	}{
		{"Single sentence", "Hola mundo\nexit", 2},
		{"Multiple sentences", "Primera linea\nSegunda linea\nexit", 4},
		{"Single word", "Golang\nexit", 1},
		{"Composed word", "Esta es una tarea read-only\nexit", 5},
		{"Multiple break lines", "Hola\n\n\nmundo\nexit", 2},
		{"Case insensitive exit", "Hola\nEXIT", 1},
	}

	for _, tc := range wordTests {
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			result := CountLogic(reader, false)
			if result != tc.expected {
				t.Errorf("Operation failed [%s]: Expected %d words, but got %d", tc.name, tc.expected, result)
			}
		})
	}

	// test by lines
	lineTests := []struct {
		name     string
		input    string
		expected int
	}{
		{"Single line", "Hola mundo\nexit", 1},
		{"Multiple lines no breaks", "Linea1\nLinea2\nexit", 2},
		{"Multiple lines with breaks", "Linea1\n\nLinea2\nexit", 3},
		{"Exit at start", "exit\nHola", 0},
		{"Exit in middle", "Hola\nexit\nMundo", 1},
		{"Exit case variants", "Linea1\nExiT", 1},
	}

	for _, tc := range lineTests {
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			result := CountLogic(reader, true)
			if result != tc.expected {
				t.Errorf("Operation failed [%s]: Expected %d lines, but got %d", tc.name, tc.expected, result)
			}
		})
	}
}
