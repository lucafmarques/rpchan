package main

import (
	"fmt"

	channel "github.com/lucafmarques/rpchan"
)

type T struct {
	A, B int
	C    string
}

func main() {
	ch, err := channel.NewChannel[T]("", ":9091", 1)
	if err != nil {
		panic(err)
	}
	for {
		v, ok := ch.Receive()
		fmt.Printf("%+v - %v\n", v, ok)
		if !ok {
			return
		}
	}
}
