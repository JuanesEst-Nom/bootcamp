package main

import (
	"bufio"
	"flag"
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

func count(r io.Reader, lines bool, bytes bool) int {
	scanner := bufio.NewScanner(r)

	if lines {
		scanner.Split(bufio.ScanLines)
	} else if bytes {
		scanner.Split(bufio.ScanBytes)
	} else {
		scanner.Split(bufio.ScanWords)
	}

	counter := 0
	for scanner.Scan() {
		counter++
	}
	return counter
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

	var result int
	if *countBytes {

		result = count(os.Stdin, false, true)
	} else {

		result = CountLogic(os.Stdin, *countLines)
	}

	if *countLines {
		fmt.Printf("Total lines: %d\n", result)
	} else if *countBytes {
		fmt.Printf("Total bytes: %d\n", result)
	} else {
		fmt.Printf("Total words: %d\n", result)
	}
}
