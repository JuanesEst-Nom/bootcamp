package main

import (
	"fmt"
	"time"
)

func main() {
	canalSeguridad := make(chan string)
	canalMantenimiento := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		canalSeguridad <- "Alerta de seguridad"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		canalMantenimiento <- "Mantenimiento completado"
	}()

	fmt.Println("Esperando sistemas...")
	contador := 0
	for contador < 2 {
		select {
		case mensajeSeguridad := <-canalSeguridad:
			contador++
			fmt.Println("Mensaje recibido del canal de seguridad:", mensajeSeguridad)
		case mensajeMantenimiento := <-canalMantenimiento:
			contador++
			fmt.Println("Mensaje recibido del canal de mantenimiento:", mensajeMantenimiento)
		default:
			fmt.Println("Escaneando... 🔍")
			time.Sleep(500 * time.Millisecond)

		}
	}
	fmt.Println("Todos los sistemas operativos.")
}
