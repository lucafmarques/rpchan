package receiver

import "errors"

type Receiver[T any] struct {
	queue chan *T
}

// Push implements the function signature for an rpc handler.
func (r *Receiver[T]) Push(item *T, ok *bool) (err error) {
	defer func() {
		recover()
		// safe to assume that we're recovering from a send on a closed channel
		err = errors.New("queue is closed")
	}()

	r.queue <- item

	return err
}

func (r Receiver[T]) Stop()           { close(r.queue) }
func (r Receiver[T]) Listen() chan *T { return r.queue }

func NewReceiver[T any](buf uint) *Receiver[T] { return &Receiver[T]{make(chan *T, buf)} }
