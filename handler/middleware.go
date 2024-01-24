package main

import (
	"fmt"
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Middleware!")
}


func main() {
	mux := http.NewServeMux()
	mux.Handle("/",logMiddleware(http.HandlerFunc(mainHandler)))
	http.ListenAndServe(":9009",mux)
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Printf("Middleware function\n")
		next.ServeHTTP(w,r)
	})
}


