package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Go Environment Variables .... ")
	//Setenv ... Set new value
	os.Setenv("GOPATH", "2024")

	//Getenv

	gopath := os.Getenv("GOPATH")
	fmt.Println("GOPATH : ", gopath)

	shell := os.Getenv("SHELL")
	fmt.Println("The name of the current shell is : ", shell)
	path := os.Getenv("PATH")
	fmt.Println("The current path is : ", path)
	bin := os.Getenv("GOBIN")
	fmt.Println("BIN : ", bin)
	root := os.Getenv("GOROOT")
	fmt.Println("ROOT : ", root)

	//Lookupenv
	lookenv := func(key string){
		val,ok := os.LookupEnv(key)
		if !ok{
			fmt.Printf("%s not set\n", key)
        } else {
            fmt.Printf("%s=%s\n", key, val)
        }
    }
	lookenv("EDITOR")
	lookenv("SHELL")

	//Expandenv
	os.Setenv("EDITOR", "emacs")

    fmt.Println(os.ExpandEnv("My editor is $EDITOR."))
    fmt.Println(os.ExpandEnv("My shell is $SHELL."))
	
}
