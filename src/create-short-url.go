package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", get).Methods(http.MethodGet)
	r.HandleFunc("/", post).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message" :"Get Called!" }`))
}
func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Post Called!"}`))
}
