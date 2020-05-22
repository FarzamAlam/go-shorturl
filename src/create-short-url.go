package main

import (
	"fmt"
	"log"
	"net/http"
	"shorturl"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", get).Methods(http.MethodGet)
	r.HandleFunc("/find/{shortURL}", getURL).Methods(http.MethodGet)
	r.HandleFunc("/create/{destURL}", post).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func getURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["shortURL"]
	dest, _, _ := shorturl.FindURL(shortURL)
	fmt.Fprintf(w, "Short : %s --> %s\n", shortURL, dest)
}
func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message" :"Get Called!" }`))
}
func post(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	destURL := vars["destURL"]
	fmt.Println(destURL)
	_ = shorturl.CreateURL(destURL)
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Post Called!"}`))
}
