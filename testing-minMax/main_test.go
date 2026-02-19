package main

import (
	"reflect"
	"testing"
)

func TestMinMax(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string
		min      float64
		max      float64
		values   []float64
		expected []float64
	}{
		{
			name:     "Standard values within range",
			min:      10.0,
			max:      20.0,
			values:   []float64{5.0, 10.0, 15.0, 20.0, 25.0},
			expected: []float64{10.0, 15.0, 20.0},
		},
		{
			name:     "Single value within range",
			min:      10.0,
			max:      20.0,
			values:   []float64{15.0},
			expected: []float64{15.0},
		},
		{
			name:     "Min greater than Max (impossible range)",
			min:      50.0,
			max:      10.0,
			values:   []float64{15.0, 25.0, 60.0},
			expected: nil,
		},
		{
			name:     "No values within range",
			min:      100.0,
			max:      200.0,
			values:   []float64{10.0, 20.0, 30.0},
			expected: nil,
		},
		{
			name:     "Negative Min and Max",
			min:      -50.0,
			max:      -10.0,
			values:   []float64{-60.0, -30.0, -5.0},
			expected: []float64{-30.0},
		},
		{
			name:     "Equal Min and Max",
			min:      10.0,
			max:      10.0,
			values:   []float64{5.0, 10.0, 15.0},
			expected: []float64{10.0},
		},
	}

	// Iterate through test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minmax(tt.min, tt.max, tt.values...)

			// Compare results with expected values
			if !reflect.DeepEqual(result, tt.expected) && !(len(result) == 0 && len(tt.expected) == 0) {
				t.Errorf("Operation error: %s\nInputs: min=%.2f, max=%.2f, values=%v\nResult: %v, Expected: %v",
					tt.name, tt.min, tt.max, tt.values, result, tt.expected)
			}
		})
	}
}
