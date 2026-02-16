// DON'T CHANGE THIS FILE
package main

import (
	"flag"
	"fmt"
)

func main() {
	lines := flag.Bool("l", false, "Contar líneas en lugar de palabras")
	fmt.Println("antes :", *lines)
	flag.Parse()
	fmt.Println("despues:", *lines)

	if *lines {
		fmt.Println("Contando líneas del programa")
	} else {
		fmt.Println("Contando palabras del programa")
	}
}
