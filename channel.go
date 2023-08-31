package rpchan

import (
	"net"
	"net/rpc"
	"sync"

	receiver "github.com/lucafmarques/rpchan/internal"
)

type rpchan[T any] struct {
	addr     string
	setupC   func()
	setupR   func()
	client   *rpc.Client
	receiver *receiver.Receiver[T]
}

func (ch *rpchan[T]) Send(v any) error {
	ch.setupC()
	return ch.client.Call("Channel.Send", v, nil)
}

func (ch *rpchan[T]) Receive() (*T, bool) {
	ch.setupR()
	v, ok := <-ch.receiver.Channel()
	return v, ok
}

func (ch *rpchan[T]) Iter() <-chan *T {
	ch.setupR()
	return ch.receiver.Channel()
}

func New[T any](addr string, buf ...uint) (*rpchan[T], error) {
	var bufsize uint
	if len(buf) > 0 {
		bufsize = buf[0]
	}

	ch := &rpchan[T]{addr: addr}
	ch.setupC = sync.OnceFunc(func() {
		cli, err := rpc.Dial("tcp", addr)
		if err != nil {
			panic(err)
		}

		ch.client = cli
	})
	ch.setupR = sync.OnceFunc(func() {
		srv := rpc.NewServer()
		rec := receiver.NewReceiver[T](bufsize)

		list, err := net.Listen("tcp", addr)
		if err != nil {
			panic(err)
		}
		if err := srv.RegisterName("Channel", rec); err != nil {
			panic(err)
		}

		go func() {
			srv.Accept(list)
			rec.Close()
		}()

		ch.receiver = rec
	})

	return ch, nil
}
