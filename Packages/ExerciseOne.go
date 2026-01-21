package main

import "fmt"

// 1. Definimos el molde
type Producto struct {
	Nombre   string
	Cantidad int
	Precio   float64
}

func main() {
	miLaptop := Producto{
		Nombre:   "Portatil",
		Cantidad: 5,
		Precio:   1200000.50,
	}

	total := float64(miLaptop.Cantidad) * miLaptop.Precio

	fmt.Printf("El producto %s tiene un total de: %.2f\n", miLaptop.Nombre, total)
}
