package main

import (
	"fmt"
	"time"

	"github.com/lucafmarques/rpchan"
)

type T struct {
	A, B int
	C    string
}

func main() {
	ticker := time.NewTicker(time.Millisecond * 100)
	t := T{
		A: 1,
		B: 2,
		C: "string",
	}
	ch, _ := channel.NewChannel[T]("", ":9091")
	for range ticker.C {
		fmt.Println(ch.Send(&t))
	}
}
