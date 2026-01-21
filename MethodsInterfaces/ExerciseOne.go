package main

import "fmt"

type Producto struct {
	Nombre    string
	Cantidad  int
	Precio    float64
	Descuento float64
}
type Servicio struct {
	Nombre      string
	TarifaPlana float64
}

func (p Producto) PrecioFinal() float64 {
	return p.Precio * float64(p.Cantidad)
}
func (p Producto) ObtenerNombre() string {
	return p.Nombre
}
func (s Servicio) PrecioFinal() float64 {
	return s.TarifaPlana
}
func (s Servicio) ObtenerNombre() string {
	return s.Nombre
}

type Cobro interface {
	PrecioFinal() float64
	ObtenerNombre() string
}

func main() {
	carrito := []Cobro{
		Producto{Nombre: "iPhone 15", Cantidad: 10, Precio: 900, Descuento: 0.10},
		Servicio{Nombre: "Streaming", TarifaPlana: 15.99},
		Producto{Nombre: "Cargador", Cantidad: 2, Precio: 25, Descuento: 0.20},
		Servicio{Nombre: "Nube", TarifaPlana: 9.99},
	}

	var total float64
	for _, item := range carrito {
		total += item.PrecioFinal()
		fmt.Printf("%s: $%.2f\n", item.ObtenerNombre(), item.PrecioFinal())
	}

	fmt.Printf("Total a pagar: $%.2f\n", total)
}
