package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Determine the counting mode
	countLines := false
	if len(os.Args) > 1 && os.Args[1] == "-l" {
		countLines = true
	}
	scanner := bufio.NewScanner(os.Stdin)
	count := 0

	if countLines {
		fmt.Println("Counting lines (type 'exit' to stop):")
	} else {
		fmt.Println("Counting words (type 'exit' to stop):")
	}

	// Capture input line by line
	for scanner.Scan() {
		line := scanner.Text()

		if strings.EqualFold(strings.TrimSpace(line), "exit") {
			break
		}

		if countLines {
			count++
		} else {
			words := strings.Fields(line)
			count += len(words)
		}
	}

	// Print the final result
	if countLines {
		fmt.Printf("Total lines: %d\n", count)
	} else {
		fmt.Printf("Total words: %d\n", count)
	}
}
