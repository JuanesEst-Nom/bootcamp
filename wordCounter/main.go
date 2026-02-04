package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Determinar el modo de conteo
	countLines := false
	if len(os.Args) > 1 && os.Args[1] == "-l" {
		countLines = true
	}
	scanner := bufio.NewScanner(os.Stdin)
	count := 0

	if countLines {
		fmt.Println("Contando líneas (escribe 'exit' para parar):")
	} else {
		fmt.Println("Contando palabras (escribe 'exit' para parar):")
	}

	// Capturar entrada línea por línea
	for scanner.Scan() {
		line := scanner.Text()

		// Condición de salida
		if strings.ToLower(strings.TrimSpace(line)) == "exit" {
			break
		}

		if countLines {
			count++
		} else {
			words := strings.Fields(line)
			count += len(words)
		}
	}

	// 3. Imprimir resultado final
	if countLines {
		fmt.Printf("Lineas totales: %d\n", count)
	} else {
		fmt.Printf("Palabras totales: %d\n", count)
	}
}
