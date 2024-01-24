package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	h1 := func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("template.html")
		Data := "Greetings!!! of the day"
		//Multiple data can be sent by declaring inside a structure
		t.Execute(w, Data)
	}

	http.HandleFunc("/", h1)

	http.HandleFunc("/cat", catHandler)

	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}

}

func catHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is message 4")
}
