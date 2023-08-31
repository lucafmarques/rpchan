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
	ch, _ := rpchan.NewChannel[T](":9091", 100)
	for {
		v, ok := ch.Receive()
		fmt.Printf("%+v - %v\n", v, ok)
		if !ok {
			return
		}
	}
}
