package main

import (
	"log"
	"net/http"
	"shorturl"
)

func main() {

	handler := shorturl.RedirectHandler()

	log.Fatal(http.ListenAndServe(":8080", handler))
}
