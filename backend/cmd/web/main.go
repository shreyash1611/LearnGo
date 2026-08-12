package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// NewServeMux() *ServeMux
	// Empty router (path → handler phone book).
	homemux := http.NewServeMux()

	// (*ServeMux).HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	// Register: this path → call this handler. Handler returns nothing.
	// "/{$}" = exact "/" only (Go 1.22+). Register ALL routes BEFORE ListenAndServe.
	homemux.HandleFunc("/{$}", home)
	homemux.HandleFunc("/snippet/view", snippetView)
	homemux.HandleFunc("/snippet/create", snippetCreate)

	// log.Println(v ...any)  →  (no return)  — log to stderr
	log.Println("Starting server on :4000")

	// ListenAndServe(addr string, handler Handler) error
	// Bind port and serve forever (BLOCKS — code below only runs if listen fails).
	// handler = your mux. Returns error only if the server fails to start / crashes out.
	err := http.ListenAndServe(":4000", homemux)
	if err != nil {
		
		log.Fatal(err)
	}

	// log.Fatal(v ...any)  →  logs then os.Exit(1)  — common pattern: log.Fatal(err)
}
