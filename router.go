package main

import "net/http"

func foo (handler http.HandlerFunc) http.HandlerFunc {//todo auto register default routes (send type in param or something)
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}
