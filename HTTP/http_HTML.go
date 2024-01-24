package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "<h1>This is message 1</h1>")
		fmt.Fprintf(w, "<h3>This is message 2</h>")
	})

	h1 := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "<h1>This is message 3</h1>")
	}

	http.HandleFunc("/dog", h1)

	http.HandleFunc("/cat", catHandler)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}

}

func catHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is message 4")
}
