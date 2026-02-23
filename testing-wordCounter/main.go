package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func CountLogic(input io.Reader, countLines bool) int {
	scanner := bufio.NewScanner(input)
	count := 0

	for scanner.Scan() {
		line := scanner.Text()

		if strings.EqualFold(strings.TrimSpace(line), "exit") {
			break
		}

		if countLines {
			count++
		} else {
			count += len(strings.Fields(line))
		}
	}
	return count
}

func main() {
	countLines := false
	if len(os.Args) > 1 && os.Args[1] == "-l" {
		countLines = true
	}

	if countLines {
		fmt.Println("Counting lines (type 'exit' to stop):")
	} else {
		fmt.Println("Counting words (type 'exit' to stop):")
	}

	result := CountLogic(os.Stdin, countLines)

	if countLines {
		fmt.Printf("Total lines: %d\n", result)
	} else {
		fmt.Printf("Total words: %d\n", result)
	}
}
