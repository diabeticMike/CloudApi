package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`Hello World 3`))
	}).Methods(http.MethodGet)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`Here you are`))
	}).Methods(http.MethodGet)
	println("listen to 80")
	if err := http.ListenAndServe(":80", router); err != nil {
		panic(err)
	}
}
