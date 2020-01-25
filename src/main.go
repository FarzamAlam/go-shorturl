package main

import (
	"fmt"
	"log"
	"net/http"
	"shorturl"
)

func main() {

	dest := "https://golang.org/pkg/net/http"
	src := "/gobaby"

	redirect := shorturl.RedirectHandler(dest)
	http.HandleFunc(src, redirect)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			fmt.Fprintln(w, "Welcome to go-shorurl")
		} else {
			fmt.Fprintf(w, "%q Path in NOT FOUND", r.URL.Path)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
