package main

import "fmt"

func prepararBanquete(canal chan string) {
	canal <- "Entrada 🥗"
	canal <- "Plato Fuerte 🥩"
	canal <- "Postre 🍨"
	close(canal)
}

func contador(veces int, canal chan int) {
	for i := 1; i <= veces; i++ {
		canal <- i
	}
	close(canal)
}

func main() {
	tubería := make(chan string)
	go prepararBanquete(tubería)

	for plato := range tubería {
		fmt.Println("Recibido:", plato)
	}
	fmt.Println("¡Banquete terminado!")

	contadorCanal := make(chan int)
	go contador(5, contadorCanal)

	for numero := range contadorCanal {
		fmt.Println("Contador:", numero)
	}
	fmt.Println("Contador terminado!")
}
