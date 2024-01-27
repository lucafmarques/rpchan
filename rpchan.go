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

// Send implements Go channels' send operation.
//
// A call to Send, much like sending over a Go channel, may block.
// The first call to Send may panic on dialing the TCP address of the RPChan.
//
// Since this involves a network call, Send can return an error.
func (ch *RPChan[T]) Send(v T) error {
	ch.setupC()
	return ch.client.Call("Channel.Send", v, nil)
}

// Receive implements Go channels' receive operation.
//
// A call to Receive, much like receiving over a Go channel, may block.
// The first call to Receive may panic on listening on the TCP address of RPChan.
func (ch *RPChan[T]) Receive() (*T, bool) {
	ch.setupR()
	v, ok := <-ch.receiver.Channel
	return v, ok
}

// Listen implements a GOEXPERIMENT=rangefunc iterator.
//
// When used in a for-range loop it works exacly like a Go channel.
func (ch *RPChan[T]) Listen() func(func(T) bool) {
	return func(yield func(T) bool) {
		for {
			if v, ok := ch.Receive(); !ok || !yield(*v) {
				return
			}
		}
	}
}

// Close implements the close built-in.
// Since closing involves I/O, it can return an error containing
// the RPC client's Close() error and/or the TCP listener Close() error.
func (ch *RPChan[T]) Close() error {
	var errs []error

	if ch.client != nil {
		errs = append(errs, ch.client.Call("Channel.Close", 0, nil), ch.client.Close())

	}
	if ch.listener != nil {
		errs = append(errs, (*(ch.listener)).Close())
	}
	if ch.receiver != nil {
		close(ch.receiver.Channel)
	}

	return errors.Join(errs...)
}

// New creates an RPChan[T] over addr, with an optional N buffer size, and
// returns a reference to it.
//
// The returned RPChan[T] will not start a client nor a server unless their
// related methods are called, [RPChan.Send] and [RPChan.Receive] or [RPChan.Listen], respectively.
func New[T any](addr string, n ...uint) *RPChan[T] {
	var bufsize uint
	if len(n) > 0 {
		bufsize = n[0]
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

		ch.receiver, ch.listener = rec, &list
	})

	return ch
}
