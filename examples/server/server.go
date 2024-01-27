package main

import (
	"fmt"

	"github.com/lucafmarques/rpchan"
)

type T struct {
	A, B int
	C    string
}

func main() {
	ch := rpchan.New[T](":9091", 100)
	for v := range ch.Listen() {
		fmt.Printf("%+v - %T\n", v, v)
	}
}
