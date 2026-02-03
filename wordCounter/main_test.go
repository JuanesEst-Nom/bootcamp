package main

import (
	"strings"
	"testing"
)

func TestCount(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		lines    bool
		bytes    bool
		expected int
	}{
		// ESCENARIOS DE PALABRAs
		{name: "Palabras: Frase simple", input: "hola mundo go", lines: false, bytes: false, expected: 3},
		{name: "Palabras: Múltiples oraciones", input: "Hola. Como estas? Todo bien.", lines: false, bytes: false, expected: 5},
		{name: "Palabras: Una sola palabra", input: "Golang", lines: false, bytes: false, expected: 1},
		{name: "Palabras: Palabra compuesta", input: "read-only", lines: false, bytes: false, expected: 1},
		{name: "Palabras: Múltiples saltos de línea", input: "uno\n\n\ndos", lines: false, bytes: false, expected: 2},
		{name: "Palabras: Palabra Exit", input: "Exit", lines: false, bytes: false, expected: 1},

		// ESCENARIOS DE LÍNEAS
		{name: "Líneas: Una sola línea", input: "una linea", lines: true, bytes: false, expected: 1},
		{name: "Líneas: Múltiples líneas seguidas", input: "linea1\nlinea2\nlinea3", lines: true, bytes: false, expected: 3},
		{name: "Líneas: Líneas con saltos extra", input: "linea1\n\nlinea2", lines: true, bytes: false, expected: 3},
		{name: "Líneas: Exit en diferentes posiciones", input: "Exit al inicio\nMedio Exit aqui\nFinal Exit", lines: true, bytes: false, expected: 3},

		// ESCENARIO DE BYTES
		{name: "Bytes: Contar bytes", input: "abc", lines: false, bytes: true, expected: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Simulamos la entrada del usuario usando strings.NewReader
			reader := strings.NewReader(tt.input)

			result := count(reader, tt.lines, tt.bytes)

			if result != tt.expected {
				t.Errorf("Operación fallida: %s\nInput: %q\nFlags: lines=%v, bytes=%v\nResultado: %d, Esperado: %d",
					tt.name, tt.input, tt.lines, tt.bytes, result, tt.expected)
			}
		})
	}
}
