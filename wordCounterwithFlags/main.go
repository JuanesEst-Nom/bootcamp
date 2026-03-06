package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	countLines := flag.Bool("l", false, "Cuenta líneas ")
	countBytes := flag.Bool("b", false, "Cuenta bytes")
	flag.Parse()

	fmt.Println(count(os.Stdin, *countLines, *countBytes))
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
