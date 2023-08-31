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
	if client, err = rpc.Dial("tcp", "localhost:9091"); err != nil {
		panic(err)
	}
})

func Send(t any) error {
	setup()
	return client.Call("Queue.Push", t, nil)
}

func Listen[T any](ctx context.Context, buf uint) <-chan *T {
	srv := rpc.NewServer()
	rec := receiver.NewReceiver[T](buf)

	list, err := net.Listen("tcp", ":9091")
	if err != nil {
		panic(err)
	}
	if err := srv.RegisterName("Queue", rec); err != nil {
		panic(err)
	}

	go srv.Accept(list)
	go func() {
		<-ctx.Done()
		rec.Stop()
	}()

	return rec.Listen()
}
