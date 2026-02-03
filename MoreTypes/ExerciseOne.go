package main

import "fmt"

type Producto struct {
	Nombre    string
	Cantidad  int
	Precio    float64
	Descuento float64
}

// Usamos punteros para modificar el producto real
func (p *Producto) AplicarDescuentoReal() {
	p.Precio = p.Precio * (1 - p.Descuento)
}

func main() {
	// 1. Lista de productos
	inventario := []Producto{
		{Nombre: "iPhone 15", Cantidad: 10, Precio: 900, Descuento: 0.10},
		{Nombre: "Cargador", Cantidad: 2, Precio: 25, Descuento: 0.20},
		{Nombre: "iPhone 16", Cantidad: 5, Precio: 1000, Descuento: 0.15},
	}

	// 2. Categorizar los productos
	categorias := make(map[string][]Producto)

	categorias["Telefonos"] = []Producto{inventario[0], inventario[2]}
	categorias["Accesorios"] = []Producto{inventario[1]}

	fmt.Println("--- Reporte por Categorías ---")

	// 3. Rrecoorer el map de categorías
	for cat, lista := range categorias {
		fmt.Printf("\nCategoría: %s\n", cat)
		fmt.Println("---------------------------")

		for _, p := range lista {
			// Aplicar descuento solo si es un teléfono
			if cat == "Telefonos" {
				p.AplicarDescuentoReal()
			}
			fmt.Printf("- %s: $%.2f\n", p.Nombre, p.Precio)
		}
	}

	// Crear un nuevo slice basado en uno existente
	ofertasRelampago := inventario[0:2]
	fmt.Printf("\nTenemos %d ofertas relámpago hoy.\n", len(ofertasRelampago))
}
