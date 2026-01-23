package main

import "fmt"

type Producto struct {
	Nombre    string
	Cantidad  int
	Precio    float64
	Descuento float64
}

// Método para calcular el stock total de un producto
func (p Producto) CalcularStock() float64 {
	return p.Precio * float64(p.Cantidad)
}

// Método para determinar el estado del inventario
func (p Producto) estadoInventario() string {
	if p.Cantidad == 0 {
		return "No hay stock"
	} else if p.Cantidad < 5 {
		return "Bajo stock"
	} else {
		return "Stock alto"
	}
}

func main() {
	// 1. Se ejecuta al final de la función main
	defer fmt.Println("--- Fin del reporte de inventario ---")

	fmt.Println("--- Iniciando Sistema de Gestión ---")

	// 2. SLICE (Lista) de productos
	inventario := []Producto{
		{Nombre: "Iphone 17 Pro Max", Cantidad: 1, Precio: 999.99, Descuento: 0.10},
		{Nombre: "Cargador USB-C", Cantidad: 10, Precio: 25.50, Descuento: 0.0},
		{Nombre: "Airpods Pro", Cantidad: 0, Precio: 249.00, Descuento: 0.05},
	}

	// 3. Recorremos la lista de productos
	for i := 0; i < len(inventario); i++ {
		p := inventario[i]
		fmt.Printf("\nProducto: %s\n", p.Nombre)

		// Comparar valores según su precio
		switch {
		case p.Precio > 500:
			fmt.Println("Categoría: Artículo de Lujo")
		case p.Precio >= 100 && p.Precio <= 500:
			fmt.Println("Categoría: Artículo Premium")
		default:
			fmt.Println("Categoría: Artículo Estándar")
		}

		fmt.Printf("Estado: %s (Total: $%.2f)\n", p.estadoInventario(), p.CalcularStock())
	}
}
