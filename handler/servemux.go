package main

import (
	"fmt"
	"net/http"
)

func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("/",handler)
	http.ListenAndServe(":9006",nil)
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"ServeMux Creation")
}