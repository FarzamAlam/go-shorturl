package main

import (
	"fmt"
	"urlshort"
	"net/http"
)

func main(){
	mux := defaultMux()

	// Create a map of shorturl and url.
	pathsToUrls := map[string]string{
			"/urlshort-godoc":"https://godoc.org/github.com/gophercises/urlshort",
			"/yaml-godoc":"https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler :=urlshort.MapHandler(pathsToUrls,mux)
	fmt.Println("Starting the server at :8080")
	http.ListenAndServe(":8080",mapHandler)
}

func defaultMux()*http.ServeMux{
	mux := http.NewServeMux()
	mux.HandleFunc("/",hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello world!")
}