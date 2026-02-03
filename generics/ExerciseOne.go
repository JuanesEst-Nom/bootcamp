package main

import "fmt"

type Producto struct {
	Nombre string
}

type Pila[T any] struct {
	elementos []T
}

func CrearPar[T any](a T, b T) []T {
	return []T{a, b}
}
func ContarElementos[T any](lista []T) int {
	contador := 0
	for range lista {
		contador++
	}
	return contador
}
func ObtenerUltimo[T any](lista []T) T {
	return lista[len(lista)-1]
}
func (p *Pila[T]) Empujar(valor T) {
	p.elementos = append(p.elementos, valor)
}

func main() {
	// 1
	//Imprimir ejemplos de uso de la función CrearPar
	// Usando strings
	parNombres := CrearPar("Manzana", "Pera")
	fmt.Println("Nombres:", parNombres)
	// Usando números (int)
	parNumeros := CrearPar(10, 20)
	fmt.Println("Números:", parNumeros)
	// Usando tus propios Structs
	p1 := Producto{Nombre: "iPhone"}
	p2 := Producto{Nombre: "MacBook"}
	parProductos := CrearPar(p1, p2)

	fmt.Println("Productos:", parProductos)

	// 2
	//Imprimir ejemplos de uso de la función ContarElementos
	//Contar elementos en una lista de strings
	listaNombres := []string{"Ana", "Luis", "Carlos"}
	cantidadNombres := ContarElementos(listaNombres)
	fmt.Println("Cantidad de nombres:", cantidadNombres)

	//Contar elementos en una lista de números
	listaNumeros := []int{1, 2, 3, 4, 5}
	cantidadNumeros := ContarElementos(listaNumeros)
	fmt.Println("Cantidad de números:", cantidadNumeros)

	// 3
	//Imprimir ejemplos de uso de la función ObtenerUltimo
	//Obtener el último elemento de una lista de strings
	ultimoNombre := ObtenerUltimo(listaNombres)
	fmt.Println("Último nombre:", ultimoNombre)

	//Obtener el último elemento de una lista de números
	ultimoNumero := ObtenerUltimo(listaNumeros)
	fmt.Println("Último número:", ultimoNumero)

	// 4
	//Imprimir ejemplos de uso del método Empujar de la estructura Pila
	//Crear una pila de enteros y empujar algunos valores
	pilaEnteros := Pila[int]{}
	pilaEnteros.Empujar(10)
	pilaEnteros.Empujar(20)
	pilaEnteros.Empujar(30)
	fmt.Println("Pila de enteros:", pilaEnteros.elementos)

	//Crear una pila de strings y empujar algunos valores
	pilaStrings := Pila[string]{}
	pilaStrings.Empujar("Hola")
	pilaStrings.Empujar("Mundo")
	fmt.Println("Pila de strings:", pilaStrings.elementos)
}
