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
	t := T{
		A: 1,
		B: 2,
		C: "string",
	}
	ch := rpchan.New[T](":9091")
	fmt.Println(ch.Send(&t))
	fmt.Println(ch.Close())
}
