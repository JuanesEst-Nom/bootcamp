package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// CountLogic recibe un io.Reader (puede ser os.Stdin o un String en un test)
func CountLogic(input io.Reader, countLines bool) int {
	scanner := bufio.NewScanner(input)
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		if strings.ToLower(strings.TrimSpace(line)) == "exit" {
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

	// Pasamos os.Stdin a nuestra función lógica
	result := CountLogic(os.Stdin, countLines)
	fmt.Println(result)
}
