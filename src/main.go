package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"GET Called!"}`))

	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message":"POST Called!"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message":"Not Found!"}`))
	}
}
func main() {

	// handler := shorturl.RedirectHandler()

	// log.Fatal(http.ListenAndServe(":8080", handler))

	r := mux.NewRouter()
	r.HandleFunc("/", get).Methods(http.MethodGet)
	r.HandleFunc("/", post).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8080", r))
}
func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Get Called!"}`))
}
func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message":"Post Called!"}`))
}
