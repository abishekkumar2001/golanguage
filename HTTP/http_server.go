package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is message 1")
	})

	h1 := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is message 2")
	}

	http.HandleFunc("/dog", h1)

	http.HandleFunc("/cat", catHandler)

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}

}

func catHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is message 3")
}
