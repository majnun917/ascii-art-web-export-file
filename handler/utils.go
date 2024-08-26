package handler

import (
	"html/template"
	"net/http"
	"strings"
)

// Customize the 404 error page
func NotFoundPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/404.html")
	if err != nil {
		http.Error(w, "Failed to parse file", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusNotFound)
	tmpl.Execute(w, nil)
}

// Custom handler to prevent directory listing
func NoDirListing(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			NotFoundPage(w, r)
			return
		}
		h.ServeHTTP(w, r)
	})
}

// Strict path matching for safety, robustness, clarity and explicitness
func PathCheck(path string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != path {
			NotFoundPage(w, r)
			return
		}
		handler(w, r)
	}
}
