package urlshort

import (
	"net/http"
)

func MapHandler(pathsToUrls map[string]string,fallback http.Handler)http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		url := r.URL.Path
		if dest, ok := pathsToUrls[url];ok{
			 http.Redirect(w ,r,dest,http.StatusFound)}
		fallback.ServeHTTP(w,r)
	}
}

