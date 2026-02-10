package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to read input from the console
func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// minmax function that filters values within the range [min, max]
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
	// Read minimum and maximum values
	fmt.Print("Min: ")
	min, err := strconv.ParseFloat(getInput(), 64)
	if err != nil {
		fmt.Println("Error: invalid number")
		return
	}

	fmt.Print("Max: ")
	max, err := strconv.ParseFloat(getInput(), 64)
	if err != nil {
		fmt.Println("Error: invalid number")
		return
	}

	if min > max {
		fmt.Println("Error: min cannot be greater than max")
		return
	}

	fmt.Print("Enter values separated by space: ")
	allValuesStr := getInput()
	fields := strings.Fields(allValuesStr)

	var valueList []float64
	for _, f := range fields {
		val, err := strconv.ParseFloat(f, 64)
		if err != nil {
			fmt.Println("Error: invalid number")
			return
		}

		valueList = append(valueList, val)

	}
	// Pass the values to the minmax function
	filteredSlice := minmax(min, max, valueList...)

	// Print the result
	fmt.Println(filteredSlice)
}
