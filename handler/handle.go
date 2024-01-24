package main

import (
	"fmt"
	"net/http"
)

/* Interface of Handler

type Handler interface {
    ServeHTTP(http.ResponseWriter, *http.Request)
}  */

type myHandler struct{}
// MyHandler is a custom HTTP handler.

//Implementing the handler interface.Here ServeHTTP implements the http.Handler interface.
func (h myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Handle in Go!")
}

func main(){
    //Creating an instance of myHandler
    h1 := myHandler{}

    http.Handle("/",h1)
    http.ListenAndServe(":9005",nil)
}