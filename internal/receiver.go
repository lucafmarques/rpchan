package internal

import "errors"

// Receiver needs to be exported to be used as an RPC handler.
type Receiver[T any] struct {
	Channel chan *T
}

func NewReceiver[T any](buf uint) *Receiver[T] { return &Receiver[T]{make(chan *T, buf)} }

// Send implements the function signature for an RPC handler.
func (r *Receiver[T]) Send(item *T, ok *bool) error {
	var err error
	defer func() {
		recover()
		// safe to assume that we're recovering from a send on a closed channel
		err = errors.New("queue is closed")
	}()

	r.Channel <- item

	return err
}
