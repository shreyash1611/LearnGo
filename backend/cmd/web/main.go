package main

import (
	"log"
	"net/http"
	"fmt"
	"os"
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



func main() {
	homemux:= http.NewServeMux()
	
	homemux.HandleFunc("/{$}", home) // we define what we want the handler to do when the user visits home
	homemux.HandleFunc("/snippet/view", snippetView)
	homemux.HandleFunc("/snippet/create", snippetCreate)
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", homemux)
	if err == nil {
		fmt.Println("Server started successfully")
	} else {
		fmt.Println("Server started with error: ", err)
		os.Exit(1)
	}
	
	// log.Fatal(err)
	
}