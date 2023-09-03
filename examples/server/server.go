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
	for {
		v, ok := ch.Receive()
		fmt.Printf("%+v - %v\n", v, ok)
		if !ok {
			return
		}
	}

	// unreacheable code to represent hwo you could use ch.Iter()
	for v := range ch.Iter() {
		fmt.Printf("%+v\n", v)
	}
}
