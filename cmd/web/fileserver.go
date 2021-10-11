package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func fileServer(r chi.Router) {
	root := "./ui/"
	fs := http.FileServer(http.Dir(root))

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(root + r.RequestURI); os.IsNotExist(err) {
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			fs.ServeHTTP(w, r)
		}
	})
}
