package main

import (
	"log"
	"net/http"
)

// Method 1: type with ServeHTTP — satisfies http.Handler directly
type homeHandler struct{}

func (h homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

// Method 2: function converted to http.HandlerFunc type, passed to mux.Handle
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello About"))
}

// Method 3: function passed directly to mux.HandleFunc
func helpHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Help"))
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", homeHandler{})
	mux.Handle("/about", http.HandlerFunc(aboutHandler))
	mux.HandleFunc("/help", helpHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
