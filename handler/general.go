package handler

import (
	"fmt"
	"net/http"
)

func IWasHere(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("I was here.")
		h.ServeHTTP(w, r)
	})
}

func FileServer(filepath string) http.Handler {
	return http.FileServer(http.Dir(filepath))
}

func File(file string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, file)
	})
}
