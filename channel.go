package channel

import (
	"context"
	"net"
	"net/rpc"
	"sync"

	"github.com/lucafmarques/rpchan/internal"
)

var client *rpc.Client

var setup = sync.OnceFunc(func() {
	var err error
	if client, err = rpc.Dial("tcp", ":9091"); err != nil {
		panic(err)
	}
})

// Send calls the Channel.Send RPC to write to the listener's channel.
func Send(t any) error {
	setup()
	return client.Call("Channel.Send", t, nil)
}

// Receive returns the underlying channel that holds the data received via the RPC's.
// This differs semantically from a channel receive, but it provides a better API
// for callers to work with the data being received.
func Receive[T any](ctx context.Context, buf uint) <-chan *T {
	srv := rpc.NewServer()
	rec := receiver.NewReceiver[T](buf)

	list, err := net.Listen("tcp", ":9091")
	if err != nil {
		panic(err)
	}
	if err := srv.RegisterName("Channel", rec); err != nil {
		panic(err)
	}

	go srv.Accept(list)
	go func() {
		<-ctx.Done()
		rec.Close()
	}()

	return rec.Channel()
}
