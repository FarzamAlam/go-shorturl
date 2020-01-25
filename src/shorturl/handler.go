package shorturl

import (
	"fmt"
	"net/http"
)

func RedirectHandler(dest string) http.HandlerFunc {
	fmt.Println("Inside RedirectHandler()")
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, dest, http.StatusFound)
	}
}
