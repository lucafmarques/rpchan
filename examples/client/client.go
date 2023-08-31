package main

import (
	"fmt"
	"time"

	"github.com/lucafmarques/rpc-channel"
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
	ticker := time.NewTicker(time.Millisecond * 100)
	for range ticker.C {
		fmt.Println(channel.Send(&t))
	}
}
