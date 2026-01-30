package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// Definición de flags
	countLines := flag.Bool("l", false, "Cuenta líneas ")
	countBytes := flag.Bool("b", false, "Cuenta bytes")
	flag.Parse()

	// Llamada a la función count con los parámetros adecuados
	fmt.Println(count(os.Stdin, *countLines, *countBytes))
}

// Función count que cuenta líneas, palabras o bytes según los flags
func count(r io.Reader, lines bool, bytes bool) int {

	scanner := bufio.NewScanner(r)

	// organizar el escáner según los flags
	if lines {

		scanner.Split(bufio.ScanLines)
	} else if bytes {

		scanner.Split(bufio.ScanBytes)
	} else {

		scanner.Split(bufio.ScanWords)
	}

	//contador
	counter := 0
	for scanner.Scan() {
		counter++
	}

	return counter
}

// test
// echo "uno dos tres" | go run main.go -b ---- 13
// echo "uno dos tres" | go run main.go -l	---- 1
// echo "uno dos tres" | go run main.go	---- 3
