package main

import (
	"log"
	"net/http"
	// "fmt"
	"strconv"
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


func notfound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 Page Not LOL Found",http.StatusNotFound)
}

func wrongMethod(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(405)
	w.Write([]byte("Method Not Allowed by shreyash"))
}

// http.HandleFunc("/", ...)           // registers on the *default* global mux
// http.ListenAndServe(":4000", nil)   // nil = use that default mux
func home(w http.ResponseWriter, r *http.Request) {
	
	w.Write([]byte("Hello From Snippetbox!"))

}


func snippetView(w http.ResponseWriter, r *http.Request) {
	
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 0 { // we only want no errors or real numbers
		notfound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}
	w.WriteHeader(200)
	w.Write([]byte("Viewing a valid snippet with id: " + strconv.Itoa(id)))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		wrongMethod(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte(`{"message": "Creating a snippet"}`))
}

func main() {
	homemux:= http.NewServeMux()
	homemux.HandleFunc("/{$}", home) // we define what we want the handler to do when the user visits home
	homemux.HandleFunc("/snippet/view", snippetView)
	homemux.HandleFunc("/snippet/create", snippetCreate)
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", homemux)
	
	log.Fatal(err)
	
}