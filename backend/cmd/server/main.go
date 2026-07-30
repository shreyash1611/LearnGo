package main

import (
	"log"
	"net/http"
)

// package main                        ← keyword + special name
// import "net/http"                   ← keyword + stdlib package path
// 
// http                                ← package alias from import
//   .HandleFunc                       ← stdlib function
//   .ListenAndServe                   ← stdlib function
//   .ResponseWriter                   ← stdlib type
//   .Request                          ← stdlib type
// 
// "/"                                 ← your path string
// ":8080"                             ← your port string
// w, r, err                           ← YOUR variable names
// func(...) { ... }                   ← your handler body
// "Hello, World!"                     ← your response text
// nil                                 ← language built-in value
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello From Snippetbox!"))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home) // we define what we want the handler to do when the user visits home
	
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	
	log.Fatal(err)
	
}