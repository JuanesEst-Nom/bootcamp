package main

import (
	"flags/fl"
	"fmt"
)

func main() {
	lines := fl.Bool("-l", false, "Flag to determine if the program should count words or lines")

	fmt.Println("Before parsing:", *lines)
	fl.Parse()
	fmt.Println("After parsing:", *lines)

	if *lines {
		fmt.Println("The program should count lines")
	} else {
		fmt.Println("The program should count words")
	}
}
