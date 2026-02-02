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
	minStr := getInput()
	min, _ := strconv.ParseFloat(minStr, 64)

	fmt.Print("Max: ")
	maxStr := getInput()
	max, _ := strconv.ParseFloat(maxStr, 64)

	if min > max {
		fmt.Println("Error: min no puede ser mayor que max")
		return
	}

	fmt.Print("Enter all values separated by space: ")
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
