package main

import (
	"reflect"
	"testing"
)

// aqui declaramos la func de testeo
func TestMinMax(t *testing.T) {
	// Definimos los escenarios de prueba
	tests := []struct {
		name     string
		min      float64
		max      float64
		values   []float64
		expected []float64
	}{
		{
			name:     "Valores estándar dentro del rango",
			min:      10.0,
			max:      20.0,
			values:   []float64{5.0, 10.0, 15.0, 20.0, 25.0},
			expected: []float64{10.0, 15.0, 20.0},
		},
		{
			name:     "Un solo valor dentro del rango",
			min:      10.0,
			max:      20.0,
			values:   []float64{15.0},
			expected: []float64{15.0},
		},
		{
			name:     "Min mayor que Max (rango imposible)",
			min:      50.0,
			max:      10.0,
			values:   []float64{15.0, 25.0, 60.0},
			expected: nil,
		},
		{
			name:     "Ningún valor dentro del rango",
			min:      100.0,
			max:      200.0,
			values:   []float64{10.0, 20.0, 30.0},
			expected: nil,
		},
		{
			name:     "Min y Max negativos",
			min:      -50.0,
			max:      -10.0,
			values:   []float64{-60.0, -30.0, -5.0},
			expected: []float64{-30.0},
		},
		{
			name:     "Min y Max iguales",
			min:      10.0,
			max:      10.0,
			values:   []float64{5.0, 10.0, 15.0},
			expected: []float64{10.0},
		},
	}

	// recorremos los escenarios de prueba
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Ejecutamos la función que creamso en el main.go
			result := minmax(tt.min, tt.max, tt.values...)

			// Comparamos el resultado con lo esperado
			if !reflect.DeepEqual(result, tt.expected) && !(len(result) == 0 && len(tt.expected) == 0) {
				t.Errorf("Error en operación: %s\nInputs: min=%.2f, max=%.2f, values=%v\nResult: %v, Expected: %v",
					tt.name, tt.min, tt.max, tt.values, result, tt.expected)
			}
		})
	}
}
