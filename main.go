package main

import "net/http"

func main() {

	mux := http.ServeMux{}

	server := &http.Server{
		Handler: &mux,
		Addr:    ":" + "8080"}

	server.ListenAndServe()

}
