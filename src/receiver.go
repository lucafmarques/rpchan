package rpc_queue

import (
	"context"
	"net"
	"net/rpc"
)

type receiverer interface {
	Listen(ctx context.Context) <-chan any
	Receive(item any, ok *bool)
	Stop()
}

var (
	t        = true
	Receiver = receiver{make(chan *any), "1001"}
)

type receiver struct {
	queue chan *any
	port  string
}

func (r receiver) Listen(ctx context.Context) <-chan any {
	c := make(chan any)
	srv := rpc.NewServer()
	list, err := net.Listen("tcp", r.port)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			select {
			case v, ok := <-r.queue:
				if !ok {
					close(c)
				}
				c <- *v
			case <-ctx.Done():
				close(c)
			}
		}
	}()

	go func() {
		err := srv.Register(r)
		if err != nil {
			panic(err)
		}
		srv.Accept(list)
	}()

	return c
}

func (r receiver) Receive(item *any, ok *bool) {
	r.queue <- item
	ok = &t
}
