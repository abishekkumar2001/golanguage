package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/oldpath", func(w http.ResponseWriter, r *http.Request) {
		// Redirect to a new path ("/newpath") with status code 302 (Found)
		http.Redirect(w, r, "/newpath", http.StatusFound)
	})

	http.HandleFunc("/newpath", func(w http.ResponseWriter, r *http.Request) {
		// Handle the new path ("/newpath")
		fmt.Fprintln(w, "This is the new path.")
	})

	fmt.Println("Running on port 9003")
	err := http.ListenAndServe(":9003", nil)
	if err != nil {
		fmt.Println(err)
	}
}
