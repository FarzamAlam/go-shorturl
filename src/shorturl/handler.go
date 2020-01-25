package m

import (
	"net/http"
)

func RedirectHandler(dest string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, dest, http.StatusFound)
	}
}
