package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("POST /posts/{id}", postWithIDInThePath)
	http.HandleFunc("GET /gets/{id}", getWithIDInThePath)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func postWithIDInThePath(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	_, err := w.Write([]byte("post: '" + id + "'"))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func getWithIDInThePath(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	_, err := w.Write([]byte("get: '" + id + "'"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
