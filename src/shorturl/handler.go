package urlshort

import (
	"net/http"
	"os"

	"github.com/boltdb/bolt"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	db, err := bolt.Open("my.Db", 0600, nil)
	if err != nil {
		os.Exit(1)
	}
	defer db.Close()
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		if dest, ok := pathsToUrls[url]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
		}
		fallback.ServeHTTP(w, r)
	}
}
