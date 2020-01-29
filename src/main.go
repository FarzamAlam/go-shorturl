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

	// r := mux.NewRouter()
	// r.HandleFunc("/create", post).Methods(http.MethodPost)
	// r.HandleFunc("/", get).Methods(http.MethodPost)

	// // log.Fatal(http.ListenAndServe(":8090", r))

}

func queryBucket(path string) (string, bool, error) {
	// Query Db to check if the src is present in the db. If found then found is true.
	dest, found, err := shorturl.FindURL(path)
	return dest, found, err
}

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
