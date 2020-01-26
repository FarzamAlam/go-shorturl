package shorturl

import (
	"fmt"
	"net/http"
	"os"
)

//RedirectHandler ... is default HandlerFunc. When ever user enters a url, it queries the db
// If found then redirects it to dest url if not then prints a message of url not found on screen.
// If url is "/", then it shows "Welcome to go-shorturl".
func RedirectHandler() http.HandlerFunc {
	fmt.Println("Inside RedirectHandler()")
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Inside the return of RedirectHandler")

		// Query Db to check if the src is present in the db. If found then found is true.
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
