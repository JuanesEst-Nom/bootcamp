package main

import (
	"fmt"
	"strings"
)

// 1. DECLARACIÓN DE CONSTANTES

const ImpuestoIVA = 0.19
const NombreTienda = "TechStore Central"

// 2. DEFINICIÓN DE ESTRUCTURA
type Producto struct {
	Nombre   string
	Cantidad int
	Precio   float64
	Activo   bool
}

func main() {
	// 3. DECLARACIÓN DE VARIABLES DE DISTINTO TIPO
	var mensajeBienvenida string = "Resumen de Inventario"
	descuentoAplicado := 50000.0
	unidadesMinimas := 1

	// ESTRUCTURA DEL PRODUCTO
	miLaptop := Producto{
		Nombre:   "Portatil Gaming",
		Cantidad: 5,
		Precio:   1200000.50,
		Activo:   false,
	}

	fmt.Printf("--- %s (%s) ---\n", NombreTienda, mensajeBienvenida)

	// 4. USO DE FUNCIONES

	totalSinIVA, totalConIVA := calcularTotales(miLaptop.Cantidad, miLaptop.Precio)

	infoFormateada := formatearDetalle(miLaptop.Nombre, totalConIVA, miLaptop.Activo)

	// Imprimir resultados
	fmt.Println(infoFormateada)
	fmt.Printf("Subtotal: $%.2f | IVA: $%.2f\n", totalSinIVA, totalConIVA-totalSinIVA)
	fmt.Printf("Descuento disponible: %.2f | Stock mínimo: %d\n", descuentoAplicado, unidadesMinimas)
}

// 5. FUNCIONES

func calcularTotales(cantidad int, precio float64) (float64, float64) {
	subtotal := float64(cantidad) * precio
	totalConIVA := subtotal * (1 + ImpuestoIVA)
	return subtotal, totalConIVA
}

func formatearDetalle(nombre string, precioFinal float64, disponible bool) string {
	estado := "AGOTADO"
	if disponible {
		estado = "DISPONIBLE"
		return fmt.Sprintf("PRODUCTO: %s | PRECIO FINAL: %.2f | ESTADO: %s", strings.ToUpper(nombre), precioFinal, estado)
	}
	return "Producto no disponible actualmente."
}
