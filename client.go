package main

import (
	"fmt"

	_ "github.com/lucafmarques/rpc-queue/producer"
)

type T struct {
	a, b int
	c    string
}

func main() {
	t := T{
		a: 1,
		b: 2,
		c: "string",
	}
	fmt.Println(Producer.Send(&t))
}
