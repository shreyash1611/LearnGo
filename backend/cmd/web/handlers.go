package main

import (
	// "log"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Handler shape (stdlib contract):
//   func(http.ResponseWriter, *http.Request)
// w = response pipe (already given by the server; you WRITE into it, you don't assign w)
// r = incoming request

// http.Error(w ResponseWriter, error string, code int)  →  (no return)
// Writes status code + plain-text body. Forces Content-Type: text/plain.
func notfound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 Page Not LOL Found", http.StatusNotFound) // StatusNotFound = 404
}

func notAvailable(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Some issue. Please recheck the code.", http.StatusNotFound)
}

func wrongMethod(w http.ResponseWriter, r *http.Request) {
	// WriteHeader(statusCode int)  →  (no return)  — send status; do BEFORE body (or first Write sends 200)
	w.WriteHeader(405) // 405 Method Not Allowed
	// Write(b []byte) (int, error)  — append body bytes through w
	w.Write([]byte("Method Not Allowed by shreyash"))
}

func home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
		"./ui/html/partials/nav.tmpl",
	}

	// template.ParseFiles(filenames ...string) (*Template, error)
	// Loads .tmpl files into one set of named chunks (define "base", "nav", …). Paths = cwd-relative.
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		notAvailable(w, r)
		return
	}

	// (*Template).ExecuteTemplate(wr io.Writer, name string, data any) error
	// Renders named template into wr (here w). "base" pulls in {{template "nav"}}, "title", "main".
	// data=nil → no dynamic values. Only returns error (nil = HTML already written into w).
	err = tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		notAvailable(w, r)
		return
	}

	// Extra Write APPENDS after the HTML (usually remove later).
	// Write(b []byte) (int, error)
	w.Write([]byte("Hello From Snippetbox!"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	// r.Method string  — field, not a function ("GET", "POST", …)
	// http.MethodGet = "GET"
	if r.Method != http.MethodGet {
		// Header() http.Header
		// Header.Set(key, value string)  →  (no return)  — replace header values for key
		w.Header().Set("Allow", "GET")
		w.WriteHeader(405) // WriteHeader(statusCode int)
		w.Write([]byte("Method Not Allowed")) // Write([]byte) (int, error)
		return
	}

	// r.URL *url.URL
	// (*url.URL).Query() url.Values           — parse ?id=12
	// (url.Values).Get(key string) string     — first value or ""
	// strconv.Atoi(s string) (int, error)     — string → int
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 0 {
		notfound(w, r)
		return
	}

	w.WriteHeader(200) // WriteHeader(statusCode int) — optional; Write alone defaults to 200
	// strconv.Itoa(i int) string  — int → string (unused below; kept for reference)
	// fmt.Fprintf(w io.Writer, format string, a ...any) (n int, err error)  — print into w
	fmt.Fprintf(w, "Viewing a valid snippet with id: %d", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// http.MethodPost = "POST"
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")                             // Header().Set(key, value string)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		wrongMethod(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // before Write/WriteHeader
	w.WriteHeader(200)                                               // WriteHeader(statusCode int)
	w.Write([]byte(`{"message": "Creating a snippet"}`))             // Write([]byte) (int, error)
}
