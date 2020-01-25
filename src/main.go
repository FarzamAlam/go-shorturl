package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"shorturl"
)

func main() {

	//dest := "https://golang.org/pkg/net/http"
	src := "/gobaby"

	srcHash := shorturl.GenerateHash(src)[:8]
	fmt.Println(srcHash)

	// Quering the db to find the dest url
	dest, found, err := shorturl.FindURL(srcHash)

	fmt.Println("dest : " + dest)
	if err != nil {
		fmt.Fprintf(os.Stderr, "main.go : %v", err)
		os.Exit(1)
	}
	if found {
		fmt.Println("Inside found")
		redirect := shorturl.RedirectHandler(dest)
		http.HandleFunc(srcHash, redirect)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			fmt.Fprintln(w, "Welcome to go-shorurl")
		} else {
			fmt.Fprintf(w, "%q Path in NOT FOUND", r.URL.Path)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
