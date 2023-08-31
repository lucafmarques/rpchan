package queue

import (
	"context"
	"net"
	"net/rpc"
	"sync"

	"github.com/lucafmarques/rpc-queue/internal/receiver"
)

var client *rpc.Client

var setup = sync.OnceFunc(func() {
	var err error
	client, err = rpc.Dial("tcp", "localhost:9091")
	if err != nil {
		panic(err)
	}
})

func Send(t any) error {
	setup()
	return client.Call("Queue.Receive", t, nil)
}

func Listen[T any](ctx context.Context) <-chan *T {
	rec := &receiver.Receiver[T]{Queue: make(chan *T)}
	srv := rpc.NewServer()
	list, err := net.Listen("tcp", ":9091")
	if err != nil {
		panic(err)
	}

	go func() {
		err := srv.RegisterName("Queue", rec)
		if err != nil {
			panic(err)
		}
		go srv.Accept(list)

		<-ctx.Done()
		close(rec.Queue)
	}()

	return rec.Queue
}
