package main

import "fmt"

type Producto struct {
	Nombre    string
	Cantidad  int
	Precio    float64
	Descuento float64
}

func (p Producto) CalcularStock() float64 {
	return p.Precio * float64(p.Cantidad)
}

func (p Producto) AplicarDescuento() float64 {
	return p.Precio * (1 - p.Descuento)
}

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
	iphone := Producto{
		Nombre:    "Iphone 17 Pro Max",
		Cantidad:  1,
		Precio:    999.99,
		Descuento: 0.10,
	}

	fmt.Printf("Stock total de %s: $%.2f\n", iphone.Nombre, iphone.CalcularStock())
	fmt.Println("Estado del inventario:", iphone.estadoInventario())
	fmt.Printf("Precio con descuento: $%.2f\n", iphone.AplicarDescuento())
}
