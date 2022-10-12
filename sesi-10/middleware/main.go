package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	endpoint := http.HandlerFunc(greet)

	mux.Handle("/", middleware1(middleware2(endpoint)))

	fmt.Println("listening on PORT")

	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func middleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware1")
		next.ServeHTTP(w, r)
	})
}
func middleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware2")
		next.ServeHTTP(w, r)
	})
}
