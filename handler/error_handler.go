package main

import (
	"fmt"
	"net/http"
)

func errorHandler(w http.ResponseWriter, r *http.Request){
	http.Error(w,"Internal Server error", http.StatusInternalServerError)
	fmt.Println("Error code : ", http.StatusInternalServerError)
}

func errorHandler1(w http.ResponseWriter, r *http.Request){
	http.Error(w,"Bad Request", http.StatusBadRequest)
	fmt.Println("Error code : ",http.StatusBadRequest)
}

func errorHandler2(w http.ResponseWriter, r *http.Request){
	http.Error(w,"Status not found", http.StatusNotFound)
	fmt.Println("Error code : ",http.StatusNotFound)
}

// http.StatusUnauthorized (401) for authentication errors
// http.StatusForbidden (403) for authorization errors

func main(){
	http.HandleFunc("/err",errorHandler)
	http.HandleFunc("/err1",errorHandler1)
	http.HandleFunc("/err2",errorHandler2)
	fmt.Println("Running in port 9003")
	err := http.ListenAndServe(":9003",nil)

	if err != nil{
		fmt.Println(err)
	}
}