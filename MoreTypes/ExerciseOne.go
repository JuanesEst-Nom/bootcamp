package main

import "fmt"

type Producto struct {
	Nombre    string
	Cantidad  int
	Precio    float64
	Descuento float64
}

func (p *Producto) AplicarDescuentoReal() {
	p.Precio = p.Precio * (1 - p.Descuento)
}

func (p *Producto) ActualizarPrecio(nuevoPrecio float64) {
	p.Precio = nuevoPrecio
}

func main() {
	inventario := []Producto{
		{Nombre: "iPhone 15", Cantidad: 10, Precio: 900, Descuento: 0.10},
		{Nombre: "Cargador", Cantidad: 2, Precio: 25, Descuento: 0.20},
		{Nombre: "iphone 16", Cantidad: 5, Precio: 1000, Descuento: 0.15},
	}

	fmt.Println("Aplicando descuentos permanentes...")

	for i := range inventario[:1] {
		inventario[i].AplicarDescuentoReal()
	}

	for _, item := range inventario {
		fmt.Printf("Producto: %-10s | Nuevo Precio: $%.2f\n", item.Nombre, item.Precio)
	}
	fmt.Println("________________________________")
	inventario[2].ActualizarPrecio(850)
	fmt.Printf("Nuevo precio especial para el iPhone 16 $%.2f\n", inventario[2].Precio)

}
