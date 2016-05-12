package main

import "net/http"

//TODO auto register default routes (send type in param or something)
func foo(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}
