package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func startServer() (bool, error) {
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "search, %q", html.EscapeString(r.URL.Path))
	})
	http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "admin, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe("localhost:4000", nil))

	return true, nil
}
