package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func CountLogic(input io.Reader, countLines bool, countBytes bool) int {
	scanner := bufio.NewScanner(input)
	count := 0

	if countBytes {
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			count++
		}
		return count
	}

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
	countLines := flag.Bool("l", false, "Count lines")
	countBytes := flag.Bool("b", false, "Count bytes")
	flag.Parse()

	if *countLines {
		fmt.Println("Counting lines (type 'exit' to stop):")
	} else if *countBytes {
		fmt.Println("Counting bytes:")
	} else {
		fmt.Println("Counting words (type 'exit' to stop):")
	}

	result := CountLogic(os.Stdin, *countLines, *countBytes)

	if *countLines {
		fmt.Printf("Total lines: %d\n", result)
	} else if *countBytes {
		fmt.Printf("Total bytes: %d\n", result)
	} else {
		fmt.Printf("Total words: %d\n", result)
	}
}
