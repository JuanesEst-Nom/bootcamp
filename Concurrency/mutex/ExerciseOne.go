package main

import (
	"sync" // Paquete para Mutex y WaitGroup
)

type sitioWeb struct {
	visitas int
	mutex   sync.Mutex
}

func (s *sitioWeb) registrarVisita(wg *sync.WaitGroup) {
	defer wg.Done()
	s.mutex.Lock()
	s.visitas++
	s.mutex.Unlock()
}

func main() {

	var wg sync.WaitGroup
	sitio := sitioWeb{}

	numVisitas := 1000
	wg.Add(numVisitas)

	for i := 0; i < numVisitas; i++ {
		go sitio.registrarVisita(&wg)
	}

	wg.Wait()
	println("Total de visitas registradas:", sitio.visitas)

}
