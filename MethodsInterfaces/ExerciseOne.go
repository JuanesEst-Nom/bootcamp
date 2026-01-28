package main

import (
	"errors"
	"fmt"
)

// DEFINICIÓN DE INTERFACES
type Cobro interface {
	PrecioFinal() float64
	ObtenerNombre() string
}

// ESTRUCTURAS
type Producto struct {
	Nombre   string
	Cantidad int
	Precio   float64
}

type Servicio struct {
	Nombre      string
	TarifaPlana float64
}

// MÉTODOS Y ERRORES
func (p Producto) PrecioFinal() float64 {
	return p.Precio * float64(p.Cantidad)
}

func (p Producto) ObtenerNombre() string {
	return p.Nombre
}

// Manejo de errores
func (p Producto) ValidarStock() error {
	if p.Cantidad <= 0 {

		return errors.New("error: el producto " + p.Nombre + " no tiene stock disponible")
	}
	return nil
}

// Implementación para Servicio
func (s Servicio) PrecioFinal() float64 {
	return s.TarifaPlana
}

func (s Servicio) ObtenerNombre() string {
	return s.Nombre
}

// 4. Interfaz stringer
func (p Producto) String() string {
	return fmt.Sprintf("PRODUCTO: %s (Cant: %d)", p.Nombre, p.Cantidad)
}

func (s Servicio) String() string {
	return fmt.Sprintf("SERVICIO: %s", s.Nombre)
}

func main() {
	// CARRITO DE COMPRAS
	carrito := []Cobro{
		Producto{Nombre: "iPhone 15", Cantidad: 10, Precio: 900},
		Servicio{Nombre: "Netflix", TarifaPlana: 15.99},
		Producto{Nombre: "Cargador", Cantidad: 0, Precio: 25},
	}

	fmt.Println("--- Carrito ---")

	var total float64

	for _, item := range carrito {
		if p, ok := item.(Producto); ok {
			err := p.ValidarStock()
			if err != nil {
				fmt.Println(err)
				continue
			}
		}

		fmt.Printf("%v | Subtotal: $%.2f\n", item, item.PrecioFinal())
		total += item.PrecioFinal()
	}

	fmt.Printf("\nTOTAL FINAL: $%.2f\n", total)
}
