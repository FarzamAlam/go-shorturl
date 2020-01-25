package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"shorturl"
)

func main() {

	// dest := "https://golang.org/pkg/net/http"
	// src := "/gobaby"

	// srcHash := shorturl.GenerateHash(src)[:8]
	// fmt.Println(srcHash)
	count := 0
	var dest string
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		count++
		fmt.Printf("Inside handle func: %d\n ", count)
		dest, _, err := shorturl.FindURL(r.URL.Path[1:])
		if err != nil {
			os.Exit(1)
		}
		if r.URL.Path == "/" {
			fmt.Fprintln(w, "Welcome to go-shorurl")
		} else {
			fmt.Fprintf(w, "%q Path in NOT FOUND", r.URL.Path)
		}
	})

	http.HandleFunc("/", shorturl.RedirectHandler(dest))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
