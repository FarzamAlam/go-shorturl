package shorturl

import (
	"fmt"
	"net/http"
	"os"
)

func RedirectHandler() http.HandlerFunc {
	fmt.Println("Inside RedirectHandler()")
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Inside the return of RedirectHandler")
		dest, found, err := FindURL(r.URL.Path[1:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error in RedirectHandler() : %v", err)
			os.Exit(1)
		}
		if found {
			http.Redirect(w, r, dest, http.StatusFound)
		} else if r.URL.Path == "/" {
			fmt.Fprintln(w, "Welcome to go-shorturl")
		} else {
			fmt.Fprintf(w, "%q Path in NOT FOUND", r.URL.Path)
		}
	}
}
