package main

import (
	"context"
	"fmt"
)

//The main entity in the package is Context itself, which is an interface. It has only four methods:

/* type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
} */

func main() {
	ctx := context.Background()
	fmt.Println("ctx.Err() : ", ctx.Err())
	fmt.Println("ctx.Done() : ", ctx.Done())
	fmt.Println("ctx.Value(\"key\") : ", ctx.Value("key"))
	fmt.Print("ctx.Deadline() : ")
	fmt.Print(ctx.Deadline())
}