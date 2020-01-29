package main

import (
	"log"
	"net/http"
	"shorturl"
)

func main() {

	handler := shorturl.RedirectHandler()

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
