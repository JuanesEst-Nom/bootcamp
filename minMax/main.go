package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Función para leer entrada desde la consola
func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// Función minmax que filtra valores dentro del rango [min, max]
func minmax(min float64, max float64, values ...float64) []float64 {
	var result []float64
	for _, v := range values {
		if v >= min && v <= max {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	// Leer valores mínimos y máximos
	fmt.Print("Min: ")
	min, err := strconv.ParseFloat(getInput(), 64)
	if err != nil {
		fmt.Println("Error: numero inválido")
		return
	}

	fmt.Print("Max: ")
	max, err := strconv.ParseFloat(getInput(), 64)
	if err != nil {
		fmt.Println("Error: numero inválido")
		return
	}

	if min > max {
		fmt.Println("Error: min no puede ser mayor que max")
		return
	}

	fmt.Print("Ingrese los valores separados por espacio: ")
	allValuesStr := getInput()
	fields := strings.Fields(allValuesStr)

	var valueList []float64
	for _, f := range fields {
		val, _ := strconv.ParseFloat(f, 64)
		valueList = append(valueList, val)
	}
	// Pasamos los valores a la función minmax
	filteredSlice := minmax(min, max, valueList...)

	// Imprimir el resultado
	fmt.Println(filteredSlice)
}
