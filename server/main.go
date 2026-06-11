package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.Int("p", 8080, "port to listen on")
	host := flag.String("h", "localhost", "host to listen on")
	flag.Parse()

	s := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", *host, *port),
		Handler: newMux(),
	}

	log.Printf("Listening on %s", s.Addr)
	log.Fatal(s.ListenAndServe())
}
