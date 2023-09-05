// Package rpchan implements Go's channel semantics over a TCP connection using net/rpc.
package rpchan

import (
	"errors"
	"net"
	"net/rpc"
	"sync"

	"github.com/lucafmarques/rpchan/internal"
)

// RPChan
type RPChan[T any] struct {
	addr     string
	setupC   func()
	setupR   func()
	client   *rpc.Client
	listener *net.Listener
	receiver *internal.Receiver[T]
}

// Send imitates a Go channels' send operation, but on an [RPChan].
//
//	ch <- "value"
//	err := rpchan.Send("value")
//
// Since Send involves a network call, it can return an error.
// A call to Send, much like sending over a Go channel, may block.
func (ch *RPChan[T]) Send(v any) error {
	ch.setupC()
	return ch.client.Call("Channel.Send", v, nil)
}

// Receive imitates a Go channels' receive operation, but on an [RPChan].
//
//	v, ok := <- ch
//	v, ok = rpchan.Receive()
//
// A call to Receive, much like receiving over a Go channel, may block.
func (ch *RPChan[T]) Receive() (*T, bool) {
	v, ok := <-ch.Iter()
	return v, ok
}

// Iter returns the underlying channel of type *T.
// Iter is a convenience method for iteration over an [RPChan].
//
//	for v := range rpchan.Iter() {
//		doSomething(v)
//	}
func (ch *RPChan[T]) Iter() <-chan *T {
	ch.setupR()
	return ch.receiver.Channel
}

// Close imitates a close() call on a normal Go channel.
// Since closing involves I/O, it can return an error containing
// the RPC client's Close() error and/or the TCP listener Close() error.
func (ch *RPChan[T]) Close() error {
	var errs []error

	if ch.client != nil {
		errs = append(errs, ch.client.Close())
	}
	if ch.listener != nil {
		errs = append(errs, (*(ch.listener)).Close())
	}

	return errors.Join(errs...)
}

// New creates an RPChan[T], with an optional bufferSize, over
// addr and returns a reference to it.
//
// The returned RPChan[T] will not start a client nor a server unless their
// related methods are called, [RPChan.Send] and [RPChan.Receive], respectively.
func New[T any](addr string, buf ...uint) *RPChan[T] {
	var bufsize uint
	if len(buf) > 0 {
		bufsize = buf[0]
	}

	ch := &RPChan[T]{addr: addr}
	ch.setupC = sync.OnceFunc(func() {
		cli, err := rpc.Dial("tcp", addr)
		if err != nil {
			panic(err)
		}

		ch.client = cli
	})
	ch.setupR = sync.OnceFunc(func() {
		srv := rpc.NewServer()
		rec := internal.NewReceiver[T](bufsize)
		srv.RegisterName("Channel", rec)

		list, err := net.Listen("tcp", addr)
		if err != nil {
			panic(err)
		}

		go func() {
			srv.Accept(list)
			close(rec.Channel)
		}()

		ch.receiver = rec
		ch.listener = &list
	})

	return ch
}
